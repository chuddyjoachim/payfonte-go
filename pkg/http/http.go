package httpclient

import (
	"bytes"
	"encoding/json"
	"net/http"

	types "github.com/chuddyjoachim/payfonte-go/pkg/types"
)

type PayfonteInitValues types.PayfonteInitValues

// NewHttpClient new instance of PayfonteInitValues.
func NewHttpClient(v *PayfonteInitValues) *PayfonteInitValues {
	return v
}

// Post sends a POST request to the specified endpoint with the given payload.
func (c *PayfonteInitValues) Post(endpoint string, payload interface{}) (*http.Response, error) {
	url := c.BaseURL + endpoint

	// Marshal payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("client-id", c.ClientId)
	req.Header.Set("client-secret", c.ClientSecret)

	client := http.DefaultClient
	return client.Do(req)
}

// Get:- sends a Get request to the specified endpoint.
func (c *PayfonteInitValues) Get(endpoint string) (*http.Response, error) {
	url := c.BaseURL + endpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("client-id", c.ClientId)

	client := http.DefaultClient
	return client.Do(req)
}
