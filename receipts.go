package ozon

import (
	"bytes"
	"context"
	"encoding/json"
	"mime/multipart"
	"net/http"

	"github.com/whatcrm/go-ozon/models"
	"github.com/whatcrm/go-ozon/utils/receipts"
)

func (c *Client) GetReceipt(ctx context.Context, reqBody models.GetReceiptRequest) (*models.GetReceiptResponse, error) {
	requestURL := c.BaseURL + receipts.GetReceiptEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.GetReceiptResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetSellerReceipts(ctx context.Context, reqBody models.SellerListReceiptsRequest) (*models.SellerListReceiptsResponse, error) {
	requestURL := c.BaseURL + receipts.SellerListReceiptsEndpoint

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.SellerListReceiptsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UploadReceipt(ctx context.Context, data models.UploadReceiptRequest) (*models.UploadReceiptResponse, error) {
	requestURL := c.BaseURL + receipts.UploadReceiptEndpoint

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	contentField, err := writer.CreateFormFile("content", "receipt.pdf")
	if err != nil {
		return nil, err
	}
	if _, err = contentField.Write(data.Content); err != nil {
		return nil, err
	}

	if err = writer.WriteField("operation_type", data.OperationType); err != nil {
		return nil, err
	}

	if data.ParentReceiptID != "" {
		if err = writer.WriteField("parent_receipt_id", data.ParentReceiptID); err != nil {
			return nil, err
		}
	}

	for _, pn := range data.PostingNumbers {
		if err = writer.WriteField("posting_numbers", pn); err != nil {
			return nil, err
		}
	}

	if err = writer.WriteField("receipt_number", data.ReceiptNumber); err != nil {
		return nil, err
	}

	if err = writer.WriteField("type", data.Type); err != nil {
		return nil, err
	}

	if err = writer.Close(); err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, &buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	var response models.UploadReceiptResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
