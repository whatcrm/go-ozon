package models

type RealizationV2Request struct {
	Month int32 `json:"month"`
	Year  int32 `json:"year"`
}

type RealizationV2Response struct {
	Result RealizationV2Result `json:"result"`
}

type RealizationV2Result struct {
	Header RealizationHeader  `json:"header"`
	Rows   []RealizationRowV2 `json:"rows"`
}

type RealizationHeader struct {
	ContractDate    string `json:"contract_date"`
	ContractNumber  string `json:"contract_number"`
	CurrencySysName string `json:"currency_sys_name"`
	DocDate         string `json:"doc_date"`
	Number          string `json:"number"`
	PayerInn        string `json:"payer_inn"`
	PayerKpp        string `json:"payer_kpp"`
	PayerName       string `json:"payer_name"`
	ReceiverInn     string `json:"receiver_inn"`
	ReceiverKpp     string `json:"receiver_kpp"`
	ReceiverName    string `json:"receiver_name"`
	StartDate       string `json:"start_date"`
	StopDate        string `json:"stop_date"`
}

type RealizationRowV2 struct {
	CommissionRatio        float64               `json:"commission_ratio"`
	DeliveryCommission     RealizationCommission `json:"delivery_commission"`
	Item                   RealizationItem       `json:"item"`
	ReturnCommission       RealizationCommission `json:"return_commission"`
	RowNumber              int64                 `json:"rowNumber"`
	SellerPricePerInstance float64               `json:"seller_price_per_instance"`
}

type RealizationCommission struct {
	Amount                  float64 `json:"amount"`
	Bonus                   float64 `json:"bonus"`
	Commission              float64 `json:"commission"`
	Compensation            float64 `json:"compensation"`
	PricePerInstance        float64 `json:"price_per_instance"`
	Quantity                float64 `json:"quantity"`
	StandardFee             float64 `json:"standard_fee"`
	BankCoinvestment        float64 `json:"bank_coinvestment"`
	Stars                   float64 `json:"stars"`
	PickUpPointCoinvestment float64 `json:"pick_up_point_coinvestment"`
	Total                   float64 `json:"total"`
}

type RealizationItem struct {
	Barcode string `json:"barcode"`
	Name    string `json:"name"`
	OfferID string `json:"offer_id"`
	SKU     int64  `json:"sku"`
}

type RealizationPostingRequest struct {
	Month int32 `json:"month"`
	Year  int32 `json:"year"`
}

type RealizationPostingResponse struct {
	Header RealizationHeader       `json:"header"`
	Rows   []RealizationPostingRow `json:"rows"`
}

type RealizationPostingRow struct {
	CommissionRatio        float64                               `json:"commission_ratio"`
	DeliveryCommission     RealizationCommission                 `json:"delivery_commission"`
	Item                   RealizationItem                       `json:"item"`
	ReturnCommission       RealizationCommission                 `json:"return_commission"`
	RowNumber              int64                                 `json:"row_number"`
	SellerPricePerInstance float64                               `json:"seller_price_per_instance"`
	Order                  RealizationPostingOrder               `json:"order"`
	LegalEntityDocument    RealizationPostingLegalEntityDocument `json:"legal_entity_document"`
}

type RealizationPostingOrder struct {
	PostingNumber string `json:"posting_number"`
	CreatedDate   string `json:"created_date"`
}

type RealizationPostingLegalEntityDocument struct {
	Number   string `json:"number"`
	SaleDate string `json:"sale_date"`
}

type TransactionListRequest struct {
	Filter   TransactionListFilter `json:"filter"`
	Page     int64                 `json:"page"`
	PageSize int64                 `json:"page_size"`
}

type TransactionListFilter struct {
	Date            *TransactionDateFilter `json:"date,omitempty"`
	OperationType   []string               `json:"operation_type,omitempty"`
	PostingNumber   string                 `json:"posting_number,omitempty"`
	TransactionType string                 `json:"transaction_type,omitempty"`
}

type TransactionDateFilter struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type TransactionListResponse struct {
	Result TransactionListResult `json:"result"`
}

type TransactionListResult struct {
	Operations []TransactionOperation `json:"operations"`
	PageCount  int64                  `json:"page_count"`
	RowCount   int64                  `json:"row_count"`
}

type TransactionOperation struct {
	OperationID          int64   `json:"operation_id"`
	OperationType        string  `json:"operation_type"`
	OperationDate        string  `json:"operation_date"`
	OperationTypeName    string  `json:"operation_type_name"`
	DeliveryCharge       float64 `json:"delivery_charge"`
	ReturnDeliveryCharge float64 `json:"return_delivery_charge"`
	AccrualsForSale      float64 `json:"accruals_for_sale"`
	SaleCommission       float64 `json:"sale_commission"`
	Amount               float64 `json:"amount"`
	Type                 string  `json:"type"`
}

type TransactionTotalsRequest struct {
	Date            *TransactionDateFilter `json:"date,omitempty"`
	PostingNumber   string                 `json:"posting_number,omitempty"`
	TransactionType string                 `json:"transaction_type,omitempty"`
}

type TransactionTotalsResponse struct {
	Result TransactionTotalsResult `json:"result"`
}

type TransactionTotalsResult struct {
	AccrualsForSale         float64 `json:"accruals_for_sale"`
	SaleCommission          float64 `json:"sale_commission"`
	ProcessingAndDelivery   float64 `json:"processing_and_delivery"`
	RefundsAndCancellations float64 `json:"refunds_and_cancellations"`
	ServicesAmount          float64 `json:"services_amount"`
	CompensationAmount      float64 `json:"compensation_amount"`
	MoneyTransfer           float64 `json:"money_transfer"`
	OthersAmount            float64 `json:"others_amount"`
}

type PeriodLanguageRequest struct {
	Date     string `json:"date"`
	Language string `json:"language,omitempty"`
}

type DocumentCodeResponse struct {
	Result DocumentCodeResult `json:"result"`
}

type DocumentCodeResult struct {
	Code string `json:"code"`
}

type B2BSalesJSONRequest struct {
	Date string `json:"date"`
}

type B2BSalesJSONResponse struct {
	DateFrom   string             `json:"date_from"`
	DateTo     string             `json:"date_to"`
	Invoices   []B2BSalesInvoice  `json:"invoices"`
	SellerInfo B2BSalesSellerInfo `json:"seller_info"`
}

type B2BSalesInvoice struct {
	BuyerInfo    B2BSalesBuyerInfo   `json:"buyer_info"`
	Currency     string              `json:"currency"`
	CurrencyCode int64               `json:"currency_code"`
	Info         B2BSalesInvoiceInfo `json:"info"`
	OfferID      string              `json:"offer_id"`
	Operations   []B2BSalesOperation `json:"operations"`
	ProductName  string              `json:"product_name"`
	SKU          int64               `json:"sku"`
	UnitCode     int64               `json:"unit_code"`
	UnitName     string              `json:"unit_name"`
}

type B2BSalesBuyerInfo struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	INN     string `json:"inn"`
	KPP     string `json:"kpp"`
}

type B2BSalesSellerInfo struct {
	CompanyName string `json:"company_name"`
	INN         string `json:"inn"`
	KPP         string `json:"kpp"`
}

type B2BSalesInvoiceInfo struct {
	Date   string `json:"date"`
	Number string `json:"number"`
	Status string `json:"status"`
	Type   string `json:"type"`
}

type B2BSalesOperation struct {
	Amount         float64 `json:"amount"`
	CostWithoutVAT float64 `json:"cost_without_vat"`
	Date           string  `json:"date"`
	GTDNumber      string  `json:"gtd_number"`
	OriginCountry  string  `json:"origin_country"`
	PostingNumber  string  `json:"posting_number"`
	Price          float64 `json:"price"`
	Quantity       float64 `json:"quantity"`
	RNPTNumber     string  `json:"rnpt_number"`
	Type           string  `json:"type"`
	VATAmount      float64 `json:"vat_amount"`
	VATRate        float64 `json:"vat_rate"`
}

type BuyoutProductsRequest struct {
	DateFrom string `json:"date_from"`
	DateTo   string `json:"date_to"`
}

type BuyoutProductsResponse struct {
	Products []BuyoutProductItem `json:"products"`
}

type BuyoutProductItem struct {
	Name                       string  `json:"name"`
	OfferID                    string  `json:"offer_id"`
	SKU                        string  `json:"sku"`
	PostingNumber              string  `json:"posting_number"`
	SellerPricePerInstance     float64 `json:"seller_price_per_instance"`
	DeductionByCategoryPercent string  `json:"deduction_by_category_percent"`
	BuyoutPrice                float64 `json:"buyout_price"`
	VATPercent                 int32   `json:"vat_percent"`
	Quantity                   int32   `json:"quantity"`
	Amount                     float64 `json:"amount"`
}

type CompensationRequest struct {
	Date     string `json:"date"`
	Language string `json:"language,omitempty"`
}

type CompensationResponse struct {
	Result DocumentCodeResult `json:"result"`
}

type DecompensationRequest struct {
	Date     string `json:"date"`
	Language string `json:"language,omitempty"`
}

type DecompensationResponse struct {
	Result DocumentCodeResult `json:"result"`
}
