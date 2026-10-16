package planningcenter

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"Momentum/infrastructure/integration/httpclient"
)

type fakeCache struct {
	values map[string][]byte
}

func newFakeCache() *fakeCache {
	return &fakeCache{values: make(map[string][]byte)}
}

func (store *fakeCache) Get(_ context.Context, key string) ([]byte, bool, error) {
	value, found := store.values[key]
	return value, found, nil
}

func (store *fakeCache) Set(_ context.Context, key string, value []byte, _ time.Duration) error {
	store.values[key] = value
	return nil
}

func (store *fakeCache) Delete(_ context.Context, key string) error {
	delete(store.values, key)
	return nil
}

func TestClientFetchesTokenAndEventsAndCachesToken(t *testing.T) {
	tokenCalls := 0
	eventCalls := 0
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/token":
			tokenCalls++
			if request.Method != http.MethodPost {
				t.Fatalf("token method = %q, want POST", request.Method)
			}
			_, _ = writer.Write([]byte(`{"data":{"type":"Bearer","attributes":{"token":"secret"}}}`))
		case "/events":
			eventCalls++
			if request.Method != http.MethodGet {
				t.Fatalf("event method = %q, want GET", request.Method)
			}
			if request.Header.Get("Authorization") != "Bearer secret" {
				t.Fatalf("authorization = %q, want %q", request.Header.Get("Authorization"), "Bearer secret")
			}
			if request.URL.Query().Get("per_page") != "100" {
				t.Fatalf("per_page = %q, want %q", request.URL.Query().Get("per_page"), "100")
			}
			if request.URL.Query().Get("where[ends_at][gte]") == "" {
				t.Fatal("missing event date filter")
			}
			_, _ = writer.Write([]byte(`{"data":[]}`))
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	client := NewClient(
		httpclient.New(server.Client()),
		newFakeCache(),
		server.URL+"/events",
		server.URL+"/token",
	)

	for range 2 {
		if _, err := client.Fetch(context.Background()); err != nil {
			t.Fatalf("Fetch() error = %v", err)
		}
	}
	if tokenCalls != 1 {
		t.Fatalf("token calls = %d, want 1", tokenCalls)
	}
	if eventCalls != 2 {
		t.Fatalf("event calls = %d, want 2", eventCalls)
	}
}
