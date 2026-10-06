package ozon

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/whatcrm/go-ozon/models"
	"github.com/whatcrm/go-ozon/utils/notification"
)

func (c *Client) SetNotification(ctx context.Context, reqBody models.SetNotificationRequest) error {
	requestURL := c.BaseURL + notification.SetEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	return c.Send(req, nil)
}

func (c *Client) UpdateNotification(ctx context.Context, reqBody models.UpdateNotificationRequest) error {
	requestURL := c.BaseURL + notification.UpdateEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	return c.Send(req, nil)
}

func (c *Client) DeleteNotification(ctx context.Context, reqBody models.DeleteNotificationRequest) error {
	requestURL := c.BaseURL + notification.DeleteEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	return c.Send(req, nil)
}

func (c *Client) CheckNotification(ctx context.Context, reqBody models.CheckNotificationRequest) (*models.CheckNotificationResponse, error) {
	requestURL := c.BaseURL + notification.CheckEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.CheckNotificationResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) EnableNotification(ctx context.Context, reqBody models.EnableNotificationRequest) error {
	requestURL := c.BaseURL + notification.EnableEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	return c.Send(req, nil)
}

func (c *Client) ListNotifications(ctx context.Context, reqBody models.ListNotificationsRequest) (*models.ListNotificationsResponse, error) {
	requestURL := c.BaseURL + notification.ListEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.ListNotificationsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) ListNotificationPushTypes(ctx context.Context) (*models.ListNotificationPushTypesResponse, error) {
	requestURL := c.BaseURL + notification.PushTypeListEndpoint

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, nil)
	if err != nil {
		return nil, err
	}

	var response models.ListNotificationPushTypesResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
