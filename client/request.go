package client

import (
	"fmt"
	"io"
	"net/http"
)

// GetRootUrl constructs and returns the root URL for the Safeguard API.
func (c *SafeguardClient) GetRootUrl() string {
	return fmt.Sprintf("%s/service/core/%s", c.ApplicanceURL, c.ApiVersion)
}

// GetRequest makes a GET request to the specified path on the Safeguard API.
// Parameters:
// - path: The API endpoint path.
// Returns the response body as a byte slice and an error if the request fails.
func (c *SafeguardClient) GetRequest(path string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", c.GetRootUrl(), path)
	log.Debugf("Making request to: %s", url)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return c.sendHttpRequest(req)
}

func (c *SafeguardClient) PostRequest(path string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", c.GetRootUrl(), path)
	log.Debugf("Making request to: %s", url)

	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}

	return c.sendHttpRequest(req)
}

func (c *SafeguardClient) PutRequest(path string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", c.GetRootUrl(), path)
	log.Debugf("Making request to: %s", url)

	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		return nil, err
	}

	return c.sendHttpRequest(req)
}

func (c *SafeguardClient) DeleteRequest(path string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", c.GetRootUrl(), path)
	log.Debugf("Making request to: %s", url)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	return c.sendHttpRequest(req)
}

func (c *SafeguardClient) sendHttpRequest(req *http.Request) ([]byte, error) {
	req = c.getAuthorizationHeader(req)
	log.Debugf("Request headers: %+v", req.Header)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
		log.Debugf("Response Status: %s", resp.Status)
		log.Debugf("Response headers: %+v", resp.Header)
		log.Debugf("Response body: %s", string(body))
		return nil, fmt.Errorf("error retrieving me: HTTP %d", resp.StatusCode)
	}

	return body, nil
}
