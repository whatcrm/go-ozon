package ozon

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/whatcrm/go-ozon/utils"
)

const DefaultBaseURL = "https://api-seller.ozon.ru"

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
	ClientID   string
	APIKey     string
}

func NewClient(clientID, apiKey string) (*Client, error) {
	if clientID == "" {
		return nil, errors.New("client id is required")
	}

	if apiKey == "" {
		return nil, errors.New("api key is required")
	}

	return &Client{
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second},
		BaseURL:  DefaultBaseURL,
		ClientID: clientID,
		APIKey:   apiKey,
	}, nil
}

func (c *Client) Send(req *http.Request, v interface{}) error {
	if req.Header.Get("Content-Type") == "" && req.Method != "DELETE" {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Client-Id", c.ClientID)
	req.Header.Set("Api-Key", c.APIKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "go-ozon-sdk")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return utils.ParseAPIError(resp)
	}

	if v == nil {
		return nil
	}

	if w, ok := v.(io.Writer); ok {
		_, err = io.Copy(w, resp.Body)
		return err
	}

	return json.NewDecoder(resp.Body).Decode(v)
}
