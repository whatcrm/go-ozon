package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type APIError struct {
	StatusCode int

	Code    int    `json:"code"`
	Message string `json:"message"`
	Details []struct {
		TypeUrl string `json:"typeUrl"`
		Value   string `json:"value"`
	} `json:"details"`

	RawBody    []byte
	RetryAfter string
}

func (e *APIError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return "OZON SELLER API ERROR"
}
func ParseAPIError(resp *http.Response) error {
	raw, _ := io.ReadAll(resp.Body)

	apiErr := &APIError{
		StatusCode: resp.StatusCode,
		RawBody:    raw,
		RetryAfter: resp.Header.Get("Retry-After"),
	}

	_ = json.Unmarshal(raw, apiErr)

	if apiErr.Message == "" {
		apiErr.Message = strings.TrimSpace(string(raw))
	}

	return apiErr
}
