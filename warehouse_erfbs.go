package ozon

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/whatcrm/go-ozon/models"
	"github.com/whatcrm/go-ozon/utils/warehouse"
)

func (c *Client) CreateErfbsAggregatorWarehouse(ctx context.Context, reqBody models.ErfbsAggregatorCreateRequest) (*models.WarehouseOperationResponse, error) {
	requestURL := c.BaseURL + warehouse.ErfbsAggregatorCreateEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.WarehouseOperationResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UpdateErfbsWarehouse(ctx context.Context, reqBody models.ErfbsWarehouseUpdateRequest) (*models.WarehouseOperationResponse, error) {
	requestURL := c.BaseURL + warehouse.ErfbsUpdateWarehouseEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.WarehouseOperationResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UpdateErfbsAggregatorDeliveryMethod(ctx context.Context, reqBody models.ErfbsAggregatorDeliveryMethodUpdateRequest) (*models.WarehouseOperationResponse, error) {
	requestURL := c.BaseURL + warehouse.ErfbsAggregatorDeliveryMethodUpdateEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.WarehouseOperationResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) CreateErfbsNonIntegratedWarehouse(ctx context.Context, reqBody models.ErfbsNonIntegratedCreateRequest) (*models.WarehouseOperationResponse, error) {
	requestURL := c.BaseURL + warehouse.ErfbsNonIntegratedCreateEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.WarehouseOperationResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UpdateErfbsNonIntegratedDeliveryMethod(ctx context.Context, reqBody models.ErfbsNonIntegratedDeliveryMethodUpdateRequest) (*models.WarehouseOperationResponse, error) {
	requestURL := c.BaseURL + warehouse.ErfbsNonIntegratedDeliveryMethodUpdateEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.WarehouseOperationResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
