package event

import (
	"context"
	"errors"
	"testing"
	"time"
)

type fakeProvider struct {
	body  []byte
	err   error
	calls int
}

func (provider *fakeProvider) Fetch(context.Context) ([]byte, error) {
	provider.calls++
	return provider.body, provider.err
}

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

func TestListUsesCachedResponse(t *testing.T) {
	cacheStore := newFakeCache()
	cacheStore.values["events"] = []byte(`{"data":[{"id":"cached"}]}`)
	provider := &fakeProvider{body: []byte(`{"data":[{"id":"provider"}]}`)}
	service := GetInstance(provider, cacheStore)

	response, err := service.List(context.Background())
	if err != nil {
		t.Fatalf("List() error = %v", err)
	}
	if len(response.Data) != 1 || response.Data[0].ID != "cached" {
		t.Fatalf("response = %#v, want cached response", response)
	}
	if provider.calls != 0 {
		t.Fatalf("provider calls = %d, want 0", provider.calls)
	}
}

func TestListFetchesAndCachesResponse(t *testing.T) {
	cacheStore := newFakeCache()
	provider := &fakeProvider{body: []byte(`{"data":[{"id":"fresh"}]}`)}
	service := GetInstance(provider, cacheStore)

	response, err := service.List(context.Background())
	if err != nil {
		t.Fatalf("List() error = %v", err)
	}
	if len(response.Data) != 1 || response.Data[0].ID != "fresh" {
		t.Fatalf("response = %#v, want provider response", response)
	}
	if provider.calls != 1 {
		t.Fatalf("provider calls = %d, want 1", provider.calls)
	}
	if _, found := cacheStore.values["events"]; !found {
		t.Fatal("List() did not cache provider response")
	}
}

func TestListReturnsProviderError(t *testing.T) {
	wantErr := errors.New("planning center unavailable")
	service := GetInstance(&fakeProvider{err: wantErr}, nil)

	_, err := service.List(context.Background())
	if !errors.Is(err, wantErr) {
		t.Fatalf("List() error = %v, want %v", err, wantErr)
	}
}
