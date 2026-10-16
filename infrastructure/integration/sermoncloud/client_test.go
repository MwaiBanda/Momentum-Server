package sermoncloud

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClientFetchBuildsPageRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/sermons" {
			t.Fatalf("path = %q, want %q", request.URL.Path, "/sermons")
		}
		if request.URL.Query().Get("page") != "2" {
			t.Fatalf("page = %q, want %q", request.URL.Query().Get("page"), "2")
		}
		_, _ = writer.Write([]byte(`{"data":[]}`))
	}))
	defer server.Close()

	body, err := NewClient(server.Client(), server.URL+"/sermons").Fetch(context.Background(), "2")
	if err != nil {
		t.Fatalf("Fetch() error = %v", err)
	}
	if string(body) != `{"data":[]}` {
		t.Fatalf("Fetch() body = %q, want %q", body, `{"data":[]}`)
	}
}

func TestClientReturnsUpstreamError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
		http.Error(writer, "upstream failed", http.StatusBadGateway)
	}))
	defer server.Close()

	_, err := NewClient(server.Client(), server.URL).Fetch(context.Background(), "1")
	if err == nil {
		t.Fatal("Fetch() error = nil, want upstream error")
	}
}
