package models

import "time"

const (
	NotificationTypeNewMessage                    = "TYPE_NEW_MESSAGE"
	NotificationTypeUpdateMessage                 = "TYPE_UPDATE_MESSAGE"
	NotificationTypeMessageRead                   = "TYPE_MESSAGE_READ"
	NotificationTypeChatClosed                    = "TYPE_CHAT_CLOSED"
	NotificationTypeNewPosting                    = "TYPE_NEW_POSTING"
	NotificationTypePostingCancelled              = "TYPE_POSTING_CANCELLED"
	NotificationTypeStateChanged                  = "TYPE_STATE_CHANGED"
	NotificationTypeDeliveryDateChanged           = "TYPE_DELIVERY_DATE_CHANGED"
	NotificationTypeCutoffDateChanged             = "TYPE_CUTOFF_DATE_CHANGED"
	NotificationTypeCreateItem                    = "TYPE_CREATE_ITEM"
	NotificationTypeUpdateItem                    = "TYPE_UPDATE_ITEM"
	NotificationTypeCreateOrUpdateItem            = "TYPE_CREATE_OR_UPDATE_ITEM"
	NotificationTypeStocksChanged                 = "TYPE_STOCKS_CHANGED"
	NotificationTypeFBOPostingNew                 = "TYPE_FBO_POSTING_NEW"
	NotificationTypeFBOPostingCancelled           = "TYPE_FBO_POSTING_CANCELLED"
	NotificationTypeFBOPostingStateChanged        = "TYPE_FBO_POSTING_STATE_CHANGED"
	NotificationTypeFBOPostingDeliveryDateChanged = "TYPE_FBO_POSTING_DELIVERY_DATE_CHANGED"
	NotificationTypeFBOStocksChanged              = "TYPE_FBO_STOCKS_CHANGED"
	NotificationTypeOrderNew                      = "TYPE_ORDER_NEW"
	NotificationTypeOrderCancelled                = "TYPE_ORDER_CANCELLED"
	NotificationTypeOrderStateChanged             = "TYPE_ORDER_STATE_CHANGED"

	NotificationAvailabilityGreen  = "GREEN"
	NotificationAvailabilityYellow = "YELLOW"
	NotificationAvailabilityRed    = "RED"
)

type SetNotificationRequest struct {
	Types []string `json:"types"`
	URL   string   `json:"url"`
}

type UpdateNotificationRequest struct {
	ID    int64    `json:"id"`
	Types []string `json:"types,omitempty"`
	URL   string   `json:"url,omitempty"`
}

type DeleteNotificationRequest struct {
	ID int64 `json:"id"`
}

type CheckNotificationRequest struct {
	URL string `json:"url"`
}

type CheckNotificationError struct {
	Description string `json:"description"`
	Type        string `json:"type"`
}

type CheckNotificationResponse struct {
	Errors   []CheckNotificationError `json:"errors"`
	IsActive bool                     `json:"is_active"`
}

type EnableNotificationRequest struct {
	Enabled bool  `json:"enabled"`
	ID      int64 `json:"id"`
}

type ListNotificationsRequest struct {
	AvailabilityStatuses []string `json:"availability_statuses,omitempty"`
	Limit                uint64   `json:"limit,omitempty"`
	Offset               uint64   `json:"offset,omitempty"`
	SortDir              string   `json:"sort_dir,omitempty"`
}

type NotificationAvailabilityThreshold struct {
	Status    string `json:"status"`
	Threshold int64  `json:"threshold"`
}

type NotificationURLType struct {
	Description string `json:"description"`
	Type        string `json:"type"`
}

type NotificationURL struct {
	AvailabilityStatus     string                `json:"availability_status"`
	AvailabilityStatusDate time.Time             `json:"availability_status_date"`
	CreatedAt              time.Time             `json:"created_at"`
	DisableReason          string                `json:"disable_reason"`
	Enable                 bool                  `json:"enable"`
	ID                     int64                 `json:"id"`
	ProblematicType        string                `json:"problematic_type"`
	ReasonDetails          string                `json:"reason_details"`
	Types                  []NotificationURLType `json:"types"`
	URL                    string                `json:"url"`
}

type ListNotificationsResponse struct {
	AvailabilityStatusThresholds []NotificationAvailabilityThreshold `json:"availability_status_thresholds"`
	TotalCount                   int64                               `json:"total_count"`
	URLs                         []NotificationURL                   `json:"urls"`
}

type NotificationSellerEndpoint struct {
	ID  int64  `json:"id"`
	URL string `json:"url"`
}

type NotificationPushType struct {
	Description    string                      `json:"description"`
	SellerEndpoint *NotificationSellerEndpoint `json:"seller_endpoint,omitempty"`
	Type           string                      `json:"type"`
}

type ListNotificationPushTypesResponse struct {
	Types []NotificationPushType `json:"types"`
}
