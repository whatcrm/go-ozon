package models

type ProductFilterWithQuant struct {
	Created *bool `json:"created,omitempty"`
}

type ProductFilter struct {
	OfferID    []string                `json:"offer_id,omitempty"`
	ProductID  []string                `json:"product_id,omitempty"`
	Visibility string                  `json:"visibility,omitempty"`
	WithQuant  *ProductFilterWithQuant `json:"with_quant,omitempty"`
}

type PaginatedProductFilter struct {
	Filter ProductFilter `json:"filter"`
	Cursor string        `json:"cursor"`
	Limit  int           `json:"limit"`
}

type GetProductInfoStocksResponseStock struct {
	Present  int    `json:"present"`
	Reserved int    `json:"reserved"`
	Type     string `json:"type"`
}

type GetProductInfoStocksResponseItem struct {
	OfferID   string                              `json:"offer_id"`
	ProductID int64                               `json:"product_id"`
	Stocks    []GetProductInfoStocksResponseStock `json:"stocks"`
}

type GetProductInfoStocksResponseResult struct {
	Cursor string                             `json:"cursor"`
	Items  []GetProductInfoStocksResponseItem `json:"items"`
	Total  int                                `json:"total"`
}

type ProductsStocksList struct {
	OfferID     string `json:"offer_id"`
	ProductID   int64  `json:"product_id"`
	Stock       int    `json:"stock"`
	WarehouseID int    `json:"warehouse_id"`
}

type ProductImportProductsStocks struct {
	Stocks []ProductsStocksList `json:"stocks"`
}

type ProductImportProductsStocksResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ProductsStocksResponseProcessResult struct {
	Errors      []ProductImportProductsStocksResponseError `json:"errors"`
	OfferID     string                                     `json:"offer_id"`
	ProductID   int64                                      `json:"product_id"`
	Updated     bool                                       `json:"updated"`
	WarehouseID int                                        `json:"warehouse_id"`
}

type ProductsStocksResponseProcessResultWrapper struct {
	Result []ProductsStocksResponseProcessResult `json:"result"`
}

type ProductStocksData struct {
	OfferID     *string `json:"offer_id,omitempty"`
	ProductID   *int64  `json:"product_id,omitempty"`
	Stock       *int    `json:"stock,omitempty"`
	WarehouseID *int    `json:"warehouse_id,omitempty"`
}

type StocksData struct {
	Stocks []ProductStocksData `json:"stocks,omitempty"`
}

type SetProductStocksResponseResult struct {
	OfferID     string `json:"offer_id"`
	ProductID   int64  `json:"product_id"`
	Updated     bool   `json:"updated"`
	WarehouseID int    `json:"warehouse_id"`
}

type SetProductStocksResponseResultWrapper struct {
	Result []SetProductStocksResponseResult `json:"result"`
}
