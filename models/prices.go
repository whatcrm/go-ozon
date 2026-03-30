package models

type ItemPriceData struct {
	AutoActionEnabled string `json:"auto_action_enabled,omitempty"`
	MinPrice          string `json:"min_price,omitempty"`
	OfferID           string `json:"offer_id,omitempty"`
	OldPrice          string `json:"old_price,omitempty"`
	Price             string `json:"price,omitempty"`
	ProductID         int64  `json:"product_id,omitempty"`
}

type PricesData struct {
	Prices []ItemPriceData `json:"prices,omitempty"`
}

type GetProductImportPriceResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type GetProductImportPriceResponseResult struct {
	Errors    []GetProductImportPriceResponseError `json:"errors"`
	OfferID   string                               `json:"offer_id"`
	ProductID int64                                `json:"product_id"`
	Updated   bool                                 `json:"updated"`
}

type GetProductImportPriceResponseResultWrapper struct {
	Result []GetProductImportPriceResponseResult `json:"result"`
}
