package message

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"testing"
	"time"

	"Momentum/application/cache"
	domainmessage "Momentum/domain/message"
)

type fakeRepository struct {
	published       []domainmessage.Message
	unpublished     []domainmessage.Message
	publishedErr    error
	unpublishedErr  error
	createParams    domainmessage.CreateParams
	createErr       error
	updateParams    domainmessage.UpdateParams
	updateErr       error
	deleteID        string
	deleteErr       error
	createNoteParam domainmessage.CreateNoteParams
	createNoteErr   error
	updateNoteParam domainmessage.UpdateNoteParams
	updateNote      domainmessage.Note
	updateNoteErr   error
	deleteNoteID    string
	deleteNote      domainmessage.Note
	deleteNoteErr   error
}

func (repository *fakeRepository) ListPublishedByUser(context.Context, string) ([]domainmessage.Message, error) {
	return repository.published, repository.publishedErr
}

func (repository *fakeRepository) ListUnpublished(context.Context) ([]domainmessage.Message, error) {
	return repository.unpublished, repository.unpublishedErr
}

func (repository *fakeRepository) Create(_ context.Context, params domainmessage.CreateParams) error {
	repository.createParams = params
	return repository.createErr
}

func (repository *fakeRepository) Update(_ context.Context, params domainmessage.UpdateParams) error {
	repository.updateParams = params
	return repository.updateErr
}

func (repository *fakeRepository) Delete(_ context.Context, id string) error {
	repository.deleteID = id
	return repository.deleteErr
}

func (repository *fakeRepository) CreateNote(_ context.Context, params domainmessage.CreateNoteParams) error {
	repository.createNoteParam = params
	return repository.createNoteErr
}

func (repository *fakeRepository) UpdateNote(_ context.Context, params domainmessage.UpdateNoteParams) (domainmessage.Note, error) {
	repository.updateNoteParam = params
	return repository.updateNote, repository.updateNoteErr
}

func (repository *fakeRepository) DeleteNote(_ context.Context, id string) (domainmessage.Note, error) {
	repository.deleteNoteID = id
	return repository.deleteNote, repository.deleteNoteErr
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

func (store *fakeCache) Increment(_ context.Context, key string) (int64, error) {
	current := int64(0)
	if value, found := store.values[key]; found {
		parsed, err := strconv.ParseInt(string(value), 10, 64)
		if err != nil {
			return 0, err
		}
		current = parsed
	}
	current++
	store.values[key] = []byte(strconv.FormatInt(current, 10))
	return current, nil
}

func TestCreatePreservesLegacyResponseAndGeneratesPersistenceIDs(t *testing.T) {
	repository := &fakeRepository{}
	service := GetInstance(repository, nil)
	input := CreateInput{
		Thumbnail: "thumbnail.png",
		Series:    "Series",
		Title:     "Message",
		Preacher:  "Preacher",
		Date:      "2026-10-16",
		Passages: []CreatePassageInput{
			{Type: "header", Header: "Introduction", Order: 1},
			{Type: "message", Verse: "John 1:1", Message: "In the beginning", Order: 2},
			{Type: "ignored", Message: "not persisted", Order: 3},
		},
	}

	response, err := service.Create(context.Background(), input)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	if response.ID == "" {
		t.Fatal("Create() response has no generated message ID")
	}
	if response.Title != input.Title || response.Preacher != input.Preacher {
		t.Fatalf("Create() response = %#v, want fields from request %#v", response, input)
	}
	if repository.createParams.ID == "" {
		t.Fatal("Create() did not generate a message ID")
	}
	if len(repository.createParams.Passages) != 2 {
		t.Fatalf("persisted passages = %d, want 2 valid passages", len(repository.createParams.Passages))
	}
	for index, passage := range repository.createParams.Passages {
		if passage.ID == "" {
			t.Fatalf("persisted passage %d has no generated ID", index)
		}
	}
	if repository.createParams.Passages[0].Header == nil || repository.createParams.Passages[1].Message == nil {
		t.Fatal("Create() did not map passage types to their persistence fields")
	}
}

func TestListPublishedByUserReadsCache(t *testing.T) {
	cached := []MessageResponse{{ID: "message-1", Title: "Cached"}}
	encoded, err := json.Marshal(cached)
	if err != nil {
		t.Fatalf("marshal cache: %v", err)
	}

	cacheStore := newFakeCache()
	cacheStore.values["messages:version"] = []byte("0")
	cacheStore.values["messages-v0-user-1"] = encoded
	service := GetInstance(&fakeRepository{
		published: []domainmessage.Message{{ID: "database-message"}},
	}, cacheStore)

	response, err := service.ListPublishedByUser(context.Background(), "user-1")
	if err != nil {
		t.Fatalf("ListPublishedByUser() error = %v", err)
	}
	if !reflect.DeepEqual(response.Data, cached) {
		t.Fatalf("response = %#v, want cached %#v", response.Data, cached)
	}
}

func TestListPublishedByUserMapsAggregateAndCaches(t *testing.T) {
	createdOn := time.Date(2026, time.October, 16, 0, 0, 0, 0, time.UTC)
	service := GetInstance(&fakeRepository{
		published: []domainmessage.Message{{
			ID:        "message-1",
			Title:     "Message",
			CreatedOn: createdOn,
			Passages: []domainmessage.Passage{{
				ID:     "passage-1",
				Header: "Introduction",
				Order:  1,
				Notes:  []domainmessage.Note{{ID: "note-1", Content: "Remember this", UserID: "user-1"}},
			}},
		}},
	}, newFakeCache())

	response, err := service.ListPublishedByUser(context.Background(), "user-1")
	if err != nil {
		t.Fatalf("ListPublishedByUser() error = %v", err)
	}
	if len(response.Data) != 1 || len(response.Data[0].Passages) != 1 || len(response.Data[0].Passages[0].Notes) != 1 {
		t.Fatalf("response = %#v, want hydrated message aggregate", response)
	}
	if response.Data[0].CreatedOn != createdOn.Format(time.RFC3339Nano) {
		t.Fatalf("CreatedOn = %q, want %q", response.Data[0].CreatedOn, createdOn.Format(time.RFC3339Nano))
	}
}

func TestUpdateMapsNotFound(t *testing.T) {
	repository := &fakeRepository{updateErr: domainmessage.ErrMessageNotFound}
	service := GetInstance(repository, nil)

	_, err := service.Update(context.Background(), UpdateInput{ID: "message-1"})
	if !errors.Is(err, ErrMessageNotFound) {
		t.Fatalf("Update() error = %v, want %v", err, ErrMessageNotFound)
	}
}

func TestCreateNoteInvalidatesUserCache(t *testing.T) {
	cacheStore := newFakeCache()
	cacheStore.values["messages:version"] = []byte("0")
	repository := &fakeRepository{}
	service := GetInstance(repository, cacheStore)

	input := CreateNoteInput{PassageID: "passage-1", UserID: "user-1", Content: "Note"}
	response, err := service.CreateNote(context.Background(), input)
	if err != nil {
		t.Fatalf("CreateNote() error = %v", err)
	}
	if !reflect.DeepEqual(response, input) {
		t.Fatalf("CreateNote() response = %#v, want %#v", response, input)
	}
	if string(cacheStore.values["messages:version"]) != "1" {
		t.Fatalf("message cache version = %q, want %q", cacheStore.values["messages:version"], "1")
	}
}

func TestMessageCacheVersionMakesOlderEntriesUnreachable(t *testing.T) {
	cacheStore := newFakeCache()
	cacheStore.values["messages:version"] = []byte("0")
	cacheStore.values["messages-v0-user-1"] = []byte(`[{"id":"stale"}]`)
	repository := &fakeRepository{
		published: []domainmessage.Message{{ID: "fresh"}},
	}
	service := GetInstance(repository, cacheStore)
	service.bumpMessageCacheVersion(context.Background())

	response, err := service.ListPublishedByUser(context.Background(), "user-1")
	if err != nil {
		t.Fatalf("ListPublishedByUser() error = %v", err)
	}
	if len(response.Data) != 1 || response.Data[0].ID != "fresh" {
		t.Fatalf("response = %#v, want fresh data after version bump", response.Data)
	}
}

func TestDeleteReturnsLegacyEmptyResponse(t *testing.T) {
	repository := &fakeRepository{}
	service := GetInstance(repository, nil)

	response, err := service.Delete(context.Background(), "message-1")
	if err != nil {
		t.Fatalf("Delete() error = %v", err)
	}
	if !reflect.DeepEqual(response, MessageResponse{}) {
		t.Fatalf("Delete() response = %#v, want legacy empty response", response)
	}
	if repository.deleteID != "message-1" {
		t.Fatalf("Delete() ID = %q, want %q", repository.deleteID, "message-1")
	}
}

var _ cache.Store = (*fakeCache)(nil)
var _ cache.CounterStore = (*fakeCache)(nil)
