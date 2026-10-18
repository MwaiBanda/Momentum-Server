package message

import (
	"time"

	domainmessage "Momentum/domain/message"
)

type MessageListResponse struct {
	Data []MessageResponse `json:"data" binding:"required"`
}

type CreateInput struct {
	Thumbnail string               `json:"thumbnail" binding:"required"`
	Series    string               `json:"series" binding:"required"`
	Title     string               `json:"title" binding:"required"`
	Preacher  string               `json:"preacher" binding:"required"`
	Date      string               `json:"date" binding:"required"`
	Passages  []CreatePassageInput `json:"passages" binding:"required"`
}

type CreatePassageInput struct {
	Header  string `json:"header"`
	Verse   string `json:"verse"`
	Message string `json:"message"`
	Type    string `json:"type" binding:"required"`
	Order   int    `json:"order" binding:"required"`
}

type UpdateInput struct {
	ID        string               `json:"id" binding:"required"`
	Thumbnail string               `json:"thumbnail" binding:"required"`
	Series    string               `json:"series" binding:"required"`
	Title     string               `json:"title" binding:"required"`
	Preacher  string               `json:"preacher" binding:"required"`
	Date      string               `json:"date" binding:"required"`
	Published bool                 `json:"published" binding:"required"`
	Passages  []UpdatePassageInput `json:"passages" binding:"required"`
}

type UpdatePassageInput struct {
	ID      string `json:"id" binding:"required"`
	Header  string `json:"header"`
	Verse   string `json:"verse"`
	Message string `json:"message"`
}

type MessageResponse struct {
	ID        string            `json:"id" binding:"required"`
	Thumbnail string            `json:"thumbnail" binding:"required"`
	Series    string            `json:"series" binding:"required"`
	Title     string            `json:"title" binding:"required"`
	Preacher  string            `json:"preacher" binding:"required"`
	Date      string            `json:"date" binding:"required"`
	CreatedOn string            `json:"createdOn" binding:"required"`
	Published bool              `json:"published" binding:"required"`
	HasOrder  bool              `json:"hasOrder" binding:"required"`
	Passages  []PassageResponse `json:"passages" binding:"required"`
}

type PassageResponse struct {
	ID      string         `json:"id" binding:"required"`
	Header  string         `json:"header" binding:"required"`
	Verse   string         `json:"verse" binding:"required"`
	Message string         `json:"message" binding:"required"`
	Type    string         `json:"type" binding:"required"`
	Order   int            `json:"order" binding:"required"`
	Notes   []NoteResponse `json:"notes" binding:"required"`
}

type NoteResponse struct {
	ID      string `json:"id" binding:"required"`
	Content string `json:"content" binding:"required"`
	UserID  string `json:"userId" binding:"required"`
}

type UpdateNoteInput struct {
	ID      string `json:"id" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type CreateNoteInput struct {
	PassageID string `json:"passageId" binding:"required"`
	UserID    string `json:"userId" binding:"required"`
	Content   string `json:"content" binding:"required"`
}

type CreateNoteResponse struct {
	ID      string `json:"id" binding:"required"`
	Content string `json:"content" binding:"required"`
	UserID  string `json:"userId" binding:"required"`
}

func messageResponseFromCreateInput(input CreateInput, id string) MessageResponse {
	response := MessageResponse{
		ID:        id,
		Thumbnail: input.Thumbnail,
		Series:    input.Series,
		Title:     input.Title,
		Preacher:  input.Preacher,
		Date:      input.Date,
		Passages:  make([]PassageResponse, 0, len(input.Passages)),
	}
	for _, passage := range input.Passages {
		response.Passages = append(response.Passages, PassageResponse{
			Header:  passage.Header,
			Verse:   passage.Verse,
			Message: passage.Message,
			Type:    passage.Type,
			Order:   passage.Order,
		})
	}
	return response
}

func messageResponseFromUpdateInput(input UpdateInput) MessageResponse {
	response := MessageResponse{
		ID:        input.ID,
		Thumbnail: input.Thumbnail,
		Series:    input.Series,
		Title:     input.Title,
		Preacher:  input.Preacher,
		Date:      input.Date,
		Published: input.Published,
		Passages:  make([]PassageResponse, 0, len(input.Passages)),
	}
	for _, passage := range input.Passages {
		response.Passages = append(response.Passages, PassageResponse{
			ID:      passage.ID,
			Header:  passage.Header,
			Verse:   passage.Verse,
			Message: passage.Message,
		})
	}
	return response
}

func toMessageResponse(messageEntity domainmessage.Message) MessageResponse {
	passages := make([]PassageResponse, 0, len(messageEntity.Passages))
	for _, passage := range messageEntity.Passages {
		notes := make([]NoteResponse, 0, len(passage.Notes))
		for _, note := range passage.Notes {
			notes = append(notes, NoteResponse{
				ID:      note.ID,
				Content: note.Content,
				UserID:  note.UserID,
			})
		}

		passages = append(passages, PassageResponse{
			ID:      passage.ID,
			Header:  passage.Header,
			Verse:   passage.Verse,
			Message: passage.Message,
			Order:   passage.Order,
			Notes:   notes,
		})
	}

	return MessageResponse{
		ID:        messageEntity.ID,
		Thumbnail: messageEntity.Thumbnail,
		Series:    messageEntity.Series,
		Title:     messageEntity.Title,
		Preacher:  messageEntity.Preacher,
		Date:      messageEntity.Date,
		CreatedOn: messageEntity.CreatedOn.Format(time.RFC3339Nano),
		Published: messageEntity.Published,
		HasOrder:  messageEntity.HasOrder,
		Passages:  passages,
	}
}

func toMessageListResponse(messages []domainmessage.Message) MessageListResponse {
	response := MessageListResponse{
		Data: make([]MessageResponse, 0, len(messages)),
	}
	for _, messageEntity := range messages {
		response.Data = append(response.Data, toMessageResponse(messageEntity))
	}
	return response
}

func toNoteResponse(note domainmessage.Note) NoteResponse {
	return NoteResponse{
		ID:      note.ID,
		Content: note.Content,
		UserID:  note.UserID,
	}
}
