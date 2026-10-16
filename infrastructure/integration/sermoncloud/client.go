package sermoncloud

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"Momentum/infrastructure/integration/httpclient"
)

type Client struct {
	transport *httpclient.Client
	baseURL   string
}

func NewClient(httpClient *http.Client, baseURL string) *Client {
	return &Client{
		transport: httpclient.New(httpClient),
		baseURL:   strings.TrimRight(baseURL, "/"),
	}
}

func NewWithTransport(transport *httpclient.Client, baseURL string) *Client {
	return &Client{
		transport: transport,
		baseURL:   strings.TrimRight(baseURL, "/"),
	}
}

func (client *Client) Fetch(ctx context.Context, page string) ([]byte, error) {
	if client == nil || client.transport == nil || client.baseURL == "" {
		return nil, fmt.Errorf("SermonCloud client is not configured")
	}

	return client.transport.Get(ctx, client.baseURL, url.Values{"page": {page}})
}
