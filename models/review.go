package models

import "time"

type CreateReviewCommentRequest struct {
	MarkReviewAsProcessed bool   `json:"mark_review_as_processed,omitempty"`
	ParentCommentID       string `json:"parent_comment_id,omitempty"`
	ReviewID              string `json:"review_id"`
	Text                  string `json:"text"`
}

type CreateReviewCommentResponse struct {
	CommentID string `json:"comment_id"`
}

type DeleteReviewCommentRequest struct {
	CommentID string `json:"comment_id"`
	SKU       int64  `json:"sku"`
}

type DeleteReviewCommentV1Request struct {
	CommentID string `json:"comment_id"`
}

type ReviewCommentListFilter struct {
	SKU           int64  `json:"sku,omitempty"`
	PublishedFrom string `json:"published_from,omitempty"`
	PublishedTo   string `json:"published_to,omitempty"`
}

type ListReviewCommentsRequest struct {
	ReviewID string                   `json:"review_id"`
	Filter   *ReviewCommentListFilter `json:"filter,omitempty"`
	Limit    int32                    `json:"limit"`
	Offset   int32                    `json:"offset,omitempty"`
	SortDir  string                   `json:"sort_dir,omitempty"`
}

type ReviewComment struct {
	ID              string    `json:"id"`
	Text            string    `json:"text"`
	ParentCommentID string    `json:"parent_comment_id"`
	IsOwner         bool      `json:"is_owner"`
	IsOfficial      bool      `json:"is_official"`
	IsPublished     bool      `json:"is_published"`
	IsRejected      bool      `json:"is_rejected"`
	LikesAmount     int32     `json:"likes_amount"`
	DislikesAmount  int32     `json:"dislikes_amount"`
	DeviationReason string    `json:"deviation_reason"`
	PublishedAt     time.Time `json:"published_at"`
}

type ListReviewCommentsResponse struct {
	Comments []ReviewComment `json:"comments"`
	Offset   int32           `json:"offset"`
}

type ChangeReviewStatusRequest struct {
	ReviewIDs []string `json:"review_ids"`
	Status    string   `json:"status"`
}

type ChangeReviewStatusV1Request struct {
	ReviewIDs []string `json:"review_ids"`
	Status    string   `json:"status"`
}

type GetReviewCountResponse struct {
	New       int32 `json:"new"`
	Processed int32 `json:"processed"`
	Total     int32 `json:"total"`
	Viewed    int32 `json:"viewed"`
}

type GetReviewCountV1Response struct {
	Processed   int32 `json:"processed"`
	Total       int32 `json:"total"`
	Unprocessed int32 `json:"unprocessed"`
}

type GetReviewInfoRequest struct {
	ReviewID string `json:"review_id"`
}

type ReviewPhoto struct {
	Height int32  `json:"height"`
	URL    string `json:"url"`
	Width  int32  `json:"width"`
}

type ReviewVideo struct {
	Height               int32  `json:"height"`
	PreviewURL           string `json:"preview_url"`
	ShortVideoPreviewURL string `json:"short_video_preview_url"`
	URL                  string `json:"url"`
	Width                int32  `json:"width"`
}

type GetReviewInfoResponse struct {
	CommentsAmount      int32         `json:"comments_amount"`
	DislikesAmount      int32         `json:"dislikes_amount"`
	ID                  string        `json:"id"`
	IsRatingParticipant bool          `json:"is_rating_participant"`
	LikesAmount         int32         `json:"likes_amount"`
	OrderStatus         string        `json:"order_status"`
	Photos              []ReviewPhoto `json:"photos"`
	PhotosAmount        int32         `json:"photos_amount"`
	PublishedAt         time.Time     `json:"published_at"`
	Rating              int32         `json:"rating"`
	SKU                 int64         `json:"sku"`
	Status              string        `json:"status"`
	Text                string        `json:"text"`
	Videos              []ReviewVideo `json:"videos"`
	VideosAmount        int32         `json:"videos_amount"`
}

type ReviewListFilters struct {
	SKU           []int64 `json:"sku,omitempty"`
	OrderStatus   string  `json:"order_status,omitempty"`
	Status        string  `json:"status,omitempty"`
	PublishedFrom string  `json:"published_from,omitempty"`
	PublishedTo   string  `json:"published_to,omitempty"`
}

type ListReviewsRequest struct {
	Filters *ReviewListFilters `json:"filters,omitempty"`
	LastID  string             `json:"last_id,omitempty"`
	Limit   int32              `json:"limit"`
	SortDir string             `json:"sort_dir,omitempty"`
}

type ReviewListItem struct {
	ID                  string    `json:"id"`
	SKU                 string    `json:"sku"`
	Text                string    `json:"text"`
	PublishedAt         time.Time `json:"published_at"`
	Rating              int32     `json:"rating"`
	Status              string    `json:"status"`
	CommentsAmount      int32     `json:"comments_amount"`
	PhotosAmount        int32     `json:"photos_amount"`
	VideosAmount        int32     `json:"videos_amount"`
	OrderStatus         string    `json:"order_status"`
	IsRatingParticipant bool      `json:"is_rating_participant"`
}

type ListReviewsResponse struct {
	HasNext bool             `json:"has_next"`
	LastID  string           `json:"last_id"`
	Reviews []ReviewListItem `json:"reviews"`
}
