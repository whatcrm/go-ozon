package models

import (
	"encoding/json"
	"time"
)

type PostingCodesUploadRequest struct {
	ExemplarsBySKU []PostingCodesUploadRequestExemplarsBySKU `json:"exemplars_by_sku"`
	PostingNumber  string                                    `json:"posting_number"`
}

type PostingCodesUploadRequestExemplarsBySKU struct {
	ExemplarQty             int32    `json:"exemplar_qty"`
	NotAvailableExemplarQty int32    `json:"not_available_exemplar_qty"`
	SKU                     int64    `json:"sku"`
	ExemplarKeys            []string `json:"exemplar_keys"`
}

type PostingCodesUploadResponse struct {
	ExemplarsBySKU []PostingCodesUploadResponseExemplarsBySKU `json:"exemplars_by_sku"`
}

type PostingCodesUploadResponseExemplarsBySKU struct {
	SKU             int64                                      `json:"sku"`
	ReceivedQty     int32                                      `json:"received_qty"`
	RejectedQty     int32                                      `json:"rejected_qty"`
	FailedExemplars []PostingCodesUploadResponseFailedExemplar `json:"failed_exemplars"`
}

type PostingCodesUploadResponseFailedExemplar struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}

type PostingListRequest struct {
	Dir    string            `json:"dir"`
	Filter PostingListFilter `json:"filter"`
	Limit  int64             `json:"limit"`
	Offset int64             `json:"offset"`
	With   PostingListWith   `json:"with"`
}

type PostingListFilter struct {
	Since         time.Time `json:"since"`
	PostingNumber []string  `json:"posting_number,omitempty"`
	To            time.Time `json:"to"`
}

type PostingListWith struct {
	AnalyticsData bool `json:"analytics_data"`
	FinancialData bool `json:"financial_data"`
	LegalInfo     bool `json:"legal_info"`
}

type PostingListResponse struct {
	Result []PostingListResponseItem `json:"result"`
}

type PostingListResponseItem struct {
	OrderID                int64                                   `json:"order_id"`
	OrderNumber            string                                  `json:"order_number"`
	PostingNumber          string                                  `json:"posting_number"`
	Status                 string                                  `json:"status"`
	CancelReasonID         int64                                   `json:"cancel_reason_id"`
	CreatedAt              time.Time                               `json:"created_at"`
	InProcessAt            time.Time                               `json:"in_process_at"`
	WaitingDeadlineForCode time.Time                               `json:"waiting_deadline_for_digital_code"`
	LegalInfo              *PostingListResponseLegalInfo           `json:"legal_info,omitempty"`
	Products               []PostingListResponseProduct            `json:"products"`
	AnalyticsData          *PostingListResponseAnalyticsData       `json:"analytics_data,omitempty"`
	FinancialData          *PostingListResponseFinancialData       `json:"financial_data,omitempty"`
	AdditionalData         []PostingListResponseAdditionalDataItem `json:"additional_data"`
}

type PostingListResponseLegalInfo struct {
	CompanyName string `json:"company_name"`
	Inn         string `json:"inn"`
	Kpp         string `json:"kpp"`
}

type PostingListResponseProduct struct {
	SKU                int64  `json:"sku"`
	Name               string `json:"name"`
	OfferID            string `json:"offer_id"`
	Price              string `json:"price"`
	RequiredQtyForCode int32  `json:"required_qty_for_digital_code"`
	CurrencyCode       string `json:"currency_code"`
}

type PostingListResponseAnalyticsData struct {
	City                 string `json:"city"`
	DeliveryType         string `json:"delivery_type"`
	IsPremium            bool   `json:"is_premium"`
	PaymentTypeGroupName string `json:"payment_type_group_name"`
	WarehouseID          int64  `json:"warehouse_id"`
	WarehouseName        string `json:"warehouse_name"`
	IsLegal              bool   `json:"is_legal"`
}

type PostingListResponseFinancialData struct {
	Products []PostingListResponseFinancialDataProduct `json:"products"`
}

type PostingListResponseFinancialDataProduct struct {
	CommissionAmount     float64  `json:"commission_amount"`
	CommissionPercent    int32    `json:"commission_percent"`
	Payout               float64  `json:"payout"`
	ProductID            int64    `json:"product_id"`
	CurrencyCode         string   `json:"currency_code"`
	OldPrice             float64  `json:"old_price"`
	Price                float64  `json:"price"`
	TotalDiscountValue   float64  `json:"total_discount_value"`
	TotalDiscountPercent float64  `json:"total_discount_percent"`
	Actions              []string `json:"actions"`
}

type PostingListResponseAdditionalDataItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ProductStocksImportRequest struct {
	Stocks []ProductStocksImportRequestItem `json:"stocks"`
}

type ProductStocksImportRequestItem struct {
	Stock   string `json:"stock"`
	OfferID string `json:"offer_id"`
}

type ProductStocksImportResponse struct {
	Status []ProductStocksImportResponseStatus `json:"status"`
}

type ProductStocksImportResponseStatus struct {
	Errors    []ProductStocksImportResponseError `json:"errors"`
	OfferID   string                             `json:"offer_id"`
	Updated   bool                               `json:"updated"`
	ProductID int64                              `json:"product_id"`
	SKU       int64                              `json:"sku"`
}

type ProductStocksImportResponseError struct {
	Code    int32             `json:"code"`
	Details []json.RawMessage `json:"details"`
	Message string            `json:"message"`
}
