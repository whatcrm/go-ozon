package ozon

import (
	"context"
	"net/http"

	"github.com/whatcrm/go-ozon/models"
	"github.com/whatcrm/go-ozon/utils/others"
)

func (c *Client) Role(ctx context.Context) (*models.RoleResponse, error) {
	requestURL := c.BaseURL + others.RoleEndpoint

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, nil)
	if err != nil {
		return nil, err
	}

	var response models.RoleResponse
	err = c.Send(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
