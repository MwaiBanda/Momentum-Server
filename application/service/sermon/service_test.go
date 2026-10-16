package sermon

import (
	"context"
	"errors"
	"testing"
	"time"
)

type fakeProvider struct {
	body  []byte
	err   error
	pages []string
}

func (provider *fakeProvider) Fetch(_ context.Context, page string) ([]byte, error) {
	provider.pages = append(provider.pages, page)
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
	cacheStore.values["sermon-page-2"] = []byte(`{"current_page":2,"data":[{"id":"cached"}]}`)
	provider := &fakeProvider{body: []byte(`{"data":[{"id":"provider"}]}`)}
	service := GetInstance(provider, cacheStore)

	response, err := service.List(context.Background(), "2")
	if err != nil {
		t.Fatalf("List() error = %v", err)
	}
	if len(response.Data) != 1 || response.Data[0].ID != "cached" {
		t.Fatalf("response = %#v, want cached response", response)
	}
	if len(provider.pages) != 0 {
		t.Fatalf("provider calls = %d, want 0", len(provider.pages))
	}
}

func TestListFetchesAndCachesResponse(t *testing.T) {
	cacheStore := newFakeCache()
	provider := &fakeProvider{body: []byte(`{"current_page":3,"data":[{"id":"fresh"}]}`)}
	service := GetInstance(provider, cacheStore)

	response, err := service.List(context.Background(), "03")
	if err != nil {
		t.Fatalf("List() error = %v", err)
	}
	if len(response.Data) != 1 || response.Data[0].ID != "fresh" {
		t.Fatalf("response = %#v, want provider response", response)
	}
	if len(provider.pages) != 1 || provider.pages[0] != "3" {
		t.Fatalf("provider pages = %#v, want [3]", provider.pages)
	}
	if _, found := cacheStore.values["sermon-page-3"]; !found {
		t.Fatal("List() did not cache provider response")
	}
}

func TestListRejectsInvalidPage(t *testing.T) {
	provider := &fakeProvider{body: []byte(`{}`)}
	service := GetInstance(provider, nil)

	_, err := service.List(context.Background(), "not-a-page")
	if !errors.Is(err, ErrInvalidPage) {
		t.Fatalf("List() error = %v, want %v", err, ErrInvalidPage)
	}
	if len(provider.pages) != 0 {
		t.Fatal("provider was called for invalid page")
	}
}

func TestListReturnsProviderError(t *testing.T) {
	wantErr := errors.New("sermon provider unavailable")
	service := GetInstance(&fakeProvider{err: wantErr}, nil)

	_, err := service.List(context.Background(), "1")
	if !errors.Is(err, wantErr) {
		t.Fatalf("List() error = %v, want %v", err, wantErr)
	}
}
