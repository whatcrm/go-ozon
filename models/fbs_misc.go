package models

type FbsPackageLabelTaskCreateRequest struct {
	PostingNumber []string `json:"posting_number"`
}

type FbsPackageLabelTaskCreateResponse struct {
	Result FbsPackageLabelTaskResult `json:"result"`
}

type FbsPackageLabelTaskResult struct {
	TaskID int64 `json:"task_id"`
}

type FbsPackageLabelTaskGetRequest struct {
	TaskID int64 `json:"task_id"`
}

type FbsPackageLabelTaskGetResponse struct {
	Result FbsPackageLabelTaskGetResult `json:"result"`
}

type FbsPackageLabelTaskGetResult struct {
	Error                  string                `json:"error"`
	Status                 string                `json:"status"`
	FileURL                string                `json:"file_url"`
	PrintedPostingsCount   int32                 `json:"printed_postings_count"`
	UnprintedPostingsCount int32                 `json:"unprinted_postings_count"`
	UnprintedPostings      []FbsUnprintedPosting `json:"unprinted_postings"`
}

type FbsUnprintedPosting struct {
	PostingNumber string   `json:"posting_number"`
	Errors        []string `json:"errors,omitempty"`
}

type FbsCancelReasonForPostingRequest struct {
	RelatedPostingNumbers []string `json:"related_posting_numbers"`
}

type FbsCancelReasonForPostingResponse struct {
	Result []FbsCancelReasonForPostingItem `json:"result"`
}

type FbsCancelReasonForPostingItem struct {
	PostingNumber string                 `json:"posting_number"`
	Reasons       []FbsCancelReasonEntry `json:"reasons"`
}

type FbsCancelReasonEntry struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	TypeID string `json:"type_id"`
}

type FbsCancelReasonListResponse struct {
	Result []FbsCancelReasonListItem `json:"result"`
}

type FbsCancelReasonListItem struct {
	ID                         int64  `json:"id"`
	Title                      string `json:"title"`
	TypeID                     string `json:"type_id"`
	IsAvailableForCancellation bool   `json:"is_available_for_cancellation"`
}

type FbsProductCancelItem struct {
	Quantity int32 `json:"quantity"`
	SKU      int64 `json:"sku"`
}

type FbsProductCancelRequest struct {
	CancelReasonID      int64                  `json:"cancel_reason_id"`
	CancelReasonMessage string                 `json:"cancel_reason_message"`
	Items               []FbsProductCancelItem `json:"items"`
	PostingNumber       string                 `json:"posting_number"`
}

type FbsProductCancelResponse struct {
	Result string `json:"result"`
}

type FbsCancelRequest struct {
	CancelReasonID      int64  `json:"cancel_reason_id"`
	CancelReasonMessage string `json:"cancel_reason_message,omitempty"`
	PostingNumber       string `json:"posting_number"`
}

type FbsCancelResponse struct {
	Result bool `json:"result"`
}
