package ozon

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/whatcrm/go-ozon/models"
	"github.com/whatcrm/go-ozon/utils/finance"
)

func (c *Client) GetRealizationV2(ctx context.Context, reqBody models.RealizationV2Request) (*models.RealizationV2Response, error) {
	requestURL := c.BaseURL + finance.RealizationV2Endpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.RealizationV2Response
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetRealizationPosting(ctx context.Context, reqBody models.RealizationPostingRequest) (*models.RealizationPostingResponse, error) {
	requestURL := c.BaseURL + finance.RealizationPostingEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.RealizationPostingResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetTransactionList(ctx context.Context, reqBody models.TransactionListRequest) (*models.TransactionListResponse, error) {
	requestURL := c.BaseURL + finance.TransactionListV3Endpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.TransactionListResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetTransactionTotals(ctx context.Context, reqBody models.TransactionTotalsRequest) (*models.TransactionTotalsResponse, error) {
	requestURL := c.BaseURL + finance.TransactionTotalsV3Endpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.TransactionTotalsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetDocumentB2BSales(ctx context.Context, reqBody models.PeriodLanguageRequest) (*models.DocumentCodeResponse, error) {
	requestURL := c.BaseURL + finance.DocumentB2BSalesEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.DocumentCodeResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetDocumentB2BSalesJSON(ctx context.Context, reqBody models.B2BSalesJSONRequest) (*models.B2BSalesJSONResponse, error) {
	requestURL := c.BaseURL + finance.DocumentB2BSalesJSONEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.B2BSalesJSONResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetMutualSettlement(ctx context.Context, reqBody models.PeriodLanguageRequest) (*models.DocumentCodeResponse, error) {
	requestURL := c.BaseURL + finance.MutualSettlementEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.DocumentCodeResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetBuyoutProducts(ctx context.Context, reqBody models.BuyoutProductsRequest) (*models.BuyoutProductsResponse, error) {
	requestURL := c.BaseURL + finance.ProductsBuyoutEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.BuyoutProductsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetCompensationReport(ctx context.Context, reqBody models.CompensationRequest) (*models.CompensationResponse, error) {
	requestURL := c.BaseURL + finance.CompensationEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.CompensationResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetDecompensationReport(ctx context.Context, reqBody models.DecompensationRequest) (*models.DecompensationResponse, error) {
	requestURL := c.BaseURL + finance.DecompensationEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.DecompensationResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
