package ozon

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/whatcrm/go-ozon/models"
	"github.com/whatcrm/go-ozon/utils/rating"
)

func (c *Client) GetRatingSummary(ctx context.Context) (*models.RatingSummaryResponse, error) {
	requestURL := c.BaseURL + rating.SummaryEndpoint

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, nil)
	if err != nil {
		return nil, err
	}

	var response models.RatingSummaryResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetRatingHistory(ctx context.Context, reqBody models.RatingHistoryRequest) (*models.RatingHistoryResponse, error) {
	requestURL := c.BaseURL + rating.HistoryEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.RatingHistoryResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetFbsIndexInfo(ctx context.Context) (*models.RatingFbsIndexInfoResponse, error) {
	requestURL := c.BaseURL + rating.FbsIndexInfoEndpoint

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, nil)
	if err != nil {
		return nil, err
	}

	var response models.RatingFbsIndexInfoResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetFbsIndexPostingList(ctx context.Context, reqBody models.RatingFbsIndexPostingListRequest) (*models.RatingFbsIndexPostingListResponse, error) {
	requestURL := c.BaseURL + rating.FbsIndexPostingListEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.RatingFbsIndexPostingListResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
