package models

import "time"

type RatingSummaryRequest struct{}

type RatingSummaryResponse struct {
	Groups               []RatingSummaryGroup        `json:"groups"`
	LocalizationIndex    []RatingSummaryLocalization `json:"localization_index"`
	PenaltyScoreExceeded bool                        `json:"penalty_score_exceeded"`
	Premium              bool                        `json:"premium"`
	PremiumPlus          bool                        `json:"premium_plus"`
}

type RatingSummaryGroup struct {
	GroupName string                   `json:"group_name"`
	Items     []RatingSummaryGroupItem `json:"items"`
}

type RatingSummaryGroupItem struct {
	Change          RatingSummaryChange `json:"change"`
	CurrentValue    float64             `json:"current_value"`
	Name            string              `json:"name"`
	PastValue       float64             `json:"past_value"`
	Rating          string              `json:"rating"`
	RatingDirection string              `json:"rating_direction"`
	Status          string              `json:"status"`
	ValueType       string              `json:"value_type"`
}

type RatingSummaryChange struct {
	Direction string `json:"direction"`
	Meaning   string `json:"meaning"`
}

type RatingSummaryLocalization struct {
	CalculationDate        time.Time `json:"calculation_date"`
	LocalizationPercentage float64   `json:"localization_percentage"`
}

type RatingHistoryRequest struct {
	DateFrom          time.Time `json:"date_from"`
	DateTo            time.Time `json:"date_to"`
	Ratings           []string  `json:"ratings"`
	WithPremiumScores bool      `json:"with_premium_scores"`
}

type RatingHistoryResponse struct {
	PremiumScores []RatingHistoryPremiumScore `json:"premium_scores"`
	Ratings       []RatingHistoryRating       `json:"ratings"`
}

type RatingHistoryPremiumScore struct {
	Rating string                          `json:"rating"`
	Scores []RatingHistoryPremiumScoreItem `json:"scores"`
}

type RatingHistoryPremiumScoreItem struct {
	Date        time.Time `json:"date"`
	RatingValue float64   `json:"rating_value"`
	Value       float64   `json:"value"`
}

type RatingHistoryRating struct {
	DangerThreshold  float64                    `json:"danger_threshold"`
	PremiumThreshold float64                    `json:"premium_threshold"`
	Rating           string                     `json:"rating"`
	Values           []RatingHistoryRatingValue `json:"values"`
	WarningThreshold float64                    `json:"warning_threshold"`
}

type RatingHistoryRatingValue struct {
	DateFrom time.Time                 `json:"date_from"`
	DateTo   time.Time                 `json:"date_to"`
	Status   RatingHistoryRatingStatus `json:"status"`
	Value    float64                   `json:"value"`
}

type RatingHistoryRatingStatus struct {
	Danger  bool `json:"danger"`
	Premium bool `json:"premium"`
	Warning bool `json:"warning"`
}

type RatingFbsIndexInfoResponse struct {
	CurrencyCode       string                 `json:"currency_code"`
	Defects            []RatingFbsIndexDefect `json:"defects"`
	Index              float64                `json:"index"`
	PeriodFrom         string                 `json:"period_from"`
	PeriodTo           string                 `json:"period_to"`
	ProcessingCostsSum float64                `json:"processing_costs_sum"`
}

type RatingFbsIndexDefect struct {
	Date                     string  `json:"date"`
	IndexByDate              float64 `json:"index_by_date"`
	ProcessingCostsSumByDate float64 `json:"processing_costs_sum_by_date"`
}

type RatingFbsIndexPostingListRequest struct {
	Cursor string                          `json:"cursor"`
	Filter RatingFbsIndexPostingListFilter `json:"filter"`
	Limit  int64                           `json:"limit"`
}

type RatingFbsIndexPostingListFilter struct {
	DateFrom       time.Time `json:"date_from"`
	DateTo         time.Time `json:"date_to"`
	PostingNumbers []string  `json:"posting_numbers"`
}

type RatingFbsIndexPostingListResponse struct {
	Cursor  string                           `json:"cursor"`
	Errors  []RatingFbsIndexPostingListError `json:"errors"`
	HasNext bool                             `json:"has_next"`
}

type RatingFbsIndexPostingListError struct {
	ChargePercent            float64   `json:"charge_percent"`
	ChargePrice              float64   `json:"charge_price"`
	ChargePriceCurrencyCode  string    `json:"charge_price_currency_code"`
	DeliverySchema           string    `json:"delivery_schema"`
	ErrorAt                  time.Time `json:"error_at"`
	HasGraceStatus           bool      `json:"has_grace_status"`
	Index                    float64   `json:"index"`
	PostingErrorType         string    `json:"posting_error_type"`
	PostingNumber            string    `json:"posting_number"`
	ProductPrice             float64   `json:"product_price"`
	ProductPriceCurrencyCode string    `json:"product_price_currency_code"`
}
