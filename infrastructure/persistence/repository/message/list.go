package message

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	domainmessage "Momentum/domain/message"
	"Momentum/infrastructure/persistence"
)

type aggregateRow struct {
	MessageID        string         `db:"message_id"`
	Thumbnail        string         `db:"thumbnail"`
	Series           string         `db:"series"`
	Title            string         `db:"title"`
	Preacher         string         `db:"preacher"`
	Date             string         `db:"date"`
	MessageCreatedOn time.Time      `db:"message_created_on"`
	Published        bool           `db:"published"`
	HasOrder         bool           `db:"has_order"`
	PassageID        sql.NullString `db:"passage_id"`
	Header           sql.NullString `db:"header"`
	Verse            sql.NullString `db:"verse"`
	PassageMessage   sql.NullString `db:"passage_message"`
	PassageMessageID sql.NullString `db:"passage_message_id"`
	PassageCreatedOn sql.NullTime   `db:"passage_created_on"`
	PassageOrder     sql.NullInt64  `db:"passage_order"`
	NoteID           sql.NullString `db:"note_id"`
	NoteContent      sql.NullString `db:"note_content"`
	NoteUserID       sql.NullString `db:"note_user_id"`
	NotePassageID    sql.NullString `db:"note_passage_id"`
	NoteCreatedOn    sql.NullTime   `db:"note_created_on"`
	NoteUpdatedOn    sql.NullTime   `db:"note_updated_on"`
}

func (repository *Repository) ListPublishedByUser(ctx context.Context, userID string) ([]domainmessage.Message, error) {
	var rows []aggregateRow
	query := persistence.LoadFile("repository/message/sql/list_published_by_user.sql")
	if err := repository.db.SelectContext(ctx, &rows, query, userID); err != nil {
		return nil, fmt.Errorf("list published messages for user %q: %w", userID, err)
	}

	return aggregateRows(rows), nil
}

func (repository *Repository) ListUnpublished(ctx context.Context) ([]domainmessage.Message, error) {
	var rows []aggregateRow
	query := persistence.LoadFile("repository/message/sql/list_unpublished.sql")
	if err := repository.db.SelectContext(ctx, &rows, query); err != nil {
		return nil, fmt.Errorf("list unpublished messages: %w", err)
	}

	return aggregateRows(rows), nil
}

func aggregateRows(rows []aggregateRow) []domainmessage.Message {
	messages := make([]domainmessage.Message, 0)
	messageIndexes := make(map[string]int)
	passageIndexes := make(map[string]map[string]int)

	for _, row := range rows {
		messageIndex, exists := messageIndexes[row.MessageID]
		if !exists {
			messageIndex = len(messages)
			messageIndexes[row.MessageID] = messageIndex
			passageIndexes[row.MessageID] = make(map[string]int)
			messages = append(messages, domainmessage.Message{
				ID:        row.MessageID,
				Thumbnail: row.Thumbnail,
				Series:    row.Series,
				Title:     row.Title,
				Preacher:  row.Preacher,
				Date:      row.Date,
				CreatedOn: row.MessageCreatedOn,
				Published: row.Published,
				HasOrder:  row.HasOrder,
				Passages:  make([]domainmessage.Passage, 0),
			})
		}

		if !row.PassageID.Valid {
			continue
		}

		passageIndex, exists := passageIndexes[row.MessageID][row.PassageID.String]
		if !exists {
			passageIndex = len(messages[messageIndex].Passages)
			passageIndexes[row.MessageID][row.PassageID.String] = passageIndex
			messages[messageIndex].Passages = append(messages[messageIndex].Passages, domainmessage.Passage{
				ID:        row.PassageID.String,
				Header:    row.Header.String,
				Verse:     row.Verse.String,
				Message:   row.PassageMessage.String,
				MessageID: row.PassageMessageID.String,
				CreatedOn: nullTimeValue(row.PassageCreatedOn),
				Order:     int(nullInt64Value(row.PassageOrder)),
				Notes:     make([]domainmessage.Note, 0),
			})
		}

		if !row.NoteID.Valid {
			continue
		}

		messages[messageIndex].Passages[passageIndex].Notes = append(
			messages[messageIndex].Passages[passageIndex].Notes,
			domainmessage.Note{
				ID:        row.NoteID.String,
				Content:   row.NoteContent.String,
				UserID:    row.NoteUserID.String,
				PassageID: row.NotePassageID.String,
				CreatedOn: nullTimeValue(row.NoteCreatedOn),
				UpdatedOn: nullTimeValue(row.NoteUpdatedOn),
			},
		)
	}

	return messages
}
