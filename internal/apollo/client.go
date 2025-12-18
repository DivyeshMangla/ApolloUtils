package apollo

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

// Client wraps Apollo.io API interactions with authentication and timeout handling.
type Client struct {
	apiKey string
	http   *http.Client
}

// New creates an Apollo API client with the provided API key and 15-second timeout.
func New(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		http: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

// post sends a JSON POST request to the Apollo API and unmarshals the response.
func (c *Client) post(endpoint string, payload any, out any) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		"POST",
		"https://api.apollo.io"+endpoint,
		bytes.NewBuffer(body),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key", c.apiKey)

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(out)
}
