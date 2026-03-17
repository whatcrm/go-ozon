package ozon

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/whatcrm/go-ozon/models"
	"github.com/whatcrm/go-ozon/utils/product"
)

func (c *Client) ImportProducts(ctx context.Context, body interface{}) (json.RawMessage, error) {
	requestURL := c.BaseURL + product.ImportProductEndpoint

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response json.RawMessage
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetImportProductsInfo(ctx context.Context, body interface{}) (json.RawMessage, error) {
	requestURL := c.BaseURL + product.InfoProductEndpoint

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response json.RawMessage
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ImportProductsBySKU(ctx context.Context, body interface{}) (json.RawMessage, error) {
	requestURL := c.BaseURL + product.ImportProductBySkuEndpoint

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response json.RawMessage
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) UpdateProductAttributes(ctx context.Context, body interface{}) (json.RawMessage, error) {
	requestURL := c.BaseURL + product.UpdateProductAttributesEndpoint

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response json.RawMessage
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ImportProductPictures(ctx context.Context, body interface{}) (json.RawMessage, error) {
	requestURL := c.BaseURL + product.ImportProductPicturesEndpoint

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response json.RawMessage
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ListProducts(ctx context.Context, body interface{}) (json.RawMessage, error) {
	requestURL := c.BaseURL + product.ListProductEndpoint

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response json.RawMessage
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetProductRatingBySKU(ctx context.Context, body interface{}) (json.RawMessage, error) {
	requestURL := c.BaseURL + product.InfoRatingBySkuProductEndpoint

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response json.RawMessage
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetProductsByIdentifier(ctx context.Context, body interface{}) (json.RawMessage, error) {
	requestURL := c.BaseURL + product.ListProductByIdentifierEndpoint

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response json.RawMessage
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetProductAttributesInfo(ctx context.Context, body interface{}) (json.RawMessage, error) {
	requestURL := c.BaseURL + product.InfoProductAttributesEndpoint

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response json.RawMessage
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetProductInfoDiscounted(ctx context.Context, reqBody models.ProductInfoDiscountedRequest) (*models.ProductInfoDiscountedResponse, error) {
	requestURL := c.BaseURL + product.ProductInfoDiscountedEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.ProductInfoDiscountedResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UpdateProductDiscount(ctx context.Context, reqBody models.ProductUpdateDiscountRequest) (*models.ProductUpdateDiscountResponse, error) {
	requestURL := c.BaseURL + product.ProductUpdateDiscountEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.ProductUpdateDiscountResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetProductDescription(ctx context.Context, body interface{}) (json.RawMessage, error) {
	requestURL := c.BaseURL + product.ProductInfoDescriptionEndpoint

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response json.RawMessage
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) UpdateProductsStocks(ctx context.Context, body interface{}) (json.RawMessage, error) {
	requestURL := c.BaseURL + product.ProductsStocksV2Endpoint

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response json.RawMessage
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetProductStocks(ctx context.Context, body interface{}) (json.RawMessage, error) {
	requestURL := c.BaseURL + product.ProductInfoStocksEndpoint

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response json.RawMessage
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ImportProductPrices(ctx context.Context, body interface{}) (json.RawMessage, error) {
	requestURL := c.BaseURL + product.ImportProductsPricesEndpoint

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response json.RawMessage
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetProductInfoLimit(ctx context.Context) (*models.ProductInfoLimitResponse, error) {
	requestURL := c.BaseURL + product.ProductInfoLimitEndpoint

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer([]byte(`{}`)))
	if err != nil {
		return nil, err
	}

	var response models.ProductInfoLimitResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UpdateProductOfferID(ctx context.Context, reqBody models.ProductUpdateOfferIDRequest) (*models.ProductUpdateOfferIDResponse, error) {
	requestURL := c.BaseURL + product.ProductUpdateOfferIDEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.ProductUpdateOfferIDResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) ArchiveProducts(ctx context.Context, reqBody models.ProductArchiveRequest) (*models.ProductArchiveResponse, error) {
	requestURL := c.BaseURL + product.ProductArchiveEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.ProductArchiveResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UnarchiveProducts(ctx context.Context, reqBody models.ProductUnarchiveRequest) (*models.ProductUnarchiveResponse, error) {
	requestURL := c.BaseURL + product.ProductUnarchiveEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.ProductUnarchiveResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) DeleteProductsV2(ctx context.Context, reqBody models.ProductsDeleteRequest) (*models.ProductsDeleteResponse, error) {
	requestURL := c.BaseURL + product.ProductsDeleteV2Endpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.ProductsDeleteResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetProductSubscription(ctx context.Context, reqBody models.ProductSubscriptionRequest) (*models.ProductSubscriptionResponse, error) {
	requestURL := c.BaseURL + product.ProductInfoSubscriptionEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.ProductSubscriptionResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetRelatedSKU(ctx context.Context, reqBody models.ProductRelatedSKURequest) (*models.ProductRelatedSKUResponse, error) {
	requestURL := c.BaseURL + product.ProductRelatedSKUGetEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.ProductRelatedSKUResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetProductPicturesInfoV2(ctx context.Context, reqBody models.ProductPicturesInfoRequest) (*models.ProductPicturesInfoResponse, error) {
	requestURL := c.BaseURL + product.ProductPicturesInfoV2Endpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.ProductPicturesInfoResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

//func (c *Client) GetProductsWithWrongVolume(ctx context.Context, reqBody models.ProductWrongVolumeRequest) (*models.ProductWrongVolumeResponse, error) {
//	requestURL := c.BaseURL + product.ProductInfoWrongVolumeEndpoint
//
//	jsonBody, err := json.Marshal(reqBody)
//	if err != nil {
//		return nil, err
//	}
//
//	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
//	if err != nil {
//		return nil, err
//	}
//
//	var response models.ProductWrongVolumeResponse
//	if err = c.Send(req, &response); err != nil {
//		return nil, err
//	}
//
//	return &response, nil
//}

func (c *Client) GetProductPlacementZoneInfo(ctx context.Context, reqBody models.ProductPlacementZoneRequest) (*models.ProductPlacementZoneResponse, error) {
	requestURL := c.BaseURL + product.ProductPlacementZoneInfoEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.ProductPlacementZoneResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UpdateProductActionTimer(ctx context.Context, reqBody models.ProductActionTimerUpdateRequest) (*models.ProductActionTimerUpdateResponse, error) {
	requestURL := c.BaseURL + product.ProductActionTimerUpdateEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.ProductActionTimerUpdateResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetProductActionTimerStatus(ctx context.Context, reqBody models.ProductActionTimerStatusRequest) (*models.ProductActionTimerStatusResponse, error) {
	requestURL := c.BaseURL + product.ProductActionTimerStatusEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.ProductActionTimerStatusResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetProductPricesInfoV5(ctx context.Context, reqBody models.ProductPricesInfoRequest) (*models.ProductPricesInfoResponse, error) {
	requestURL := c.BaseURL + product.ProductInfoPricesV5Endpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.ProductPricesInfoResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
