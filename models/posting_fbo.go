package models

import "time"

type PostingFBOAdditionalFields struct {
	AnalyticsData bool `json:"analytics_data,omitempty"`
	FinancialData bool `json:"financial_data,omitempty"`
}

type GetPostingFBOListFilter struct {
	Since  *time.Time `json:"since,omitempty"`
	To     *time.Time `json:"to,omitempty"`
	Status string     `json:"status,omitempty"`
}

type PaginatedGetPostingFBOListFilter struct {
	Filter   *GetPostingFBOListFilter    `json:"filter,omitempty"`
	Dir      string                      `json:"dir,omitempty"`
	Translit bool                        `json:"translit,omitempty"`
	Limit    *int                        `json:"limit,omitempty"`
	Offset   *int                        `json:"offset,omitempty"`
	With     *PostingFBOAdditionalFields `json:"with,omitempty"`
}

type GetPostingFBOListResponseProduct struct {
	DigitalCodes []string `json:"digital_codes"`
	Name         string   `json:"name"`
	OfferID      string   `json:"offer_id"`
	Price        string   `json:"price"`
	Quantity     int      `json:"quantity"`
	SKU          int      `json:"sku"`
}

type GetPostingFBOListResponsePicking struct {
	Amount float64   `json:"amount"`
	Tag    string    `json:"tag"`
	Moment time.Time `json:"moment"`
}

type GetPostingFBOListResponseFinancialDataServices struct {
	MarketplaceServiceItemDelivToCustomer            float64 `json:"marketplace_service_item_deliv_to_customer"`
	MarketplaceServiceItemDirectFlowTrans            float64 `json:"marketplace_service_item_direct_flow_trans"`
	MarketplaceServiceItemDropoffFF                  float64 `json:"marketplace_service_item_dropoff_ff"`
	MarketplaceServiceItemDropoffPVZ                 float64 `json:"marketplace_service_item_dropoff_pvz"`
	MarketplaceServiceItemDropoffSC                  float64 `json:"marketplace_service_item_dropoff_sc"`
	MarketplaceServiceItemFulfillment                float64 `json:"marketplace_service_item_fulfillment"`
	MarketplaceServiceItemPickup                     float64 `json:"marketplace_service_item_pickup"`
	MarketplaceServiceItemReturnAfterDelivToCustomer float64 `json:"marketplace_service_item_return_after_deliv_to_customer"`
	MarketplaceServiceItemReturnFlowTrans            float64 `json:"marketplace_service_item_return_flow_trans"`
	MarketplaceServiceItemReturnNotDelivToCustomer   float64 `json:"marketplace_service_item_return_not_deliv_to_customer"`
	MarketplaceServiceItemReturnPartGoodsCustomer    float64 `json:"marketplace_service_item_return_part_goods_customer"`
}

type GetPostingFBOListResponseFinancialDataProduct struct {
	Actions              []string                                       `json:"actions"`
	ClientPrice          string                                         `json:"client_price"`
	CommissionAmount     float64                                        `json:"commission_amount"`
	CommissionPercent    int                                            `json:"commission_percent"`
	ItemServices         GetPostingFBOListResponseFinancialDataServices `json:"item_services"`
	OldPrice             float64                                        `json:"old_price"`
	Payout               float64                                        `json:"payout"`
	Picking              GetPostingFBOListResponsePicking               `json:"picking"`
	Price                float64                                        `json:"price"`
	ProductID            int                                            `json:"product_id"`
	Quantity             int                                            `json:"quantity"`
	TotalDiscountPercent float64                                        `json:"total_discount_percent"`
	TotalDiscountValue   float64                                        `json:"total_discount_value"`
}

type GetPostingFBOListResponseFinancialData struct {
	PostingServices GetPostingFBOListResponseFinancialDataServices  `json:"posting_services"`
	Products        []GetPostingFBOListResponseFinancialDataProduct `json:"products"`
}

type GetPostingFBOListResponseAnalyticsData struct {
	City                 string `json:"city"`
	DeliveryType         string `json:"delivery_type"`
	IsLegal              bool   `json:"is_legal"`
	IsPremium            bool   `json:"is_premium"`
	PaymentTypeGroupName string `json:"payment_type_group_name"`
	Region               string `json:"region"`
	WarehouseID          int    `json:"warehouse_id"`
	WarehouseName        string `json:"warehouse_name"`
}

type GetPostingFBOAdditionalDataItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type GetPostingFBOListResponseResult struct {
	AdditionalData []GetPostingFBOAdditionalDataItem       `json:"additional_data"`
	AnalyticsData  *GetPostingFBOListResponseAnalyticsData `json:"analytics_data,omitempty"`
	CancelReasonID int                                     `json:"cancel_reason_id"`
	FinancialData  *GetPostingFBOListResponseFinancialData `json:"financial_data,omitempty"`
	OrderID        int                                     `json:"order_id"`
	OrderNumber    string                                  `json:"order_number"`
	PostingNumber  string                                  `json:"posting_number"`
	Products       []GetPostingFBOListResponseProduct      `json:"products"`
	Status         string                                  `json:"status"`
	CreatedAt      *time.Time                              `json:"created_at,omitempty"`
	InProcessAt    *time.Time                              `json:"in_process_at,omitempty"`
}

type GetPostingFBOListResponseResultWrapper struct {
	Result []GetPostingFBOListResponseResult `json:"result"`
}
