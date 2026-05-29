package util

import (
	"io"
	"net/http"
	"time"
)

type HTTPClient struct {
	client  *http.Client
	baseURL string
}

func NewHTTPClient(baseURL string) *HTTPClient {
	return &HTTPClient{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL: baseURL,
	}
}

func (h *HTTPClient) Get(endpoint string) ([]byte, error) {
	url := h.baseURL + endpoint
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func (h *HTTPClient) FetchResource(resourceURL string) ([]byte, error) {
	resp, err := http.Get(resourceURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func (h *HTTPClient) DownloadFile(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func (h *HTTPClient) ProxyRequest(targetURL string) (*http.Response, error) {
	return http.Get(targetURL)
}

func (h *HTTPClient) FetchUserData(userID string, apiEndpoint string) ([]byte, error) {
	url := apiEndpoint + "/users/" + userID
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
