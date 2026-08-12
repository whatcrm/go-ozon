package ozon

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/whatcrm/go-ozon/models"
	"github.com/whatcrm/go-ozon/utils/review"
)

func (c *Client) CreateReviewComment(ctx context.Context, reqBody models.CreateReviewCommentRequest) (*models.CreateReviewCommentResponse, error) {
	requestURL := c.BaseURL + review.CommentCreateEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.CreateReviewCommentResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) DeleteReviewComment(ctx context.Context, reqBody models.DeleteReviewCommentRequest) error {
	requestURL := c.BaseURL + review.CommentDeleteV2Endpoint

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

func (c *Client) DeleteReviewCommentV1(ctx context.Context, reqBody models.DeleteReviewCommentV1Request) error {
	requestURL := c.BaseURL + review.CommentDeleteV1Endpoint

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

func (c *Client) ListReviewComments(ctx context.Context, reqBody models.ListReviewCommentsRequest) (*models.ListReviewCommentsResponse, error) {
	requestURL := c.BaseURL + review.CommentListEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.ListReviewCommentsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) ChangeReviewStatus(ctx context.Context, reqBody models.ChangeReviewStatusRequest) error {
	requestURL := c.BaseURL + review.ChangeStatusV2Endpoint

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

func (c *Client) ChangeReviewStatusV1(ctx context.Context, reqBody models.ChangeReviewStatusV1Request) error {
	requestURL := c.BaseURL + review.ChangeStatusV1Endpoint

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

func (c *Client) GetReviewCount(ctx context.Context) (*models.GetReviewCountResponse, error) {
	requestURL := c.BaseURL + review.CountV2Endpoint

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, nil)
	if err != nil {
		return nil, err
	}

	var response models.GetReviewCountResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetReviewCountV1(ctx context.Context) (*models.GetReviewCountV1Response, error) {
	requestURL := c.BaseURL + review.CountV1Endpoint

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer([]byte("{}")))
	if err != nil {
		return nil, err
	}

	var response models.GetReviewCountV1Response
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetReviewInfo(ctx context.Context, reqBody models.GetReviewInfoRequest) (*models.GetReviewInfoResponse, error) {
	requestURL := c.BaseURL + review.InfoV2Endpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.GetReviewInfoResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetReviewInfoV1(ctx context.Context, reqBody models.GetReviewInfoRequest) (*models.GetReviewInfoResponse, error) {
	requestURL := c.BaseURL + review.InfoV1Endpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.GetReviewInfoResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) ListReviews(ctx context.Context, reqBody models.ListReviewsRequest) (*models.ListReviewsResponse, error) {
	requestURL := c.BaseURL + review.ListV2Endpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.ListReviewsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
