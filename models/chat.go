package models

import "time"

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

type ChatListResponse struct {
	Chats []struct {
		CreatedAt  time.Time `json:"created_at"`
		ChatId     string    `json:"chat_id"`
		ChatStatus string    `json:"chat_status"`
		ChatType   string    `json:"chat_type"`
	} `json:"chats"`
	FirstUnreadMessageId string `json:"first_unread_message_id"`
	LastMessageId        string `json:"last_message_id"`
	UnreadCount          int    `json:"unread_count"`
	TotalChatsCount      int    `json:"total_chats_count"`
	TotalUnreadCount     int    `json:"total_unread_count"`
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
		MessageId           string    `json:"message_id"`
		ModerateImageStatus string    `json:"moderate_image_status"`
		User                struct {
			Id   string `json:"id"`
			Type string `json:"type"`
		} `json:"user"`
	} `json:"messages"`
}
