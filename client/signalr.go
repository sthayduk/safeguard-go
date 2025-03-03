package client

import (
	"context"
	"fmt"
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
}

// NewEventHandler creates a new SignalR event handler
func NewEventHandler(client *SafeguardClient) *EventHandler {
	return &EventHandler{
		client: client,
	}
}

// Run starts the SignalR connection and event handling
func (h *EventHandler) Run(ctx context.Context) error {
	if h.started {
		return fmt.Errorf("event handler already running")
	}

	// Validate access token
	if err := h.client.ValidateAccessToken(); err != nil {
		logger.Error("token validation failed", "error", err)
		return err
	}

	httpClient := h.client.HttpClient

	// Create the signalr connection
	connector := func() (signalr.Connection, error) {
		conn, err := signalr.NewHTTPConnection(
			ctx,
			fmt.Sprintf("%s/service/event/signalr", h.client.Appliance.getUrl()),
			signalr.WithHTTPHeaders(func() http.Header {
				header := http.Header{
					"Authorization": []string{"Bearer " + h.client.AccessToken.getUserToken()},
					"accept":        []string{"application/json"},
				}
				return header
			}),
			signalr.WithHTTPClient(httpClient),
		)
		if err != nil {
			logger.Error("creating signalr connection failed", "error", err)
			return nil, err
		}
		logger.Info("signalr connection (re)created")
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
	)

	if err != nil {
		logger.Error("creating signalr client failed", "error", err)
		return err
	}

	h.ctx = ctx
	h.started = true

	// Start the SignalR client
	client.Start()

	// Wait for the SignalR client to shut down
	return <-client.WaitForState(ctx, signalr.ClientClosed)
}

// OnReceive is called when a message is received from the SignalR server.
func (h *EventHandler) OnReceive(eventName string, payload []byte) {
	logger.Info("received event", "event", eventName, "payload", string(payload))
	// Handle the event based on the eventName and payload
}

// OnClose is called when the SignalR connection is closed.
func (h *EventHandler) OnClose(err error) {
	logger.Info("signalr connection closed", "error", err)
	// Handle the connection close event
}

func (h *EventHandler) Receive(msg string) {
	logger.Info(msg)
}
