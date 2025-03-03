package client

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

// NewEventHandler creates a new SignalR event handler
func NewEventHandler(client *SafeguardClient) *EventHandler {
	return &EventHandler{
		client:       client,
		logger:       client.Logger,
		EventChannel: make(chan SignalREvent, 100), // buffered channel to prevent blocking
	}
}

// Run starts the SignalR connection and event handling
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
