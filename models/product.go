package models

import "time"

// product_action_timer.go

type ProductActionTimerUpdateRequest struct {
	ProductIDs []int64 `json:"product_ids"`
}

type ProductActionTimerUpdateResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ProductActionTimerStatusRequest struct {
	ProductIDs []int64 `json:"product_ids"`
}

type ProductActionTimerStatusItem struct {
	ExpiredAt                     string `json:"expired_at"`
	MinPriceForAutoActionsEnabled bool   `json:"min_price_for_auto_actions_enabled"`
	ProductID                     int64  `json:"product_id"`
}

type ProductActionTimerStatusResponse struct {
	Statuses []ProductActionTimerStatusItem `json:"statuses"`
}

// product_archive.go

type ProductArchiveRequest struct {
	ProductIDs []int64 `json:"product_id"`
}

type ProductArchiveResponse struct {
	Result bool `json:"result"`
}

type ProductUnarchiveRequest struct {
	ProductIDs []int64 `json:"product_id"`
}

type ProductUnarchiveResponse struct {
	Result bool `json:"result"`
}

type ProductDeleteItem struct {
	OfferID string `json:"offer_id"`
}

type ProductsDeleteRequest struct {
	Products []ProductDeleteItem `json:"products"`
}

type ProductsDeleteStatus struct {
	Error     string `json:"error"`
	IsDeleted bool   `json:"is_deleted"`
	OfferID   string `json:"offer_id"`
}

type ProductsDeleteResponse struct {
	Status []ProductsDeleteStatus `json:"status"`
}

// product_description.go

type ProductDescriptionRequest struct {
	OfferID   string `json:"offer_id,omitempty"`
	ProductID int64  `json:"product_id,omitempty"`
}

type GetProductInfoDescriptionResponseResult struct {
	Description string `json:"description"`
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	OfferID     string `json:"offer_id"`
}

type GetProductInfoDescriptionResponseResultWrapper struct {
	Result GetProductInfoDescriptionResponseResult `json:"result"`
}

// product_discounted.go

type ProductInfoDiscountedRequest struct {
	DiscountedSkus []string `json:"discounted_skus"`
}

type ProductInfoDiscountedItem struct {
	CommentReasonDamaged string `json:"comment_reason_damaged"`
	Condition            string `json:"condition"`
	ConditionEstimation  string `json:"condition_estimation"`
	Defects              string `json:"defects"`
	DiscountedSku        int64  `json:"discounted_sku"`
	MechanicalDamage     string `json:"mechanical_damage"`
	PackageDamage        string `json:"package_damage"`
	PackagingViolation   string `json:"packaging_violation"`
	ReasonDamaged        string `json:"reason_damaged"`
	Repair               string `json:"repair"`
	Shortage             string `json:"shortage"`
	SKU                  int64  `json:"sku"`
	WarrantyType         string `json:"warranty_type"`
}

type ProductInfoDiscountedResponse struct {
	Items []ProductInfoDiscountedItem `json:"items"`
}

type ProductUpdateDiscountRequest struct {
	Discount  int32 `json:"discount"`
	ProductID int64 `json:"product_id"`
}

type ProductUpdateDiscountResponse struct {
	Result bool `json:"result"`
}

// product_limits.go

type ProductLimitInfo struct {
	Limit   int       `json:"limit"`
	Usage   int       `json:"usage"`
	ResetAt time.Time `json:"reset_at"`
}

type ProductLimitTotalInfo struct {
	Limit int `json:"limit"`
	Usage int `json:"usage"`
}

type ProductInfoLimitResponse struct {
	DailyCreate ProductLimitInfo      `json:"daily_create"`
	DailyUpdate ProductLimitInfo      `json:"daily_update"`
	Total       ProductLimitTotalInfo `json:"total"`
}

// product_offer_id.go

type ProductUpdateOfferIDItem struct {
	NewOfferID string `json:"new_offer_id"`
	OfferID    string `json:"offer_id"`
}

type ProductUpdateOfferIDRequest struct {
	UpdateOfferID []ProductUpdateOfferIDItem `json:"update_offer_id"`
}

type ProductUpdateOfferIDError struct {
	Message string `json:"message"`
	OfferID string `json:"offer_id"`
}

type ProductUpdateOfferIDResponse struct {
	Errors []ProductUpdateOfferIDError `json:"errors"`
}

// product_pictures.go

type ProductPicturesInfoRequest struct {
	ProductIDs []string `json:"product_id"`
}

type ProductPicturesError struct {
	Message string `json:"message"`
	URL     string `json:"url"`
}

type ProductPicturesItem struct {
	ProductID    int64                  `json:"product_id"`
	PrimaryPhoto []string               `json:"primary_photo"`
	Photo        []string               `json:"photo"`
	ColorPhoto   []string               `json:"color_photo"`
	Photo360     []string               `json:"photo_360"`
	Errors       []ProductPicturesError `json:"errors"`
}

type ProductPicturesInfoResponse struct {
	Items []ProductPicturesItem `json:"items"`
}

// product_placement_zone.go

type ProductPlacementZoneRequest struct {
	SKUs []string `json:"skus"`
}

type ProductPlacementZoneItem struct {
	PlacementZone string `json:"placement_zone"`
	SKU           int64  `json:"sku"`
}

type ProductPlacementZoneResponse struct {
	ProductsPlacement []ProductPlacementZoneItem `json:"products_placement"`
}

// product_prices_info.go

type ProductPricesInfoFilter struct {
	OfferID    []string `json:"offer_id,omitempty"`
	ProductID  []string `json:"product_id,omitempty"`
	Visibility string   `json:"visibility,omitempty"`
}

type ProductPricesInfoRequest struct {
	Cursor string                  `json:"cursor"`
	Filter ProductPricesInfoFilter `json:"filter"`
	Limit  int32                   `json:"limit"`
}

type ProductPricesInfoCommissions struct {
	FboDelivToCustomerAmount    float64 `json:"fbo_deliv_to_customer_amount"`
	FboDirectFlowTransMaxAmount float64 `json:"fbo_direct_flow_trans_max_amount"`
	FboDirectFlowTransMinAmount float64 `json:"fbo_direct_flow_trans_min_amount"`
	FboReturnFlowAmount         float64 `json:"fbo_return_flow_amount"`
	FbsDelivToCustomerAmount    float64 `json:"fbs_deliv_to_customer_amount"`
	FbsDirectFlowTransMaxAmount float64 `json:"fbs_direct_flow_trans_max_amount"`
	FbsDirectFlowTransMinAmount float64 `json:"fbs_direct_flow_trans_min_amount"`
	FbsFirstMileMaxAmount       float64 `json:"fbs_first_mile_max_amount"`
	FbsFirstMileMinAmount       float64 `json:"fbs_first_mile_min_amount"`
	FbsReturnFlowAmount         float64 `json:"fbs_return_flow_amount"`
	SalesPercentFbo             float64 `json:"sales_percent_fbo"`
	SalesPercentFbp             float64 `json:"sales_percent_fbp"`
	SalesPercentFbs             float64 `json:"sales_percent_fbs"`
	SalesPercentRfbs            float64 `json:"sales_percent_rfbs"`
}

type ProductPricesInfoAction struct {
	DateFrom string  `json:"date_from"`
	DateTo   string  `json:"date_to"`
	Title    string  `json:"title"`
	Value    float64 `json:"value"`
}

type ProductPricesInfoMarketingActions struct {
	Actions           []ProductPricesInfoAction `json:"actions"`
	CurrentPeriodFrom string                    `json:"current_period_from"`
	CurrentPeriodTo   string                    `json:"current_period_to"`
	OzonActionsExist  bool                      `json:"ozon_actions_exist"`
}

type ProductPricesInfoPrice struct {
	AutoActionEnabled               bool    `json:"auto_action_enabled"`
	AutoAddToOzonActionsListEnabled bool    `json:"auto_add_to_ozon_actions_list_enabled"`
	CurrencyCode                    string  `json:"currency_code"`
	MarketingSellerPrice            float64 `json:"marketing_seller_price"`
	MinPrice                        float64 `json:"min_price"`
	NetPrice                        float64 `json:"net_price"`
	OldPrice                        float64 `json:"old_price"`
	Price                           float64 `json:"price"`
	RetailPrice                     float64 `json:"retail_price"`
	VAT                             float64 `json:"vat"`
}

type ProductPricesInfoExternalIndexData struct {
	MinPrice         float64 `json:"min_price"`
	MinPriceCurrency string  `json:"min_price_currency"`
	PriceIndexValue  float64 `json:"price_index_value"`
}

type ProductPricesInfoPriceIndexes struct {
	ColorIndex                string                             `json:"color_index"`
	ExternalIndexData         ProductPricesInfoExternalIndexData `json:"external_index_data"`
	OzonIndexData             ProductPricesInfoExternalIndexData `json:"ozon_index_data"`
	SelfMarketplacesIndexData ProductPricesInfoExternalIndexData `json:"self_marketplaces_index_data"`
}

type ProductPricesInfoItem struct {
	Acquiring        float64                           `json:"acquiring"`
	Commissions      ProductPricesInfoCommissions      `json:"commissions"`
	MarketingActions ProductPricesInfoMarketingActions `json:"marketing_actions"`
	OfferID          string                            `json:"offer_id"`
	Price            ProductPricesInfoPrice            `json:"price"`
	PriceIndexes     ProductPricesInfoPriceIndexes     `json:"price_indexes"`
	ProductID        int64                             `json:"product_id"`
	VolumeWeight     float64                           `json:"volume_weight"`
}

type ProductPricesInfoResponse struct {
	Cursor string                  `json:"cursor"`
	Items  []ProductPricesInfoItem `json:"items"`
	Total  int32                   `json:"total"`
}

// product_related_sku.go

type ProductRelatedSKURequest struct {
	SKU []string `json:"sku"`
}

type ProductRelatedSKUItem struct {
	Availability   string `json:"availability"`
	DeletedAt      string `json:"deleted_at"`
	DeliverySchema string `json:"delivery_schema"`
	ProductID      int64  `json:"product_id"`
	SKU            int64  `json:"sku"`
}

type ProductRelatedSKUError struct {
	Code    string `json:"code"`
	SKU     int64  `json:"sku"`
	Message string `json:"message"`
}

type ProductRelatedSKUResponse struct {
	Items  []ProductRelatedSKUItem  `json:"items"`
	Errors []ProductRelatedSKUError `json:"errors"`
}

// product_stocks_by_warehouse.go

type ProductStocksByWarehouseV2Filter struct {
	OfferID []string `json:"offer_id,omitempty"`
	SKU     []string `json:"sku,omitempty"`
}

type ProductStocksByWarehouseV2Request struct {
	Cursor  string   `json:"cursor"`
	Limit   uint64   `json:"limit"`
	OfferID []string `json:"offer_id,omitempty"`
	SKU     []string `json:"sku"`
}

type ProductStocksByWarehouseV2Item struct {
	FreeStock     int64  `json:"free_stock"`
	OfferID       string `json:"offer_id"`
	Present       int64  `json:"present"`
	ProductID     int64  `json:"product_id"`
	Reserved      int64  `json:"reserved"`
	SKU           int64  `json:"sku"`
	WarehouseID   int64  `json:"warehouse_id"`
	WarehouseName string `json:"warehouse_name"`
}

type ProductStocksByWarehouseV2Response struct {
	Cursor   string                           `json:"cursor"`
	HasNext  bool                             `json:"has_next"`
	Products []ProductStocksByWarehouseV2Item `json:"products"`
}

type ProductStocksByWarehouseV1Request struct {
	SKU     []string `json:"sku,omitempty"`
	OfferID []string `json:"offer_id,omitempty"`
}

type ProductStocksByWarehouseV1Item struct {
	SKU           int64  `json:"sku"`
	OfferID       string `json:"offer_id"`
	Present       int64  `json:"present"`
	ProductID     int64  `json:"product_id"`
	Reserved      int64  `json:"reserved"`
	WarehouseID   int64  `json:"warehouse_id"`
	WarehouseName string `json:"warehouse_name"`
}

type ProductStocksByWarehouseV1Response struct {
	Result []ProductStocksByWarehouseV1Item `json:"result"`
}

// product_subscription.go

type ProductSubscriptionRequest struct {
	SKUs []string `json:"skus"`
}

type ProductSubscriptionItem struct {
	Count int64 `json:"count"`
	SKU   int64 `json:"sku"`
}

type ProductSubscriptionResponse struct {
	Result []ProductSubscriptionItem `json:"result"`
}

// product_warehouse_stocks.go

type ProductWarehouseStocksRequest struct {
	Cursor      string `json:"cursor"`
	Limit       int64  `json:"limit"`
	WarehouseID int64  `json:"warehouse_id"`
}

type ProductWarehouseStockItem struct {
	SKU         int64  `json:"sku"`
	ProductID   int64  `json:"product_id"`
	OfferID     string `json:"offer_id"`
	WarehouseID int64  `json:"warehouse_id"`
	Present     int64  `json:"present"`
	Reserved    int64  `json:"reserved"`
	FreeStock   int64  `json:"free_stock"`
	UpdatedAt   string `json:"updated_at"`
}

type ProductWarehouseStocksResponse struct {
	Cursor  string                      `json:"cursor"`
	HasNext bool                        `json:"has_next"`
	Stocks  []ProductWarehouseStockItem `json:"stocks"`
}
