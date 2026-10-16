package message

import (
	"time"

	domainmessage "Momentum/domain/message"
)

type MessageListResponse struct {
	Data []MessageResponse `json:"data"`
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

type MessageInput = MessageResponse

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
