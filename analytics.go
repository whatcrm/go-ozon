package ozon

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/whatcrm/go-ozon/models"
	"github.com/whatcrm/go-ozon/utils/analytics"
)

func (c *Client) GetStockOnWarehousesV2(ctx context.Context, reqBody models.StockOnWarehousesV2Request) (*models.StockOnWarehousesV2Response, error) {
	requestURL := c.BaseURL + analytics.StockOnWarehousesV2Endpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.StockOnWarehousesV2Response
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetTurnoverStocks(ctx context.Context, reqBody models.TurnoverStocksRequest) (*models.TurnoverStocksResponse, error) {
	requestURL := c.BaseURL + analytics.TurnoverStocksEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.TurnoverStocksResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetAverageDeliveryTime(ctx context.Context, reqBody models.AverageDeliveryTimeRequest) (*models.AverageDeliveryTimeResponse, error) {
	requestURL := c.BaseURL + analytics.AverageDeliveryTimeEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.AverageDeliveryTimeResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetAverageDeliveryTimeDetails(ctx context.Context, reqBody models.AverageDeliveryTimeDetailsRequest) (*models.AverageDeliveryTimeDetailsResponse, error) {
	requestURL := c.BaseURL + analytics.AverageDeliveryTimeDetailsEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.AverageDeliveryTimeDetailsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetAverageDeliveryTimeSummary(ctx context.Context) (*models.AverageDeliveryTimeSummaryResponse, error) {
	requestURL := c.BaseURL + analytics.AverageDeliveryTimeSummaryEndpoint

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, nil)
	if err != nil {
		return nil, err
	}

	var response models.AverageDeliveryTimeSummaryResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
