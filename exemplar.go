package ozon

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/whatcrm/go-ozon/models"
	"github.com/whatcrm/go-ozon/utils/posting"
)

func (c *Client) GetExemplarStatus(ctx context.Context, reqBody models.ExemplarStatusRequest) (*models.ExemplarStatusResponse, error) {
	requestURL := c.BaseURL + posting.PostingExemplarStatusV5Endpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.ExemplarStatusResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) ValidateExemplars(ctx context.Context, reqBody models.ExemplarValidateRequest) (*models.ExemplarValidateResponse, error) {
	requestURL := c.BaseURL + posting.PostingExemplarValidateV5Endpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.ExemplarValidateResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UpdateExemplars(ctx context.Context, reqBody models.ExemplarUpdateRequest) (*models.ExemplarUpdateResponse, error) {
	requestURL := c.BaseURL + posting.PostingExemplarUpdateV1Endpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.ExemplarUpdateResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) SetExemplars(ctx context.Context, reqBody models.ExemplarSetRequest) (*models.ExemplarSetResponse, error) {
	requestURL := c.BaseURL + posting.PostingExemplarSetV6Endpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.ExemplarSetResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) CreateOrGetExemplars(ctx context.Context, reqBody models.ExemplarCreateOrGetRequest) (*models.ExemplarCreateOrGetResponse, error) {
	requestURL := c.BaseURL + posting.PostingExemplarCreateOrGetV6Endpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.ExemplarCreateOrGetResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
