package planningcenter

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"Momentum/application/cache"
	"Momentum/constants"
	"Momentum/infrastructure/integration/httpclient"
)

const tokenTTL = 90 * time.Minute

type Client struct {
	transport  *httpclient.Client
	tokenStore cache.Store
	eventsURL  string
	tokenURL   string
}

type tokenResponse struct {
	Data token `json:"data"`
}

type token struct {
	Type       string          `json:"type"`
	Attributes tokenAttributes `json:"attributes"`
}

type tokenAttributes struct {
	Token string `json:"token"`
}

func NewClient(
	transport *httpclient.Client,
	tokenStore cache.Store,
	eventsURL string,
	tokenURL string,
) *Client {
	return &Client{
		transport:  transport,
		tokenStore: tokenStore,
		eventsURL:  eventsURL,
		tokenURL:   tokenURL,
	}
}

func (client *Client) Fetch(ctx context.Context) ([]byte, error) {
	if client == nil || client.transport == nil || client.eventsURL == "" || client.tokenURL == "" {
		return nil, fmt.Errorf("Planning Center client is not configured")
	}

	accessToken, tokenType, err := client.accessToken(ctx)
	if err != nil {
		return nil, err
	}

	requestURL, err := url.Parse(client.eventsURL)
	if err != nil {
		return nil, fmt.Errorf("parse Planning Center events URL: %w", err)
	}
	query := requestURL.Query()
	query.Set("per_page", "100")
	query.Set("where[ends_at][gte]", time.Now().Add(-4*time.Hour).Format(time.RFC3339))
	requestURL.RawQuery = query.Encode()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create Planning Center events request: %w", err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", tokenType+" "+accessToken)

	return client.transport.Do(request)
}

func (client *Client) accessToken(ctx context.Context) (string, string, error) {
	if client.tokenStore != nil {
		cached, found, err := cache.Get[tokenResponse](ctx, client.tokenStore, constants.TokenKey)
		if err != nil {
			log.Printf("[planning-center-cache] get %q: %v", constants.TokenKey, err)
		} else if found && cached.Data.Attributes.Token != "" {
			return cached.Data.Attributes.Token, cached.Data.Type, nil
		}
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, client.tokenURL, nil)
	if err != nil {
		return "", "", fmt.Errorf("create Planning Center token request: %w", err)
	}
	request.Header.Set("Content-Type", "application/json")

	body, err := client.transport.Do(request)
	if err != nil {
		return "", "", fmt.Errorf("request Planning Center token: %w", err)
	}

	var response tokenResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return "", "", fmt.Errorf("decode Planning Center token response: %w", err)
	}
	if response.Data.Attributes.Token == "" {
		return "", "", fmt.Errorf("Planning Center token response did not contain a token")
	}

	if client.tokenStore != nil {
		if err := cache.Set(ctx, client.tokenStore, constants.TokenKey, response, tokenTTL); err != nil {
			log.Printf("[planning-center-cache] set %q: %v", constants.TokenKey, err)
		}
	}

	return response.Data.Attributes.Token, response.Data.Type, nil
}
