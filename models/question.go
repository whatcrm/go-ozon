package models

import (
	"bytes"
	"encoding/json"
	"time"
)

type CreateQuestionAnswerRequest struct {
	QuestionID string `json:"question_id"`
	SKU        int64  `json:"sku"`
	Text       string `json:"text"`
}

type CreateQuestionAnswerResponse struct {
	AnswerID string `json:"answer_id"`
}

type DeleteQuestionAnswerRequest struct {
	AnswerID string `json:"answer_id"`
	SKU      int64  `json:"sku"`
}

type ListQuestionAnswersRequest struct {
	LastID     string `json:"last_id,omitempty"`
	QuestionID string `json:"question_id"`
	SKU        int64  `json:"sku"`
}

type QuestionAnswer struct {
	AuthorName        string    `json:"author_name"`
	ID                string    `json:"id"`
	PublishedAt       time.Time `json:"published_at"`
	QuestionID        string    `json:"question_id"`
	SKU               int64     `json:"sku"`
	StatusPublication string    `json:"status_publication"`
	Text              string    `json:"text"`
}

type ListQuestionAnswersResponse struct {
	Answers []QuestionAnswer `json:"answers"`
	LastID  string           `json:"last_id"`
}

type ChangeQuestionStatusRequest struct {
	QuestionIDs []string `json:"question_ids"`
	Status      string   `json:"status"`
}

type GetQuestionCountResponse struct {
	All         int64 `json:"all"`
	New         int64 `json:"new"`
	Processed   int64 `json:"processed"`
	Unprocessed int64 `json:"unprocessed"`
	Viewed      int64 `json:"viewed"`
}

type GetQuestionInfoRequest struct {
	QuestionID string `json:"question_id"`
}

type QuestionAnswersCount int64

func (c *QuestionAnswersCount) UnmarshalJSON(b []byte) error {
	b = bytes.TrimSpace(b)
	if len(b) == 0 || bytes.Equal(b, []byte("null")) {
		*c = 0
		return nil
	}

	if b[0] == '"' {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		var n json.Number = json.Number(s)
		v, err := n.Int64()
		if err != nil {
			return err
		}
		*c = QuestionAnswersCount(v)
		return nil
	}

	var v int64
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	*c = QuestionAnswersCount(v)
	return nil
}

type GetQuestionInfoResponse struct {
	AnswersCount QuestionAnswersCount `json:"answers_count"`
	AuthorName   string               `json:"author_name"`
	ID           string               `json:"id"`
	ProductURL   string               `json:"product_url"`
	PublishedAt  time.Time            `json:"published_at"`
	QuestionLink string               `json:"question_link"`
	SKU          int64                `json:"sku"`
	Status       string               `json:"status"`
	Text         string               `json:"text"`
}

type QuestionListFilter struct {
	DateFrom string `json:"date_from,omitempty"`
	DateTo   string `json:"date_to,omitempty"`
	Status   string `json:"status,omitempty"`
}

type ListQuestionsRequest struct {
	Filter  *QuestionListFilter `json:"filter,omitempty"`
	LastID  string              `json:"last_id,omitempty"`
	Limit   int64               `json:"limit,omitempty"`
	SortDir string              `json:"sort_dir,omitempty"`
}

type QuestionListItem struct {
	AnswersCount int64     `json:"answers_count"`
	AuthorName   string    `json:"author_name"`
	ID           string    `json:"id"`
	SKU          int64     `json:"sku"`
	ProductURL   string    `json:"product_url"`
	PublishedAt  time.Time `json:"published_at"`
	QuestionLink string    `json:"question_link"`
	Text         string    `json:"text"`
	Status       string    `json:"status"`
}

type ListQuestionsResponse struct {
	HasNext   bool               `json:"has_next"`
	LastID    string             `json:"last_id"`
	Questions []QuestionListItem `json:"questions"`
}

type GetQuestionTopSKURequest struct {
	Limit int64 `json:"limit"`
}

type GetQuestionTopSKUResponse struct {
	SKU []int64 `json:"sku"`
}
