package ozon

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/whatcrm/go-ozon/models"
	"github.com/whatcrm/go-ozon/utils/question"
)

func (c *Client) CreateQuestionAnswer(ctx context.Context, reqBody models.CreateQuestionAnswerRequest) (*models.CreateQuestionAnswerResponse, error) {
	requestURL := c.BaseURL + question.AnswerCreateEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.CreateQuestionAnswerResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) DeleteQuestionAnswer(ctx context.Context, reqBody models.DeleteQuestionAnswerRequest) error {
	requestURL := c.BaseURL + question.AnswerDeleteEndpoint

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

func (c *Client) ListQuestionAnswers(ctx context.Context, reqBody models.ListQuestionAnswersRequest) (*models.ListQuestionAnswersResponse, error) {
	requestURL := c.BaseURL + question.AnswerListEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.ListQuestionAnswersResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) ChangeQuestionStatus(ctx context.Context, reqBody models.ChangeQuestionStatusRequest) error {
	requestURL := c.BaseURL + question.ChangeStatusEndpoint

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

func (c *Client) GetQuestionCount(ctx context.Context) (*models.GetQuestionCountResponse, error) {
	requestURL := c.BaseURL + question.CountEndpoint

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, nil)
	if err != nil {
		return nil, err
	}

	var response models.GetQuestionCountResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetQuestionInfo(ctx context.Context, reqBody models.GetQuestionInfoRequest) (*models.GetQuestionInfoResponse, error) {
	requestURL := c.BaseURL + question.InfoEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.GetQuestionInfoResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) ListQuestions(ctx context.Context, reqBody models.ListQuestionsRequest) (*models.ListQuestionsResponse, error) {
	requestURL := c.BaseURL + question.ListEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.ListQuestionsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetQuestionTopSKU(ctx context.Context, reqBody models.GetQuestionTopSKURequest) (*models.GetQuestionTopSKUResponse, error) {
	requestURL := c.BaseURL + question.TopSKUEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.GetQuestionTopSKUResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
