package ozon

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/whatcrm/go-ozon/models"
	"github.com/whatcrm/go-ozon/utils/description"
)

func (c *Client) GetTreeCategory(ctx context.Context, language string) (*models.TreeCategoryResponse, error) {
	requestBody := c.BaseURL + description.TreeEndpoint

	jsonBody, err := json.Marshal(language)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", requestBody, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.TreeCategoryResponse
	err = c.Send(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetAttributeCategory(ctx context.Context, attribute models.AttributeCategoryRequest) (*models.AttributeCategoryResponse, error) {
	requestBody := c.BaseURL + description.AttributeEndpoint

	jsonBody, err := json.Marshal(attribute)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", requestBody, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.AttributeCategoryResponse
	err = c.Send(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetAttributeValues(ctx context.Context, attribute models.AttributeValueRequest) (*models.AttributeValueResponse, error) {
	requestBody := c.BaseURL + description.ValuesEndpoint

	jsonBody, err := json.Marshal(attribute)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", requestBody, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.AttributeValueResponse
	err = c.Send(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) SearchAttributeValues(ctx context.Context, attribute models.AttributeValueRequest) (*models.AttributeValueResponse, error) {
	requestBody := c.BaseURL + description.ValuesSearchEndpoint

	jsonBody, err := json.Marshal(attribute)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", requestBody, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.AttributeValueResponse
	err = c.Send(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
