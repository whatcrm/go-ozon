package ozon

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/whatcrm/go-ozon/models"
	"github.com/whatcrm/go-ozon/utils/actions"
)

func (c *Client) GetActions(ctx context.Context) (*models.GetSellerActionsResponseResultWrapper, error) {
	requestURL := c.BaseURL + actions.ListEndpoint

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}

	var response models.GetSellerActionsResponseResultWrapper
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetActionsCandidates(ctx context.Context, body models.PaginatedCandidatesForActions) (*models.GetActionsCandidatesResponseResultWrapper, error) {
	requestURL := c.BaseURL + actions.CandidatesEndpoint

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.GetActionsCandidatesResponseResultWrapper
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetActionsProducts(ctx context.Context, body models.PaginatedActionProducts) (*models.GetSellerProductResponseResultWrapper, error) {
	requestURL := c.BaseURL + actions.ProductsEndpoint

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.GetSellerProductResponseResultWrapper
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
