package httpclient

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestClientGetAppliesQueryAndReturnsBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Query().Get("page") != "2" {
			t.Fatalf("page = %q, want %q", request.URL.Query().Get("page"), "2")
		}
		_, _ = writer.Write([]byte(`{"ok":true}`))
	}))
	defer server.Close()

	body, err := New(server.Client()).Get(
		context.Background(),
		server.URL,
		url.Values{"page": {"2"}},
	)
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	if string(body) != `{"ok":true}` {
		t.Fatalf("Get() body = %q, want %q", body, `{"ok":true}`)
	}
}

func TestClientRejectsNonSuccessStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
		http.Error(writer, "upstream failed", http.StatusBadGateway)
	}))
	defer server.Close()

	_, err := New(server.Client()).Get(context.Background(), server.URL, nil)
	if err == nil || !strings.Contains(err.Error(), "502") {
		t.Fatalf("Get() error = %v, want status 502 error", err)
	}
}

func TestClientRejectsOversizedResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
		_, _ = writer.Write([]byte("123456"))
	}))
	defer server.Close()

	client := New(server.Client())
	client.maxBodyBytes = 5

	_, err := client.Get(context.Background(), server.URL, nil)
	if err == nil || !strings.Contains(err.Error(), "exceeds 5 bytes") {
		t.Fatalf("Get() error = %v, want response size error", err)
	}
}

func TestNewUsesDefaultTimeout(t *testing.T) {
	client := New(nil)
	if client.httpClient.Timeout != 10*time.Second {
		t.Fatalf("timeout = %v, want %v", client.httpClient.Timeout, 10*time.Second)
	}
}
