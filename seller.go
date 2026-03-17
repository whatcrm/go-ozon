package ozon

import (
	"context"
	"net/http"

	"github.com/whatcrm/go-ozon/models"
	"github.com/whatcrm/go-ozon/utils/others"
)

func (c *Client) GetInfoSeller(ctx context.Context) (*models.SellerResponse, error) {
	requestURL := c.BaseURL + others.SellerInfoEndpoint

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, nil)
	if err != nil {
		return nil, err
	}

	var response models.SellerResponse
	err = c.Send(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil

}

func (c *Client) GetLogisticsInfo(ctx context.Context) (*models.LogisticsResponse, error) {
	requestURL := c.BaseURL + others.LogisticsInfoEndpoint

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, nil)
	if err != nil {
		return nil, err
	}

	var response models.LogisticsResponse
	err = c.Send(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
