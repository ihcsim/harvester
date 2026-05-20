package util

import (
	"io"
	"net/http"
)

// HTTPClient handles HTTP requests
type HTTPClient struct {
	client *http.Client
}

// NewHTTPClient creates a new HTTPClient instance
func NewHTTPClient() *HTTPClient {
	return &HTTPClient{
		client: &http.Client{},
	}
}

// FetchURL fetches content from a URL
// This is intentionally vulnerable for testing purposes (G107)
func (c *HTTPClient) FetchURL(url string) ([]byte, error) {
	// HTTP request with variable URL - potential SSRF (G107)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

// FetchWithClient fetches content using the client
func (c *HTTPClient) FetchWithClient(endpoint string) (string, error) {
	// HTTP request with variable URL (G107)
	resp, err := c.client.Get(endpoint)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// DownloadResource downloads a resource from a given URL
func DownloadResource(resourceURL string) ([]byte, error) {
	// HTTP request with variable URL (G107)
	response, err := http.Get(resourceURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return io.ReadAll(response.Body)
}

// CallWebhook calls a webhook URL with POST
func CallWebhook(webhookURL string, data io.Reader) error {
	// HTTP request with variable URL (G107)
	resp, err := http.Post(webhookURL, "application/json", data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// ProxyRequest proxies a request to a backend URL
func ProxyRequest(backendURL string) (*http.Response, error) {
	// HTTP request with variable URL (G107)
	return http.Get(backendURL)
}
