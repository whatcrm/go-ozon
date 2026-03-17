package models

import "time"

type PostingAdditionalFields struct {
	AnalyticsData bool `json:"analytics_data,omitempty"`
	Barcodes      bool `json:"barcodes,omitempty"`
	FinancialData bool `json:"financial_data,omitempty"`
	Translit      bool `json:"translit,omitempty"`
}

type GetPostingFBSListFilter struct {
	DeliveryMethodID []int      `json:"delivery_method_id,omitempty"`
	OrderID          *int       `json:"order_id,omitempty"`
	ProviderID       []int      `json:"provider_id,omitempty"`
	Status           string     `json:"status,omitempty"`
	WarehouseID      []int      `json:"warehouse_id,omitempty"`
	Since            *time.Time `json:"since,omitempty"`
	To               *time.Time `json:"to,omitempty"`
}

type PaginatedGetPostingFBSListFilter struct {
	Filter *GetPostingFBSListFilter `json:"filter,omitempty"`
	Dir    string                   `json:"dir,omitempty"`
	Limit  *int                     `json:"limit,omitempty"`
	Offset *int                     `json:"offset,omitempty"`
	With   *PostingAdditionalFields `json:"with,omitempty"`
}

type GetPostingFBSListResponseRequirements struct {
	ProductsRequiringGTD           []int `json:"products_requiring_gtd"`
	ProductsRequiringCountry       []int `json:"products_requiring_country"`
	ProductsRequiringMandatoryMark []int `json:"products_requiring_mandatory_mark"`
	ProductsRequiringRNPT          []int `json:"products_requiring_rnpt"`
}

type GetPostingFBSListResponseProduct struct {
	MandatoryMark []string `json:"mandatory_mark"`
	Name          string   `json:"name"`
	OfferID       string   `json:"offer_id"`
	Price         string   `json:"price"`
	Quantity      int      `json:"quantity"`
	SKU           int      `json:"sku"`
	CurrencyCode  string   `json:"currency_code"`
}

type GetPostingFBSListResponsePicking struct {
	Amount float64   `json:"amount"`
	Tag    string    `json:"tag"`
	Moment time.Time `json:"moment"`
}

type GetPostingFBSListResponseFinancialDataServices struct {
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

type GetPostingFBSListResponseFinancialDataProduct struct {
	Actions              []string                                       `json:"actions"`
	ClientPrice          string                                         `json:"client_price"`
	CommissionAmount     float64                                        `json:"commission_amount"`
	CommissionPercent    int                                            `json:"commission_percent"`
	ItemServices         GetPostingFBSListResponseFinancialDataServices `json:"item_services"`
	OldPrice             float64                                        `json:"old_price"`
	Payout               float64                                        `json:"payout"`
	Picking              GetPostingFBSListResponsePicking               `json:"picking"`
	Price                float64                                        `json:"price"`
	ProductID            int                                            `json:"product_id"`
	Quantity             int                                            `json:"quantity"`
	TotalDiscountPercent float64                                        `json:"total_discount_percent"`
	TotalDiscountValue   float64                                        `json:"total_discount_value"`
}

type GetPostingFBSListResponseFinancialData struct {
	PostingServices GetPostingFBSListResponseFinancialDataServices  `json:"posting_services"`
	Products        []GetPostingFBSListResponseFinancialDataProduct `json:"products"`
}

type GetPostingFBSListResponseDeliveryMethod struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	TplProvider   string `json:"tpl_provider"`
	TplProviderID int    `json:"tpl_provider_id"`
	Warehouse     string `json:"warehouse"`
	WarehouseID   int    `json:"warehouse_id"`
}

type GetPostingFBSListResponseAddress struct {
	AddressTail     string  `json:"address_tail"`
	City            string  `json:"city"`
	Comment         string  `json:"comment"`
	Country         string  `json:"country"`
	District        string  `json:"district"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	ProviderPVZCode string  `json:"provider_pvz_code"`
	PVZCode         int     `json:"pvz_code"`
	Region          string  `json:"region"`
	ZipCode         string  `json:"zip_code"`
}

type GetPostingFBSListResponseCustomer struct {
	Address       GetPostingFBSListResponseAddress `json:"address"`
	CustomerEmail string                           `json:"customer_email"`
	CustomerID    int                              `json:"customer_id"`
	Name          string                           `json:"name"`
	Phone         string                           `json:"phone"`
}

type GetPostingFBSListResponseCancellation struct {
	AffectCancellationRating bool   `json:"affect_cancellation_rating"`
	CancelReason             string `json:"cancel_reason"`
	CancelReasonID           int    `json:"cancel_reason_id"`
	CancellationInitiator    string `json:"cancellation_initiator"`
	CancellationType         string `json:"cancellation_type"`
	CancelledAfterShip       bool   `json:"cancelled_after_ship"`
}

type GetPostingFBSListResponseBarcodes struct {
	LowerBarcode string `json:"lower_barcode"`
	UpperBarcode string `json:"upper_barcode"`
}

type GetPostingFBSListResponseAnalyticsData struct {
	City                 string     `json:"city"`
	IsPremium            bool       `json:"is_premium"`
	PaymentTypeGroupName string     `json:"payment_type_group_name"`
	Region               string     `json:"region"`
	TplProvider          string     `json:"tpl_provider"`
	TplProviderID        int        `json:"tpl_provider_id"`
	Warehouse            string     `json:"warehouse"`
	WarehouseID          int        `json:"warehouse_id"`
	DeliveryDateBegin    *time.Time `json:"delivery_date_begin,omitempty"`
}

type GetPostingFBSListResponseAddressee struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type GetPostingFBSListResponsePosting struct {
	Addressee          *GetPostingFBSListResponseAddressee      `json:"addressee,omitempty"`
	AnalyticsData      *GetPostingFBSListResponseAnalyticsData  `json:"analytics_data,omitempty"`
	Barcodes           *GetPostingFBSListResponseBarcodes       `json:"barcodes,omitempty"`
	Cancellation       *GetPostingFBSListResponseCancellation   `json:"cancellation,omitempty"`
	Customer           *GetPostingFBSListResponseCustomer       `json:"customer,omitempty"`
	DeliveryMethod     *GetPostingFBSListResponseDeliveryMethod `json:"delivery_method,omitempty"`
	FinancialData      *GetPostingFBSListResponseFinancialData  `json:"financial_data,omitempty"`
	IsExpress          bool                                     `json:"is_express"`
	OrderID            int                                      `json:"order_id"`
	OrderNumber        string                                   `json:"order_number"`
	PostingNumber      string                                   `json:"posting_number"`
	Products           []GetPostingFBSListResponseProduct       `json:"products"`
	Requirements       *GetPostingFBSListResponseRequirements   `json:"requirements,omitempty"`
	Status             string                                   `json:"status"`
	TplIntegrationType string                                   `json:"tpl_integration_type"`
	TrackingNumber     string                                   `json:"tracking_number"`
	DeliveringDate     *time.Time                               `json:"delivering_date,omitempty"`
	InProcessAt        *time.Time                               `json:"in_process_at,omitempty"`
	ShipmentDate       *time.Time                               `json:"shipment_date,omitempty"`
}

type GetPostingFBSListResponseResult struct {
	Postings []GetPostingFBSListResponsePosting `json:"postings"`
	HasNext  bool                               `json:"has_next"`
}

type GetPostingFBSListResponseResultWrapper struct {
	Result GetPostingFBSListResponseResult `json:"result"`
}

type FbsShipPackageProduct struct {
	ProductID int64 `json:"product_id"`
	Quantity  int32 `json:"quantity"`
}

type FbsShipPackage struct {
	Products []FbsShipPackageProduct `json:"products"`
}

type FbsShipWith struct {
	AdditionalData bool `json:"additional_data"`
}

type FbsShipRequest struct {
	Packages      []FbsShipPackage `json:"packages"`
	PostingNumber string           `json:"posting_number"`
	With          *FbsShipWith     `json:"with,omitempty"`
}

type FbsShipAdditionalDataProduct struct {
	CurrencyCode  string   `json:"currency_code"`
	MandatoryMark []string `json:"mandatory_mark"`
	Name          string   `json:"name"`
	OfferID       string   `json:"offer_id"`
	Price         string   `json:"price"`
	Quantity      int32    `json:"quantity"`
	SKU           int64    `json:"sku"`
}

type FbsShipAdditionalData struct {
	PostingNumber string                         `json:"posting_number"`
	Products      []FbsShipAdditionalDataProduct `json:"products"`
}

type FbsShipResponse struct {
	AdditionalData []FbsShipAdditionalData `json:"additional_data"`
	Result         []string                `json:"result"`
}

type FbsShipPackageProductRequest struct {
	ExemplarsIds []string `json:"exemplarsIds,omitempty"`
	ProductID    int64    `json:"product_id"`
	Quantity     int32    `json:"quantity"`
}

type FbsShipPackageRequest struct {
	PostingNumber string                         `json:"posting_number"`
	Products      []FbsShipPackageProductRequest `json:"products"`
}

type FbsShipPackageResponse struct {
	Result string `json:"result"`
}

type PostingFBSData struct {
	PostingNumber string                   `json:"posting_number"`
	With          *PostingAdditionalFields `json:"with,omitempty"`
}

type CountryFilter struct {
	NameSearch string `json:"name_search,omitempty"`
}

type OderData struct {
	PostingNumber  *string `json:"posting_number,omitempty"`
	ProductID      *int    `json:"product_id,omitempty"`
	CountryISOCode *string `json:"country_iso_code,omitempty"`
}

type FBSPackageData struct {
	PostingNumber []string `json:"posting_number"`
}

type PostingFSBDeliveryData struct {
	ContainersCount  *int       `json:"containers_count,omitempty"`
	DeliveryMethodID *int       `json:"delivery_method_id,omitempty"`
	DepartureDate    *time.Time `json:"departure_date,omitempty"`
}

type PostingFBSActData struct {
	ID *int `json:"id,omitempty"`
}

type FBSActData struct {
	ID int `json:"id"`
}

type PostingFSBActData struct {
	ID *int `json:"id,omitempty"`
}
