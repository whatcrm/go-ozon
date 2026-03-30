package models

import (
	"bytes"
	"encoding/json"
	"time"
)

type SendFileRequest struct {
	Base64Content string `json:"base64_content"`
	ChatId        string `json:"chat_id"`
	Name          string `json:"name"`
}

type SendMessageRequest struct {
	ChatId string `json:"chat_id"`
	Text   string `json:"text"`
}

type StartChatRequest struct {
	PostingNumber string `json:"posting_number"`
}

type ReadChatRequest struct {
	ChatId        string `json:"chat_id"`
	FromMessageId int64  `json:"from_message_id"`
}

type ChatFilter struct {
	ChatStatus string `json:"chat_status"`
	UnreadOnly bool   `json:"unread_only"`
}

type ChatListRequest struct {
	Filter ChatFilter `json:"filter"`
	Limit  int        `json:"limit"`
	Cursor string     `json:"cursor"`
}

// MessageID can be returned by API either as a JSON number (e.g. 0)
// or as a JSON string (e.g. "3000000407272079927").
type MessageID string

func (id *MessageID) UnmarshalJSON(b []byte) error {
	b = bytes.TrimSpace(b)
	if len(b) == 0 || bytes.Equal(b, []byte("null")) {
		*id = ""
		return nil
	}

	// "123"
	if b[0] == '"' {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		*id = MessageID(s)
		return nil
	}

	// 123
	var n json.Number
	if err := json.Unmarshal(b, &n); err != nil {
		return err
	}
	*id = MessageID(n.String())
	return nil
}

type ChatListItem struct {
	FirstUnreadMessageID MessageID `json:"first_unread_message_id"`
	LastMessageID        MessageID `json:"last_message_id"`
	UnreadCount          int       `json:"unread_count"`
	Chat                 struct {
		ChatID     string    `json:"chat_id"`
		ChatStatus string    `json:"chat_status"`
		ChatType   string    `json:"chat_type"`
		CreatedAt  time.Time `json:"created_at"`
	} `json:"chat"`
}

type ChatListResponse struct {
	Chats            []ChatListItem `json:"chats"`
	Cursor           string         `json:"cursor"`
	HasNext          bool           `json:"has_next"`
	TotalUnreadCount int            `json:"total_unread_count"`
}

type ChatHistoryRequest struct {
	ChatId    string `json:"chat_id"`
	Direction string `json:"direction"`
	Filter    struct {
		MessageIds []string `json:"message_ids"`
	} `json:"filter"`
	FromMessageId int64 `json:"from_message_id"`
	Limit         int   `json:"limit"`
}

type ChatHistoryResponse struct {
	HasNext  bool `json:"has_next"`
	Messages []struct {
		Context struct {
			OrderNumber string `json:"order_number"`
			Sku         string `json:"sku"`
		} `json:"context"`
		CreatedAt           time.Time `json:"created_at"`
		Data                []string  `json:"data"`
		IsImage             bool      `json:"is_image"`
		IsRead              bool      `json:"is_read"`
		MessageId           int64     `json:"message_id"`
		ModerateImageStatus string    `json:"moderate_image_status"`
		User                struct {
			Id   string `json:"id"`
			Type string `json:"type"`
		} `json:"user"`
	} `json:"messages"`
}
type T struct {
	HasNext  bool `json:"has_next"`
	Messages []struct {
		MessageId string `json:"message_id"`
		User      struct {
			Id   string `json:"id"`
			Type string `json:"type"`
		} `json:"user"`
		CreatedAt time.Time `json:"created_at"`
		IsRead    bool      `json:"is_read"`
		Data      []string  `json:"data"`
		Context   struct {
			Sku         string `json:"sku"`
			OrderNumber string `json:"order_number"`
		} `json:"context"`
		IsImage bool `json:"is_image"`
	} `json:"messages"`
}
