package httpclient

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

const (
	defaultTimeout      = 10 * time.Second
	defaultMaxBodyBytes = 10 << 20
)

type Client struct {
	httpClient   *http.Client
	maxBodyBytes int64
}

func New(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: defaultTimeout}
	}
	clientCopy := *httpClient
	httpClient = &clientCopy
	transport := httpClient.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}
	httpClient.Transport = otelhttp.NewTransport(transport)

	return &Client{
		httpClient:   httpClient,
		maxBodyBytes: defaultMaxBodyBytes,
	}
}

func (client *Client) Get(ctx context.Context, endpoint string, query url.Values) ([]byte, error) {
	if client == nil || client.httpClient == nil {
		return nil, fmt.Errorf("outbound HTTP client is not configured")
	}

	requestURL, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("parse outbound URL: %w", err)
	}
	if query != nil {
		values := requestURL.Query()
		for key, entries := range query {
			values[key] = append([]string(nil), entries...)
		}
		requestURL.RawQuery = values.Encode()
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create outbound request: %w", err)
	}

	return client.Do(request)
}

func (client *Client) Do(request *http.Request) ([]byte, error) {
	if client == nil || client.httpClient == nil {
		return nil, fmt.Errorf("outbound HTTP client is not configured")
	}
	if request == nil {
		return nil, fmt.Errorf("outbound request is nil")
	}

	response, err := client.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("execute outbound request: %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(io.LimitReader(response.Body, client.maxBodyBytes+1))
	if err != nil {
		return nil, fmt.Errorf("read outbound response: %w", err)
	}
	if int64(len(body)) > client.maxBodyBytes {
		return nil, fmt.Errorf("outbound response exceeds %d bytes", client.maxBodyBytes)
	}
	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("outbound request returned status %d: %s", response.StatusCode, body)
	}

	return body, nil
}
