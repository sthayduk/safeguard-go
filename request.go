package safeguard

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
)

// SafeHeaders is a wrapper around http.Header that implements slog.LogValuer
// to safely log headers with sensitive information masked.
type SafeHeaders struct {
	headers http.Header
}

// LogValue implements slog.LogValuer to provide safe logging of HTTP headers.
// It masks sensitive authorization headers while preserving other header information.
func (sh SafeHeaders) LogValue() slog.Value {
	attrs := make([]slog.Attr, 0, len(sh.headers))

	for key, values := range sh.headers {
		if strings.ToLower(key) == "authorization" {
			// Mask authorization header to prevent token leakage
			maskedValues := make([]string, len(values))
			for i, value := range values {
				if strings.HasPrefix(strings.ToLower(value), "bearer ") && len(value) > 7 {
					// Show only "Bearer " + first 6 and last 4 characters with asterisks in between
					tokenPart := value[7:] // Remove "Bearer " prefix
					if len(tokenPart) > 10 {
						maskedValues[i] = fmt.Sprintf("Bearer %s***%s", tokenPart[:6], tokenPart[len(tokenPart)-4:])
					} else {
						maskedValues[i] = "Bearer ***"
					}
				} else {
					maskedValues[i] = "***"
				}
			}
			attrs = append(attrs, slog.Any(key, maskedValues))
		} else {
			// Include other headers as-is
			attrs = append(attrs, slog.Any(key, values))
		}
	}

	return slog.GroupValue(attrs...)
}

// NewSafeHeaders creates a SafeHeaders wrapper for the given http.Header.
func NewSafeHeaders(headers http.Header) SafeHeaders {
	return SafeHeaders{headers: headers}
}

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

	// Use the read-only root URL for POST requests
	// Some Customers hide the cluster leader behind the firewall and only
	// allow access to the appliance through an load balancer that is only
	// accessible thru the appliance URL
	// This is a workaround to allow POST requests to work in this scenario

	url := fmt.Sprintf("%s/%s", c.getReadOnlyRootUrl(), path)
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

	result, err := c.sendHttpRequest(req)
	if err == nil {
		logger.Debug("POST request successful",
			"method", req.Method,
			"url", url,
			"responseBody", string(result),
		)
		return result, nil
	}

	logger.Debug("POST request failed on read-only URL, retrying on read-write URL",
		"url", url,
		"path", path,
	)

	url = fmt.Sprintf("%s/%s", c.getReadWriteRootUrl(), path)
	req, err = http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		logger.Error("Failed to create POST request on read-write URL",
			"error", err,
			"url", url,
		)
		return nil, fmt.Errorf("POST request failed on read-write URL: %v", err)
	}

	result, err = c.sendHttpRequest(req)
	if err == nil {
		logger.Debug("POST request successful on read-write URL",
			"method", req.Method,
			"url", url,
			"responseBody", string(result),
		)
		return result, nil
	}

	logger.Error("POST request failed on read-write URL",
		"url", url,
		"path", path,
	)
	return nil, fmt.Errorf("POST request failed on read-write URL: %v", err)
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
		"headers", NewSafeHeaders(req.Header),
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
		"headers", NewSafeHeaders(resp.Header),
	)

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted && resp.StatusCode != http.StatusCreated {
		logger.Error("Request failed with non-success status code",
			"method", req.Method,
			"url", req.URL,
			"status", resp.Status,
			"statusCode", resp.StatusCode,
			"responseHeaders", NewSafeHeaders(resp.Header),
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
	logger.Debug("Setting request headers", "existingHeaders", NewSafeHeaders(req.Header))

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
		"finalHeaders", NewSafeHeaders(req.Header),
		"method", req.Method,
		"url", req.URL,
	)
}

// getAuthorizationHeader prepares authorization headers for API requests.
// It formats the access token according to the required Bearer scheme.
//
// Parameters:
//   - req: The HTTP request to modify with authorization headers.
//
// Returns:
//   - *http.Request: The modified request with authorization headers.
func (c *SafeguardClient) getAuthorizationHeader() http.Header {
	headers := http.Header{}
	headers.Set("accept", "application/json")
	headers.Set("Authorization", "Bearer "+c.AccessToken.getUserToken())

	return headers
}
