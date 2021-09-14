package repository

import (
	"context"
	"time"

	"github.com/FotiadisM/workflow-server/internal/conversations"
)

func (r Repository) GetConversations(ctx context.Context, userID string) (convs []conversations.Conversation, err error) {
	rows, err := r.db.Query(ctx, `SELECT * FROM conversations WHERE user1_id=$1 OR user2_id=$1`, userID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id, user1ID, user2ID string
		if err = rows.Scan(&id, &user1ID, &user2ID); err != nil {
			return
		}

		switch userID {
		case user1ID:
			convs = append(convs, conversations.Conversation{ID: id, ConvUserID: user2ID})
		case user2ID:
			convs = append(convs, conversations.Conversation{ID: id, ConvUserID: user1ID})
		}
	}

	return
}

func (r Repository) CreateConversation(ctx context.Context, userID, convUserID string) (convID string, err error) {
	err = r.db.QueryRow(ctx, `INSERT INTO conversations (user1_id, user2_id) VALUES ($1, $2) RETURNING id`, userID, convUserID).Scan(&convID)
	return
}

func (r Repository) GetConversationMessages(ctx context.Context, convID string) (msgs []conversations.Message, err error) {
	rows, err := r.db.Query(ctx, `SELECT * FROM messages WHERE conv_id=$1 ORDER BY time_sent`, convID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		msg := conversations.Message{}
		if err = rows.Scan(&msg.ID, &msg.ConvID, &msg.SenterID, &msg.Text, &msg.Time); err != nil {
			return
		}
		msgs = append(msgs, msg)
	}

	return
}

func (r Repository) CreateConversationMessage(ctx context.Context, convID, senterID, text string) (msgID string, timeSent time.Time, err error) {
	err = r.db.QueryRow(ctx, `INSERT INTO messages (conv_id, senter_id, text) VALUES ($1, $2, $3) RETURNING id, time_sent`, convID, senterID, text).Scan(&msgID, &timeSent)

	return
}

func (r Repository) GetReceiver(ctx context.Context, convID, userID string) (receiverID string, err error) {
	var us1, us2 string
	err = r.db.QueryRow(ctx, `SELECT user1_id, user2_id FROM connections WHERE id=$1 AND (user1_id=$2 OR user2_id=$2)`, convID, userID).Scan(&us1, &us2)

	if us1 == userID {
		return us2, err
	}

	return us1, err
}
