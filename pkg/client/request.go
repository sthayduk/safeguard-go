package client

import (
	"fmt"
	"io"
	"net/http"
)

// getRootUrl constructs and returns the root URL for the Safeguard API.
func (c *SafeguardClient) getRootUrl() string {
	return fmt.Sprintf("%s/service/core/%s", c.ApplicanceURL, c.ApiVersion)
}

// GetRequest makes a GET request to the specified path on the Safeguard API.
// Parameters:
// - path: The API endpoint path.
// Returns the response body as a byte slice and an error if the request fails.
func (c *SafeguardClient) GetRequest(path string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", c.getRootUrl(), path)
	logger.Info("Making request", "url", url)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return c.sendHttpRequest(req)
}

// PostRequest sends an HTTP POST request to the specified path with the provided body.
// It constructs the full URL using the client's root URL and the given path, sets the
// Content-Type header to "application/json", and sends the request.
//
// Parameters:
//   - path: The endpoint path to which the request will be sent.
//   - body: An io.Reader containing the request body.
//
// Returns:
//   - A byte slice containing the response body.
//   - An error if the request fails or an error occurs during the request.
func (c *SafeguardClient) PostRequest(path string, body io.Reader) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", c.getRootUrl(), path)
	logger.Info("Making request", "url", url, "method", "POST")

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	return c.sendHttpRequest(req)
}

// PutRequest sends an HTTP PUT request to the specified path and returns the response body as a byte slice.
// It constructs the full URL using the client's root URL and the provided path, then creates and sends the request.
//
// Parameters:
//   - path: The path to append to the client's root URL for the PUT request.
//
// Returns:
//   - A byte slice containing the response body.
//   - An error if the request creation or sending fails.
func (c *SafeguardClient) PutRequest(path string, body io.Reader) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", c.getRootUrl(), path)
	logger.Info("Making request", "url", url, "method", "PUT")

	req, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return nil, err
	}

	return c.sendHttpRequest(req)
}

// DeleteRequest sends an HTTP DELETE request to the specified path and returns the response body.
//
// Parameters:
//   - path: The relative path to which the DELETE request should be made.
//
// Returns:
//   - []byte: The response body from the DELETE request.
//   - error: An error if the request could not be created or the HTTP request failed.
func (c *SafeguardClient) DeleteRequest(path string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", c.getRootUrl(), path)
	logger.Info("Making request", "url", url, "method", "DELETE")

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	return c.sendHttpRequest(req)
}

func (c *SafeguardClient) sendHttpRequest(req *http.Request) ([]byte, error) {
	c.setHeaders(req)
	logger.Debug("Request headers", "headers", req.Header)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted && resp.StatusCode != http.StatusCreated {
		logger.Warn("Request failed",
			"status", resp.Status,
			"method", req.Method,
			"url", req.URL,
			"statusCode", resp.StatusCode,
		)
		logger.Debug("Response details",
			"headers", resp.Header,
			"body", string(body),
		)
		return nil, fmt.Errorf("error during %s request to %s: HTTP %d - %s", req.Method, req.URL, resp.StatusCode, string(body))
	}

	return body, nil
}

// setHeaders sets the default headers and Content-Type header if not already set.
func (c *SafeguardClient) setHeaders(req *http.Request) {
	req = c.getAuthorizationHeader(req)

	for key, values := range c.DefaultHeaders {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}
}
