package safeguard

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/philippseith/signalr"
)

// SignalREvent represents the root structure of a SignalR event notification
type SignalREvent struct {
	ApplianceId string    `json:"ApplianceId"`
	Name        string    `json:"Name"`
	Time        time.Time `json:"Time"`
	Message     string    `json:"Message"`
	AuditLogUri *string   `json:"AuditLogUri"`
	Data        EventData `json:"Data"`
}

// EventData represents the Data field of a SignalR event
type EventData struct {
	AccessRequestType           AccessRequestType `json:"AccessRequestType"`
	AccountDistinguishedName    string            `json:"AccountDistinguishedName"`
	AccountDomainName           string            `json:"AccountDomainName"`
	AccountHasTotpAuthenticator bool              `json:"AccountHasTotpAuthenticator"`
	AccountId                   int               `json:"AccountId"`
	AccountName                 string            `json:"AccountName"`
	ActionUserIds               []int             `json:"ActionUserIds"`
	ApproverAccessRequestUri    string            `json:"ApproverAccessRequestUri"`
	AssetId                     int               `json:"AssetId"`
	AssetName                   string            `json:"AssetName"`
	AssetNetworkAddress         string            `json:"AssetNetworkAddress"`
	AssetPlatformType           string            `json:"AssetPlatformType"`
	Comment                     *string           `json:"Comment"`
	DurationInMinutes           int               `json:"DurationInMinutes"`
	OfflineWorkflowMode         bool              `json:"OfflineWorkflowMode"`
	Reason                      *string           `json:"Reason"`
	ReasonCode                  *string           `json:"ReasonCode"`
	Requester                   string            `json:"Requester"`
	RequesterAccessRequestUri   string            `json:"RequesterAccessRequestUri"`
	RequesterId                 int               `json:"RequesterId"`
	RequesterUsername           string            `json:"RequesterUsername"`
	RequestId                   string            `json:"RequestId"`
	RequiredDate                time.Time         `json:"RequiredDate"`
	ReviewerAccessRequestUri    string            `json:"ReviewerAccessRequestUri"`
	SessionSpsNodeIpAddress     *string           `json:"SessionSpsNodeIpAddress"`
	TicketNumber                *string           `json:"TicketNumber"`
	WasCheckedOut               bool              `json:"WasCheckedOut"`
	EventName                   string            `json:"EventName"`
	EventTimestamp              time.Time         `json:"EventTimestamp"`
	ApplianceId                 string            `json:"ApplianceId"`
	EventUserId                 int               `json:"EventUserId"`
	EventUserDisplayName        string            `json:"EventUserDisplayName"`
	EventUserName               string            `json:"EventUserName"`
	EventUserDomainName         *string           `json:"EventUserDomainName"`
	AuditLogUri                 *string           `json:"AuditLogUri"`
	EventDisplayName            string            `json:"EventDisplayName"`
	EventDescription            string            `json:"EventDescription"`
}

// EventHandler manages SignalR connections and event handling for the Safeguard API
type EventHandler struct {
	client  *SafeguardClient
	ctx     context.Context
	started bool

	// logger is the applications logger.
	logger *slog.Logger

	// Channel for handling Events
	EventChannel chan SignalREvent

	// signalr hub is the signalr client used to register and listen for events from the pam appliance.
	signalr.Hub
}

// Log logs messages from the signalr client in debug mode.
func (h *EventHandler) Log(keyVals ...interface{}) error {
	h.logger.WithGroup("signalr").Debug("SIGNALR", keyVals...)
	return nil
}

// NewEventHandler creates a new EventHandler instance with the provided SafeguardClient
// and initializes the EventChannel.
func NewEventHandler(client *SafeguardClient) *EventHandler {
	return &EventHandler{
		client:       client,
		logger:       client.Logger,
		EventChannel: make(chan SignalREvent, 100), // buffered channel to prevent blocking
	}
}

// Run starts the event handler if it is not already running. It validates the access token,
// creates a SignalR connection and client, and starts the SignalR client. The function
// blocks until the SignalR client shuts down.
//
// Parameters:
//   - ctx: The context to control cancellation and timeout.
//
// Returns:
//   - error: An error if the event handler is already running, token validation fails,
//     creating the SignalR connection or client fails, or if there is an error
//     during the SignalR client's operation.
func (h *EventHandler) Run(ctx context.Context) error {
	if h.started {
		return fmt.Errorf("event handler already running")
	}

	// Validate access token
	if err := h.client.ValidateAccessToken(); err != nil {
		h.logger.Error("token validation failed", "error", err)
		return err
	}

	// Create the signalr connection
	connector := func() (signalr.Connection, error) {
		conn, err := signalr.NewHTTPConnection(
			ctx,
			fmt.Sprintf("%s/service/event/signalr", h.client.getClusterLeaderUrl()),
			signalr.WithHTTPHeaders(func() http.Header {
				header := http.Header{
					"Authorization": []string{"Bearer " + h.client.AccessToken.getUserToken()},
					"accept":        []string{"application/json"},
				}
				return header
			}),
			signalr.WithHTTPClient(h.client.HttpClient),
		)
		if err != nil {
			h.logger.Error("creating signalr connection failed", "error", err)
			return nil, err
		}
		h.logger.Info("signalr connection (re)created")
		return conn, nil
	}

	// Create the SignalR client with additional options
	client, err := signalr.NewClient(
		ctx,
		signalr.WithConnector(connector),
		signalr.WithBackoff(func() backoff.BackOff {
			bo := backoff.NewExponentialBackOff()
			bo.MaxElapsedTime = 3 * 24 * time.Hour
			return bo
		}),
		signalr.TransferFormat(signalr.TransferFormatText),
		signalr.WithReceiver(h),
		signalr.Logger(h, true), // overwrite default signalr logger
	)

	if err != nil {
		h.logger.Error("creating signalr client failed", "error", err)
		return err
	}

	h.ctx = ctx
	h.started = true

	// Start the SignalR client
	client.Start()

	// Wait for the SignalR client to shut down
	return <-client.WaitForState(ctx, signalr.ClientClosed)
}

func (h *EventHandler) NotifyEventAsync(rawEvent interface{}) {
	// Convert the raw event to JSON
	jsonData, err := json.Marshal(rawEvent)
	if err != nil {
		h.logger.Error("failed to marshal raw event", "error", err)
		return
	}

	// Unmarshal JSON data to SignalREvent struct
	var event SignalREvent
	if err := json.Unmarshal(jsonData, &event); err != nil {
		h.logger.Error("failed to unmarshal event", "error", err)
		return
	}

	// Send the event to the EventChannel
	select {
	case h.EventChannel <- event:
		h.logger.Debug("event received", "type", fmt.Sprintf("%T", rawEvent), "event", event)
	default:
		h.logger.Warn("event channel is full, dropping event", "event", event)
	}
}
