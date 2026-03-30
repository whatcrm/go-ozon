package models

import "time"

type GetReceiptRequest struct {
	ReceiptID string `json:"receipt_id"`
}

type GetReceiptResponse struct {
	Content string `json:"content"`
}

type SellerListReceiptsRequest struct {
	Page           int64    `json:"page,omitempty"`
	PageSize       int64    `json:"page_size,omitempty"`
	PostingNumbers []string `json:"posting_numbers,omitempty"`
}

type SellerListReceiptsResponse struct {
	HasNext  bool                `json:"has_next"`
	Receipts []SellerReceiptItem `json:"receipts"`
}

type SellerReceiptItem struct {
	CreatedAt       time.Time `json:"created_at"`
	OperationType   string    `json:"operation_type"`
	OrderID         int64     `json:"order_id"`
	ParentReceiptID string    `json:"parent_receipt_id"`
	PostingNumbers  []string  `json:"posting_numbers"`
	ReceiptID       string    `json:"receipt_id"`
	ReceiptNumber   string    `json:"receipt_number"`
	Type            string    `json:"type"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type UploadReceiptRequest struct {
	Content         []byte
	OperationType   string
	ParentReceiptID string
	PostingNumbers  []string
	ReceiptNumber   string
	Type            string
}

type UploadReceiptResponse struct {
	ReceiptID string `json:"receipt_id"`
}
