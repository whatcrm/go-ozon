package models

import "time"

type SellerResponse struct {
	Company struct {
		Country       string `json:"country"`
		Currency      string `json:"currency"`
		Inn           string `json:"inn"`
		LegalName     string `json:"legal_name"`
		Name          string `json:"name"`
		Ogrn          string `json:"ogrn"`
		OwnershipForm string `json:"ownership_form"`
		TaxSystem     string `json:"tax_system"`
	} `json:"company"`
	Ratings []struct {
		CurrentValue struct {
			DateFrom  time.Time `json:"date_from"`
			DateTo    time.Time `json:"date_to"`
			Formatted string    `json:"formatted"`
			Status    struct {
				Danger  bool `json:"danger"`
				Premium bool `json:"premium"`
				Warning bool `json:"warning"`
			} `json:"status"`
			Value int `json:"value"`
		} `json:"current_value"`
		Name      string `json:"name"`
		PastValue struct {
			DateFrom  time.Time `json:"date_from"`
			DateTo    time.Time `json:"date_to"`
			Formatted string    `json:"formatted"`
			Status    struct {
				Danger  bool `json:"danger"`
				Premium bool `json:"premium"`
				Warning bool `json:"warning"`
			} `json:"status"`
			Value int `json:"value"`
		} `json:"past_value"`
		Rating    string `json:"rating"`
		Status    string `json:"status"`
		ValueType string `json:"value_type"`
	} `json:"ratings"`
	Subscription struct {
		IsPremium bool   `json:"is_premium"`
		Type      string `json:"type"`
	} `json:"subscription"`
}

type LogisticsResponse struct {
	AvailableSchemas     []string `json:"available_schemas"`
	OzonLogisticsEnabled bool     `json:"ozon_logistics_enabled"`
}
