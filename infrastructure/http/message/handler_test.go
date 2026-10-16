package message

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	applicationmessage "Momentum/application/service/message"

	"github.com/gofiber/fiber/v3"
)

type fakeUseCase struct {
	listPublished   func(context.Context, string) (applicationmessage.MessageListResponse, error)
	listUnpublished func(context.Context) (applicationmessage.MessageListResponse, error)
	create          func(context.Context, applicationmessage.MessageInput) (applicationmessage.MessageResponse, error)
	update          func(context.Context, applicationmessage.MessageInput) (applicationmessage.MessageResponse, error)
	delete          func(context.Context, string) (applicationmessage.MessageResponse, error)
	createNote      func(context.Context, applicationmessage.CreateNoteInput) (applicationmessage.CreateNoteInput, error)
	updateNote      func(context.Context, applicationmessage.NoteResponse) (applicationmessage.NoteResponse, error)
	deleteNote      func(context.Context, string) (applicationmessage.NoteResponse, error)
}

func (fake *fakeUseCase) ListPublishedByUser(ctx context.Context, userID string) (applicationmessage.MessageListResponse, error) {
	if fake.listPublished == nil {
		return applicationmessage.MessageListResponse{}, nil
	}
	return fake.listPublished(ctx, userID)
}

func (fake *fakeUseCase) ListUnpublished(ctx context.Context) (applicationmessage.MessageListResponse, error) {
	if fake.listUnpublished == nil {
		return applicationmessage.MessageListResponse{}, nil
	}
	return fake.listUnpublished(ctx)
}

func (fake *fakeUseCase) Create(ctx context.Context, input applicationmessage.MessageInput) (applicationmessage.MessageResponse, error) {
	if fake.create == nil {
		return applicationmessage.MessageResponse{}, nil
	}
	return fake.create(ctx, input)
}

func (fake *fakeUseCase) Update(ctx context.Context, input applicationmessage.MessageInput) (applicationmessage.MessageResponse, error) {
	if fake.update == nil {
		return applicationmessage.MessageResponse{}, nil
	}
	return fake.update(ctx, input)
}

func (fake *fakeUseCase) Delete(ctx context.Context, id string) (applicationmessage.MessageResponse, error) {
	if fake.delete == nil {
		return applicationmessage.MessageResponse{}, nil
	}
	return fake.delete(ctx, id)
}

func (fake *fakeUseCase) CreateNote(ctx context.Context, input applicationmessage.CreateNoteInput) (applicationmessage.CreateNoteInput, error) {
	if fake.createNote == nil {
		return applicationmessage.CreateNoteInput{}, nil
	}
	return fake.createNote(ctx, input)
}

func (fake *fakeUseCase) UpdateNote(ctx context.Context, input applicationmessage.NoteResponse) (applicationmessage.NoteResponse, error) {
	if fake.updateNote == nil {
		return applicationmessage.NoteResponse{}, nil
	}
	return fake.updateNote(ctx, input)
}

func (fake *fakeUseCase) DeleteNote(ctx context.Context, id string) (applicationmessage.NoteResponse, error) {
	if fake.deleteNote == nil {
		return applicationmessage.NoteResponse{}, nil
	}
	return fake.deleteNote(ctx, id)
}

func newTestApp(messages applicationmessage.UseCase) *fiber.App {
	handler := GetInstance(messages)
	app := fiber.New()
	app.Get("/messages/admin", handler.ListAdmin)
	app.Get("/messages/unpublished", handler.ListUnpublished)
	app.Get("/messages/:userId", handler.ListPublishedByUser)
	app.Post("/messages", handler.Create)
	app.Put("/messages", handler.Update)
	app.Delete("/messages/:messageId", handler.Delete)
	app.Post("/messages/notes", handler.CreateNote)
	app.Put("/messages/notes", handler.UpdateNote)
	app.Delete("/messages/notes/:noteId", handler.DeleteNote)
	return app
}

func performRequest(t *testing.T, app *fiber.App, method string, target string, body string) *http.Response {
	t.Helper()

	request := httptest.NewRequest(method, target, strings.NewReader(body))
	if body != "" {
		request.Header.Set("Content-Type", "application/json")
	}

	response, err := app.Test(request)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	t.Cleanup(func() {
		response.Body.Close()
	})
	return response
}

func TestListPublishedByUserReturnsResponse(t *testing.T) {
	expected := applicationmessage.MessageListResponse{
		Data: []applicationmessage.MessageResponse{{ID: "message-1", Title: "Message"}},
	}
	var receivedUserID string
	messages := &fakeUseCase{
		listPublished: func(_ context.Context, userID string) (applicationmessage.MessageListResponse, error) {
			receivedUserID = userID
			return expected, nil
		},
	}

	response := performRequest(t, newTestApp(messages), http.MethodGet, "/messages/user-1", "")
	if response.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.StatusCode, http.StatusOK)
	}
	if receivedUserID != "user-1" {
		t.Fatalf("user ID = %q, want %q", receivedUserID, "user-1")
	}

	var actual applicationmessage.MessageListResponse
	if err := json.NewDecoder(response.Body).Decode(&actual); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("response = %#v, want %#v", actual, expected)
	}
}

func TestListAdminReturnsPublishedMessageContract(t *testing.T) {
	expected := applicationmessage.MessageListResponse{
		Data: []applicationmessage.MessageResponse{{ID: "message-1", Title: "Message", Published: true}},
	}
	var receivedUserID string
	messages := &fakeUseCase{
		listPublished: func(_ context.Context, userID string) (applicationmessage.MessageListResponse, error) {
			receivedUserID = userID
			return expected, nil
		},
	}

	response := performRequest(t, newTestApp(messages), http.MethodGet, "/messages/admin", "")
	if response.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.StatusCode, http.StatusOK)
	}
	if receivedUserID != "admin" {
		t.Fatalf("admin user ID = %q, want %q", receivedUserID, "admin")
	}

	var actual applicationmessage.MessageListResponse
	if err := json.NewDecoder(response.Body).Decode(&actual); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("response = %#v, want %#v", actual, expected)
	}
}

func TestListUnpublishedReturnsMessageListContract(t *testing.T) {
	expected := applicationmessage.MessageListResponse{
		Data: []applicationmessage.MessageResponse{{ID: "draft-1", Title: "Draft", Published: false}},
	}
	messages := &fakeUseCase{
		listUnpublished: func(context.Context) (applicationmessage.MessageListResponse, error) {
			return expected, nil
		},
	}

	response := performRequest(t, newTestApp(messages), http.MethodGet, "/messages/unpublished", "")
	if response.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.StatusCode, http.StatusOK)
	}

	var actual applicationmessage.MessageListResponse
	if err := json.NewDecoder(response.Body).Decode(&actual); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("response = %#v, want %#v", actual, expected)
	}
}

func TestCreatePreservesLegacyResponse(t *testing.T) {
	input := applicationmessage.MessageInput{
		Title:    "Message",
		Preacher: "Preacher",
		Passages: []applicationmessage.PassageResponse{{Type: "header", Header: "Intro", Order: 1}},
	}
	var received applicationmessage.MessageInput
	messages := &fakeUseCase{
		create: func(_ context.Context, actual applicationmessage.MessageInput) (applicationmessage.MessageResponse, error) {
			received = actual
			return actual, nil
		},
	}

	response := performRequest(
		t,
		newTestApp(messages),
		http.MethodPost,
		"/messages",
		`{"title":"Message","preacher":"Preacher","passages":[{"type":"header","header":"Intro","order":1}]}`,
	)
	if response.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.StatusCode, http.StatusOK)
	}
	if !reflect.DeepEqual(received, input) {
		t.Fatalf("input = %#v, want %#v", received, input)
	}
}

func TestCreateRejectsMalformedBody(t *testing.T) {
	called := false
	messages := &fakeUseCase{
		create: func(context.Context, applicationmessage.MessageInput) (applicationmessage.MessageResponse, error) {
			called = true
			return applicationmessage.MessageResponse{}, nil
		},
	}

	response := performRequest(t, newTestApp(messages), http.MethodPost, "/messages", `{invalid`)
	if response.StatusCode != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", response.StatusCode, http.StatusBadRequest)
	}
	if called {
		t.Fatal("Create use case was called for malformed input")
	}
}

func TestDeleteMapsErrors(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want int
	}{
		{name: "invalid input", err: applicationmessage.ErrInvalidInput, want: http.StatusBadRequest},
		{name: "not found", err: applicationmessage.ErrMessageNotFound, want: http.StatusNotFound},
		{name: "internal error", err: errors.New("database unavailable"), want: http.StatusInternalServerError},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			messages := &fakeUseCase{
				delete: func(context.Context, string) (applicationmessage.MessageResponse, error) {
					return applicationmessage.MessageResponse{}, test.err
				},
			}
			response := performRequest(t, newTestApp(messages), http.MethodDelete, "/messages/message-1", "")
			if response.StatusCode != test.want {
				t.Fatalf("status = %d, want %d", response.StatusCode, test.want)
			}
		})
	}
}

func TestDeleteNoteReturnsResponse(t *testing.T) {
	expected := applicationmessage.NoteResponse{ID: "note-1", Content: "Deleted", UserID: "user-1"}
	var receivedID string
	messages := &fakeUseCase{
		deleteNote: func(_ context.Context, id string) (applicationmessage.NoteResponse, error) {
			receivedID = id
			return expected, nil
		},
	}

	response := performRequest(t, newTestApp(messages), http.MethodDelete, "/messages/notes/note-1", "")
	if response.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.StatusCode, http.StatusOK)
	}
	if receivedID != "note-1" {
		t.Fatalf("note ID = %q, want %q", receivedID, "note-1")
	}
}
