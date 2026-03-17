package models

type StockOnWarehousesV2Request struct {
	Limit         int64  `json:"limit"`
	Offset        int64  `json:"offset,omitempty"`
	WarehouseType string `json:"warehouse_type,omitempty"`
}

type StockOnWarehousesV2Response struct {
	Result StockOnWarehousesV2Result `json:"result"`
}

type StockOnWarehousesV2Result struct {
	Rows []StockOnWarehousesV2Row `json:"rows"`
}

type StockOnWarehousesV2Row struct {
	FreeToSellAmount int64  `json:"free_to_sell_amount"`
	ItemCode         string `json:"item_code"`
	ItemName         string `json:"item_name"`
	PromisedAmount   int64  `json:"promised_amount"`
	ReservedAmount   int64  `json:"reserved_amount"`
	SKU              int64  `json:"sku"`
	WarehouseName    string `json:"warehouse_name"`
}

type TurnoverStocksRequest struct {
	Limit  int32    `json:"limit,omitempty"`
	Offset int32    `json:"offset,omitempty"`
	SKU    []string `json:"sku,omitempty"`
}

type TurnoverStocksResponse struct {
	Items []TurnoverStocksItem `json:"items"`
}

type TurnoverStocksItem struct {
	ADS           float64 `json:"ads"`
	CurrentStock  int64   `json:"current_stock"`
	IDC           float64 `json:"idc"`
	IDCGrade      string  `json:"idc_grade"`
	Name          string  `json:"name"`
	OfferID       string  `json:"offer_id"`
	SKU           int64   `json:"sku"`
	Turnover      float64 `json:"turnover"`
	TurnoverGrade string  `json:"turnover_grade"`
}

type AverageDeliveryTimeRequest struct {
	DeliverySchema string   `json:"delivery_schema"`
	SKU            []string `json:"sku,omitempty"`
	SupplyPeriod   string   `json:"supply_period"`
}

type AverageDeliveryTimeResponse struct {
	Data  []AverageDeliveryTimeCluster `json:"data"`
	Total AverageDeliveryTimeTotal     `json:"total"`
}

type AverageDeliveryTimeCluster struct {
	ClustersData      []AverageDeliveryTimeClusterData `json:"clusters_data"`
	DeliveryClusterID int64                            `json:"delivery_cluster_id"`
	Metrics           AverageDeliveryTimeMetrics       `json:"metrics"`
}

type AverageDeliveryTimeClusterData struct {
	AnotherDeliveryTime []AverageDeliveryTimeAnother `json:"another_delivery_time"`
	ClusterID           int64                        `json:"cluster_id"`
	DeliveryTimeFBO     int64                        `json:"delivery_time_FBO"`
	DeliveryTimeFBS     int64                        `json:"delivery_time_FBS"`
	DeliveryTimeStatus  string                       `json:"delivery_time_status"`
	OrdersCount         int64                        `json:"orders_count"`
	OrdersPercent       int64                        `json:"orders_percent"`
}

type AverageDeliveryTimeAnother struct {
	DeliveryTime       int64  `json:"delivery_time"`
	OrdersCount        int64  `json:"orders_count"`
	OrdersPercent      int64  `json:"orders_percent"`
	DeliveryTimeStatus string `json:"delivery_time_status"`
}

type AverageDeliveryTimeMetrics struct {
	AverageDeliveryTime       int64                          `json:"average_delivery_time"`
	AverageDeliveryTimeStatus string                         `json:"average_delivery_time_status"`
	RecommendedSupply         int64                          `json:"recommended_supply"`
	OrdersCount               AverageDeliveryTimeOrdersCount `json:"orders_count"`
	ImpactShare               int64                          `json:"impact_share"`
	ExactImpactShare          string                         `json:"exact_impact_share"`
	AttentionLevel            string                         `json:"attention_level"`
	LostProfit                int64                          `json:"lost_profit"`
}

type AverageDeliveryTimeOrdersCount struct {
	Total  int64                         `json:"total"`
	Fast   AverageDeliveryTimeOrdersPart `json:"fast"`
	Medium AverageDeliveryTimeOrdersPart `json:"medium"`
	Long   AverageDeliveryTimeOrdersPart `json:"long"`
}

type AverageDeliveryTimeOrdersPart struct {
	Value   int64 `json:"value"`
	Percent int64 `json:"percent"`
}

type AverageDeliveryTimeTotal struct {
	AverageDeliveryTime       int64                          `json:"average_delivery_time"`
	AverageDeliveryTimeStatus string                         `json:"average_delivery_time_status"`
	RecommendedSupply         int64                          `json:"recommended_supply"`
	OrdersCount               AverageDeliveryTimeOrdersCount `json:"orders_count"`
	ImpactShare               int64                          `json:"impact_share"`
	ExactImpactShare          string                         `json:"exact_impact_share"`
	AttentionLevel            string                         `json:"attention_level"`
	LostProfit                int64                          `json:"lost_profit"`
}

type AverageDeliveryTimeDetailsRequest struct {
	ClusterID int64                             `json:"cluster_id"`
	Filters   AverageDeliveryTimeDetailsFilters `json:"filters"`
	Limit     int32                             `json:"limit"`
	Offset    int32                             `json:"offset"`
}

type AverageDeliveryTimeDetailsFilters struct {
	DeliverySchema string `json:"delivery_schema"`
	SupplyPeriod   string `json:"supply_period"`
}

type AverageDeliveryTimeDetailsResponse struct {
	Data      []AverageDeliveryTimeDetailsItem `json:"data"`
	TotalRows int64                            `json:"total_rows"`
}

type AverageDeliveryTimeDetailsItem struct {
	ClustersData []AverageDeliveryTimeClusterData  `json:"clusters_data"`
	Item         AverageDeliveryTimeDetailsProduct `json:"item"`
	Metrics      AverageDeliveryTimeMetrics        `json:"metrics"`
}

type AverageDeliveryTimeDetailsProduct struct {
	Name           string `json:"name"`
	DeliverySchema string `json:"delivery_schema"`
	SKU            int64  `json:"sku"`
	OfferID        string `json:"offer_id"`
}

type AverageDeliveryTimeSummaryResponse struct {
	AverageDeliveryTime int32                            `json:"average_delivery_time"`
	PerfectDeliveryTime int32                            `json:"perfect_delivery_time"`
	CurrentTariff       AverageDeliveryTimeSummaryTariff `json:"current_tariff"`
	UpdatedAt           string                           `json:"updated_at"`
	LostProfit          float64                          `json:"lost_profit"`
}

type AverageDeliveryTimeSummaryTariff struct {
	Start        int32   `json:"start"`
	TariffStatus string  `json:"tariff_status"`
	TariffValue  float64 `json:"tariff_value"`
	Fee          float64 `json:"fee"`
}
