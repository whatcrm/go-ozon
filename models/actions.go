package models

type GetSellerActionsResponseResult struct {
	ID                         float64 `json:"id"`
	Title                      string  `json:"title"`
	ActionType                 string  `json:"action_type"`
	Description                string  `json:"description"`
	DateStart                  string  `json:"date_start"`
	DateEnd                    string  `json:"date_end"`
	FreezeDate                 string  `json:"freeze_date"`
	PotentialProductsCount     float64 `json:"potential_products_count"`
	ParticipatingProductsCount float64 `json:"participating_products_count"`
	IsParticipating            bool    `json:"is_participating"`
	BannedProductsCount        float64 `json:"banned_products_count"`
	WithTargeting              bool    `json:"with_targeting"`
	OrderAmount                float64 `json:"order_amount"`
	DiscountType               string  `json:"discount_type"`
	DiscountValue              float64 `json:"discount_value"`
	IsVoucherAction            bool    `json:"is_voucher_action"`
}

type GetSellerActionsResponseResultWrapper struct {
	Result []GetSellerActionsResponseResult `json:"result"`
}

type PaginatedCandidatesForActions struct {
	ActionID *float64 `json:"action_id,omitempty"`
	Limit    *float64 `json:"limit,omitempty"`
	Offset   *float64 `json:"offset,omitempty"`
}

type GetActionsCandidatesResponseProducts struct {
	ID             float64 `json:"id"`
	Price          float64 `json:"price"`
	ActionPrice    float64 `json:"action_price"`
	MaxActionPrice float64 `json:"max_action_price"`
	AddMode        string  `json:"add_mode"`
	MinStock       float64 `json:"min_stock"`
	Stock          float64 `json:"stock"`
}

type GetActionsCandidatesResponseResult struct {
	Products []GetActionsCandidatesResponseProducts `json:"products"`
	Total    float64                                `json:"total"`
}

type GetActionsCandidatesResponseResultWrapper struct {
	Result GetActionsCandidatesResponseResult `json:"result"`
}

type PaginatedActionProducts struct {
	ActionID *float64 `json:"action_id,omitempty"`
	Limit    *float64 `json:"limit,omitempty"`
	Offset   *float64 `json:"offset,omitempty"`
}

type GetSellerProductResponseProducts struct {
	ID             int     `json:"id"`
	Price          float64 `json:"price"`
	ActionPrice    float64 `json:"action_price"`
	MaxActionPrice float64 `json:"max_action_price"`
	AddMode        string  `json:"add_mode"`
	MinStock       float64 `json:"min_stock"`
	Stock          float64 `json:"stock"`
}

type GetSellerProductResponseResult struct {
	Products []GetSellerProductResponseProducts `json:"products"`
	Total    int                                `json:"total"`
}

type GetSellerProductResponseResultWrapper struct {
	Result GetSellerProductResponseResult `json:"result"`
}
