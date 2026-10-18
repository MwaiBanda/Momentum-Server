package message

import (
	"time"

	domainmessage "Momentum/domain/message"
)

type MessageListResponse struct {
	Data []MessageResponse `json:"data"`
}

type CreateInput struct {
	Thumbnail string               `json:"thumbnail"`
	Series    string               `json:"series"`
	Title     string               `json:"title"`
	Preacher  string               `json:"preacher"`
	Date      string               `json:"date"`
	Passages  []CreatePassageInput `json:"passages"`
}

type CreatePassageInput struct {
	Header  string `json:"header"`
	Verse   string `json:"verse"`
	Message string `json:"message"`
	Type    string `json:"type"`
	Order   int    `json:"order"`
}

type UpdateInput struct {
	ID        string               `json:"id"`
	Thumbnail string               `json:"thumbnail"`
	Series    string               `json:"series"`
	Title     string               `json:"title"`
	Preacher  string               `json:"preacher"`
	Date      string               `json:"date"`
	Published bool                 `json:"published"`
	Passages  []UpdatePassageInput `json:"passages"`
}

type UpdatePassageInput struct {
	ID      string `json:"id"`
	Header  string `json:"header"`
	Verse   string `json:"verse"`
	Message string `json:"message"`
}

type MessageResponse struct {
	ID        string            `json:"id"`
	Thumbnail string            `json:"thumbnail"`
	Series    string            `json:"series"`
	Title     string            `json:"title"`
	Preacher  string            `json:"preacher"`
	Date      string            `json:"date"`
	CreatedOn string            `json:"createdOn"`
	Published bool              `json:"published"`
	HasOrder  bool              `json:"hasOrder"`
	Passages  []PassageResponse `json:"passages"`
}

type PassageResponse struct {
	ID      string         `json:"id"`
	Header  string         `json:"header"`
	Verse   string         `json:"verse"`
	Message string         `json:"message"`
	Type    string         `json:"type"`
	Order   int            `json:"order"`
	Notes   []NoteResponse `json:"notes"`
}

type NoteResponse struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	UserID  string `json:"userId"`
}

type CreateNoteInput struct {
	PassageID string `json:"passageId"`
	UserID    string `json:"userId"`
	Content   string `json:"content"`
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
