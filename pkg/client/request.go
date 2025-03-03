package client

import (
	"fmt"
	"io"
	"net/http"
)

// getReadOnlyRootUrl constructs and returns the root URL for read-only operations.
// It uses the appliance URL for operations that don't require cluster leader coordination.
//
// Returns:
//   - string: The complete root URL for read-only API operations.
func (c *SafeguardClient) getReadOnlyRootUrl() string {
	// If the appliance URL is not set, fall back to the cluster leader URL
	return fmt.Sprintf("%s/service/core/%s", c.Appliance.getUrl(), c.ApiVersion)
}

// getReadWriteRootUrl constructs and returns the root URL for write operations.
// It ensures write operations are directed to the current cluster leader.
//
// Returns:
//   - string: The complete root URL for read-write API operations.
func (c *SafeguardClient) getReadWriteRootUrl() string {
	// Update the cluster leader URL to ensure it's set correctly
	if c.ClusterLeader.isExpired() {
		c.updateClusterLeaderUrl()
	}

	// If the cluster leader URL is not set, fall back to the appliance URL
	return fmt.Sprintf("%s/service/core/%s", c.getClusterLeaderUrl(), c.ApiVersion)
}

// GetRequest makes a GET request to the specified path on the Safeguard API.
// It constructs the full URL by combining the read-only root URL with the provided path,
// creates an HTTP GET request, and sends it using the client's HTTP configuration.
//
// Parameters:
//   - path: The API endpoint path to append to the root URL.
//
// Returns:
//   - []byte: The response body from the API call.
//   - error: An error if the request fails or returns a non-successful status code.
func (c *SafeguardClient) GetRequest(path string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", c.getReadOnlyRootUrl(), path)
	logger.Debug("Preparing GET request",
		"url", url,
		"path", path,
	)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		logger.Error("Failed to create GET request",
			"error", err,
			"url", url,
		)
		return nil, err
	}

	return c.sendHttpRequest(req)
}

// PostRequest sends an HTTP POST request to the specified path with the provided body.
// It automatically handles authentication, request routing, and response processing.
//
// Parameters:
//   - path: The endpoint path to which the request will be sent.
//   - body: The request body data as an io.Reader.
//
// Returns:
//   - []byte: The response body from the API call.
//   - error: An error if the request fails or returns a non-successful status code.
func (c *SafeguardClient) PostRequest(path string, body io.Reader) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", c.getReadWriteRootUrl(), path)
	logger.Debug("Preparing POST request",
		"url", url,
		"path", path,
	)

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		logger.Error("Failed to create POST request",
			"error", err,
			"url", url,
		)
		return nil, err
	}

	return c.sendHttpRequest(req)
}

// PutRequest sends an HTTP PUT request to update resources on the Safeguard API.
// It automatically handles authentication and routes requests through the cluster leader.
//
// Parameters:
//   - path: The endpoint path for the resource to update.
//   - body: The request body containing the update data.
//
// Returns:
//   - []byte: The response body from the API call.
//   - error: An error if the request fails.
func (c *SafeguardClient) PutRequest(path string, body io.Reader) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", c.getReadWriteRootUrl(), path)
	logger.Debug("Preparing PUT request",
		"url", url,
		"path", path,
	)

	req, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		logger.Error("Failed to create PUT request",
			"error", err,
			"url", url,
		)
		return nil, err
	}

	return c.sendHttpRequest(req)
}

// DeleteRequest sends an HTTP DELETE request to remove resources.
// It ensures proper routing through the cluster leader for consistency.
//
// Parameters:
//   - path: The endpoint path identifying the resource to delete.
//
// Returns:
//   - []byte: The response body if any.
//   - error: An error if the deletion fails.
func (c *SafeguardClient) DeleteRequest(path string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", c.getReadWriteRootUrl(), path)
	logger.Debug("Preparing DELETE request",
		"url", url,
		"path", path,
	)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		logger.Error("Failed to create DELETE request",
			"error", err,
			"url", url,
		)
		return nil, err
	}

	return c.sendHttpRequest(req)
}

// sendHttpRequest handles the common logic for sending HTTP requests to the Safeguard API.
// It sets necessary headers, performs the request, and processes the response.
// Successful responses are considered to be those with status codes 200 (OK),
// 201 (Created), or 202 (Accepted).
//
// Parameters:
//   - req: A pointer to an http.Request object representing the prepared HTTP request.
//
// Returns:
//   - []byte: The response body if the request is successful.
//   - error: An error if the request fails, returns a non-successful status code,
//     or if there are issues reading the response body.
//
// The function handles logging of request details at debug level and any errors
// that occur during the request processing.
func (c *SafeguardClient) sendHttpRequest(req *http.Request) ([]byte, error) {
	c.setHeaders(req)
	logger.Debug("Sending request",
		"method", req.Method,
		"url", req.URL.String(),
		"headers", req.Header,
	)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		logger.Error("Request failed",
			"error", err,
			"method", req.Method,
			"url", req.URL,
		)
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Failed to read response body",
			"error", err,
			"method", req.Method,
			"url", req.URL,
		)
		return nil, err
	}

	logger.Debug("Response received",
		"method", req.Method,
		"url", req.URL,
		"status", resp.Status,
		"statusCode", resp.StatusCode,
		"contentLength", resp.ContentLength,
		"headers", resp.Header,
	)

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted && resp.StatusCode != http.StatusCreated {
		logger.Error("Request failed with non-success status code",
			"method", req.Method,
			"url", req.URL,
			"status", resp.Status,
			"statusCode", resp.StatusCode,
			"responseHeaders", resp.Header,
			"responseBody", string(body),
		)
		return nil, fmt.Errorf("error during %s request to %s: HTTP %d - %s", req.Method, req.URL, resp.StatusCode, string(body))
	}

	logger.Debug("Request completed successfully",
		"method", req.Method,
		"url", req.URL,
		"statusCode", resp.StatusCode,
		"bodyLength", len(body),
	)

	return body, nil
}

// setHeaders configures the HTTP request headers for Safeguard API requests.
// It applies the following headers in order:
// 1. Authorization header from the client's current authentication state
// 2. Any default headers configured in the client
// 3. Content-Type header (defaults to "application/json" if not set)
//
// Parameters:
//   - req: A pointer to an http.Request object to modify with the appropriate headers.
//
// The function modifies the request headers in place and ensures all necessary
// headers are present for successful API communication.
func (c *SafeguardClient) setHeaders(req *http.Request) {
	logger.Debug("Setting request headers", "existingHeaders", req.Header)

	req.Header = c.getAuthorizationHeader()

	for key, values := range c.DefaultHeaders {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	logger.Debug("Headers set successfully",
		"finalHeaders", req.Header,
		"method", req.Method,
		"url", req.URL,
	)
}
