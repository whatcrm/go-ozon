package ozon

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/whatcrm/go-ozon/models"
	"github.com/whatcrm/go-ozon/utils/product"
)

func (c *Client) ImportProductPricesWithModel(ctx context.Context, data models.PricesData) (*models.GetProductImportPriceResponseResultWrapper, error) {
	requestURL := c.BaseURL + product.ImportProductsPricesEndpoint

	jsonBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.GetProductImportPriceResponseResultWrapper
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
