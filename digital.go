package ozon

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/whatcrm/go-ozon/models"
	"github.com/whatcrm/go-ozon/utils/digital"
)

func (c *Client) UploadPostingCodes(ctx context.Context, reqBody models.PostingCodesUploadRequest) (*models.PostingCodesUploadResponse, error) {
	requestURL := c.BaseURL + digital.PostingCodesUploadEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.PostingCodesUploadResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetPostingList(ctx context.Context, reqBody models.PostingListRequest) (*models.PostingListResponse, error) {
	requestURL := c.BaseURL + digital.PostingListEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.PostingListResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) ImportProductStocks(ctx context.Context, reqBody models.ProductStocksImportRequest) (*models.ProductStocksImportResponse, error) {
	requestURL := c.BaseURL + digital.ProductDigitalStocksImportEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.ProductStocksImportResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
