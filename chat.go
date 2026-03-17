package ozon

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/whatcrm/go-ozon/models"
	"github.com/whatcrm/go-ozon/utils/chat"
)

func (c *Client) GetList(ctx context.Context, filter models.ChatListRequest) (*models.ChatListResponse, error) {
	requestURL := c.BaseURL + chat.ListEndpoint

	jsonBody, err := json.Marshal(filter)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.ChatListResponse
	err = c.Send(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetHistory(ctx context.Context, filter models.ChatHistoryRequest) (*models.ChatHistoryResponse, error) {
	requestURL := c.BaseURL + chat.HistoryEndpoint

	jsonBody, err := json.Marshal(filter)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.ChatHistoryResponse
	err = c.Send(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) SendFile(ctx context.Context, file models.SendFileRequest) (*models.Response, error) {
	requestURL := c.BaseURL + chat.SendFileEndpoint

	jsonBody, err := json.Marshal(file)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.Response
	err = c.Send(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) SendMessage(ctx context.Context, message models.SendMessageRequest) (*models.Response, error) {
	requestURL := c.BaseURL + chat.SendMessageEndpoint

	jsonBody, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.Response
	err = c.Send(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) CreateChat(ctx context.Context, info models.StartChatRequest) (*models.Response, error) {
	requestURL := c.BaseURL + chat.StartChatEndpoint

	jsonBody, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.Response
	err = c.Send(req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) ReadChat(ctx context.Context, info models.ReadChatRequest) (*models.Response, error) {
	requestURL := c.BaseURL + chat.ReadChatEndpoint

	jsonBody, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.Response
	err = c.Send(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
