package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type HTTPClient struct {
	BaseURL  string
	Timeout  int
	APIKey   string
	Endpoint string
	Method   string
}

func buildURL(base, endpoint string) string {
	base = strings.TrimSuffix(base, "/")
	endpoint = strings.TrimPrefix(endpoint, "/")
	return base + "/" + endpoint
}

func InitHTTPClient(baseUrl string, timeout int, apiKey string) HTTPClient {
	return HTTPClient{
		BaseURL: baseUrl,
		Timeout: timeout,
		APIKey:  apiKey,
	}
}

func (h *HTTPClient) SendJSON(ctx context.Context, endpoint, method string, payload map[string]any) (string, error) {
	var requestBody map[string]any
	var requestPayload *bytes.Buffer

	if payload != nil {
		requestBody = payload
		jsonData, err := json.Marshal(requestBody)
		if err != nil {
			return "", err
		}
		requestPayload = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequestWithContext(ctx, method, buildURL(h.BaseURL, endpoint), requestPayload)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Key: "+h.APIKey)

	client := &http.Client{Timeout: time.Duration(h.Timeout) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	isFailed := resp.StatusCode < 200 || resp.StatusCode >= 300
	if isFailed {
		return string(bodyBytes), fmt.Errorf("request failed: %s", resp.Status)
	}

	return string(bodyBytes), nil
}

func (h *HTTPClient) SendFormEncoded(ctx context.Context, endpoint, method string, payload map[string]string) (string, error) {
	var requestBody map[string]string
	var requestPayload *bytes.Buffer

	if payload != nil {
		requestBody = payload
		values := url.Values{}

		for k, v := range requestBody {
			values.Set(k, v)
		}

		requestPayload = bytes.NewBufferString(values.Encode())
	}

	req, err := http.NewRequestWithContext(ctx, method, buildURL(h.BaseURL, endpoint), requestPayload)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Key", h.APIKey)

	client := &http.Client{Timeout: time.Duration(h.Timeout) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	isFailed := resp.StatusCode < 200 || resp.StatusCode >= 300
	if isFailed {
		return string(bodyBytes), fmt.Errorf("request failed: %s", resp.Status)
	}

	return string(bodyBytes), nil
}
