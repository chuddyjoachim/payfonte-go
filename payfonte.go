package main

import (
	"encoding/json"
	"io"

	httpclient "github.com/chuddyjoachim/payfonte-go/pkg/http"
	lib "github.com/chuddyjoachim/payfonte-go/pkg/lib"
	types "github.com/chuddyjoachim/payfonte-go/pkg/types"
)

type PayfonteInitValues types.PayfonteInitValues

type NewPayfonte types.NewPayfonte

type GenerateCheckoutPayload struct {
	User struct {
		Email       string `json:"email"`
		PhoneNumber string `json:"phoneNumber"`
	} `json:"user"`
	Reference string `json:"reference"`
	Amount    int    `json:"amount"`
	Currency  string `json:"currency"`
}

type GenerateCheckoutResponseData struct {
	ID        string `json:"id"`
	Url       string `json:"url"`
	ShortUrl  string `json:"shorturl"`
	Reference string `json:"reference"`
	Amount    int    `json:"amount"`
	Currency  string `json:"currency"`
}
type GenerateCheckoutResponse struct {
	Data       GenerateCheckoutResponseData `json:"data"`
	StatusCode int                          `json:"statusCode"`
}

// Initialize payfonte api
func NewPayfonteApi(p *NewPayfonte) *PayfonteInitValues {
	baseUrl := lib.GetBaseUrl(p.IsProd)

	return &PayfonteInitValues{
		ClientId:     p.ClientId,
		ClientSecret: p.ClientSecret,
		BaseURL:      baseUrl,
	}
}

// Generate checkout url
func (p *PayfonteInitValues) GenerateCheckoutUrl(payload *GenerateCheckoutPayload) (*GenerateCheckoutResponse, error) {

	cl := httpclient.NewHttpClient(&httpclient.PayfonteInitValues{ClientId: p.ClientId, ClientSecret: p.ClientSecret, BaseURL: p.BaseURL})

	response, err := cl.Post("/payments/v1/checkouts", payload)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseErr := lib.CheckHttpError(response)
	if responseErr != nil {
		return nil, responseErr
	}
	bodyBytes, _ := io.ReadAll(response.Body)
	data := &GenerateCheckoutResponse{}
	err = json.Unmarshal(bodyBytes, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
