package ozon

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/whatcrm/go-ozon/models"
	"github.com/whatcrm/go-ozon/utils/product"
)

func (c *Client) GetProductInfoStocksWithModel(ctx context.Context, filter models.PaginatedProductFilter) (*models.GetProductInfoStocksResponseResult, error) {
	requestURL := c.BaseURL + product.ProductInfoStocksEndpoint

	jsonBody, err := json.Marshal(filter)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.GetProductInfoStocksResponseResult
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) ImportProductsStocks(ctx context.Context, stocks models.ProductImportProductsStocks) (*models.ProductsStocksResponseProcessResultWrapper, error) {
	requestURL := c.BaseURL + product.ProductsStocksV2Endpoint

	jsonBody, err := json.Marshal(stocks)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.ProductsStocksResponseProcessResultWrapper
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) SetProductsStocks(ctx context.Context, stocks models.StocksData) (*models.SetProductStocksResponseResultWrapper, error) {
	requestURL := c.BaseURL + product.ProductsStocksV2Endpoint

	jsonBody, err := json.Marshal(stocks)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.SetProductStocksResponseResultWrapper
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetProductWarehouseStocks(ctx context.Context, reqBody models.ProductWarehouseStocksRequest) (*models.ProductWarehouseStocksResponse, error) {
	requestURL := c.BaseURL + product.ProductInfoWarehouseStocksEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.ProductWarehouseStocksResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetProductStocksByWarehouseV2FBS(ctx context.Context, reqBody models.ProductStocksByWarehouseV2Request) (*models.ProductStocksByWarehouseV2Response, error) {
	requestURL := c.BaseURL + product.ProductInfoStocksByWarehouseV2Fbs

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.ProductStocksByWarehouseV2Response
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetProductStocksByWarehouseV1FBS(ctx context.Context, reqBody models.ProductStocksByWarehouseV1Request) (*models.ProductStocksByWarehouseV1Response, error) {
	requestURL := c.BaseURL + product.ProductInfoStocksByWarehouseV1Fbs

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.ProductStocksByWarehouseV1Response
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
