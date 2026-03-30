package ozon

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/whatcrm/go-ozon/models"
	"github.com/whatcrm/go-ozon/utils/posting"
)

func (c *Client) GetFbsPostingList(ctx context.Context, filter models.PaginatedGetPostingFBSListFilter) (*models.GetPostingFBSListResponseResultWrapper, error) {
	requestURL := c.BaseURL + posting.PostingFbsListV3Endpoint

	jsonBody, err := json.Marshal(filter)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.GetPostingFBSListResponseResultWrapper
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetFbsPosting(ctx context.Context, data models.PostingFBSData) (json.RawMessage, error) {
	requestURL := c.BaseURL + posting.PostingFbsGetV3Endpoint

	jsonBody, err := json.Marshal(data)
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

func (c *Client) GetFbsPostingProductCountryList(ctx context.Context, filter models.CountryFilter) (json.RawMessage, error) {
	requestURL := c.BaseURL + posting.PostingFbsProductCountryListV2Endpoint

	jsonBody, err := json.Marshal(filter)
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

func (c *Client) SetFbsPostingProductCountry(ctx context.Context, data models.OderData) (json.RawMessage, error) {
	requestURL := c.BaseURL + posting.PostingFbsProductCountrySetV2Endpoint

	jsonBody, err := json.Marshal(data)
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

func (c *Client) GetFbsPostingPackageLabel(ctx context.Context, data models.FBSPackageData) ([]byte, error) {
	requestURL := c.BaseURL + posting.PostingFbsPackageLabelV2Endpoint

	jsonBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err = c.Send(req, &buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c *Client) GetFboPostingList(ctx context.Context, filter models.PaginatedGetPostingFBOListFilter) (json.RawMessage, error) {
	requestURL := c.BaseURL + posting.PostingFboListV2Endpoint

	jsonBody, err := json.Marshal(filter)
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

func (c *Client) CreateFbsAct(ctx context.Context, data models.PostingFSBDeliveryData) (json.RawMessage, error) {
	requestURL := c.BaseURL + posting.PostingFbsActCreateV2Endpoint

	jsonBody, err := json.Marshal(data)
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

func (c *Client) GetFbsActPostings(ctx context.Context, data models.PostingFBSActData) (json.RawMessage, error) {
	requestURL := c.BaseURL + posting.PostingFbsActGetPostingsV2Endpoint

	jsonBody, err := json.Marshal(data)
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

func (c *Client) GetFbsActBarcode(ctx context.Context, data models.FBSActData) ([]byte, error) {
	requestURL := c.BaseURL + posting.PostingFbsActGetBarcodeV2Endpoint

	jsonBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err = c.Send(req, &buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c *Client) CheckFbsActStatus(ctx context.Context, data models.PostingFSBActData) (json.RawMessage, error) {
	requestURL := c.BaseURL + posting.PostingFbsActCheckStatusV2Endpoint

	jsonBody, err := json.Marshal(data)
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

func (c *Client) ShipFbsPostingV4(ctx context.Context, reqBody models.FbsShipRequest) (*models.FbsShipResponse, error) {
	requestURL := c.BaseURL + posting.PostingFbsShipV4Endpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.FbsShipResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) ShipFbsPostingPackageV4(ctx context.Context, reqBody models.FbsShipPackageRequest) (*models.FbsShipPackageResponse, error) {
	requestURL := c.BaseURL + posting.PostingFbsShipPackageV4Endpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.FbsShipPackageResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) CreateFbsPackageLabelTask(ctx context.Context, reqBody models.FbsPackageLabelTaskCreateRequest) (*models.FbsPackageLabelTaskCreateResponse, error) {
	requestURL := c.BaseURL + posting.PostingFbsPackageLabelTaskCreateV1Endpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.FbsPackageLabelTaskCreateResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetFbsPackageLabelTask(ctx context.Context, reqBody models.FbsPackageLabelTaskGetRequest) (*models.FbsPackageLabelTaskGetResponse, error) {
	requestURL := c.BaseURL + posting.PostingFbsPackageLabelTaskGetV1Endpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.FbsPackageLabelTaskGetResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetFbsCancelReasonsForPostings(ctx context.Context, reqBody models.FbsCancelReasonForPostingRequest) (*models.FbsCancelReasonForPostingResponse, error) {
	requestURL := c.BaseURL + posting.PostingFbsCancelReasonV1Endpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.FbsCancelReasonForPostingResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetFbsCancelReasonList(ctx context.Context) (*models.FbsCancelReasonListResponse, error) {
	requestURL := c.BaseURL + posting.PostingFbsCancelReasonListV2Endpoint

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, nil)
	if err != nil {
		return nil, err
	}

	var response models.FbsCancelReasonListResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) CancelFbsProducts(ctx context.Context, reqBody models.FbsProductCancelRequest) (*models.FbsProductCancelResponse, error) {
	requestURL := c.BaseURL + posting.PostingFbsProductCancelV2Endpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.FbsProductCancelResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) CancelFbsPosting(ctx context.Context, reqBody models.FbsCancelRequest) (*models.FbsCancelResponse, error) {
	requestURL := c.BaseURL + posting.PostingFbsCancelV2Endpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.FbsCancelResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
