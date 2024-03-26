package main

import (
	"encoding/json"
	"fmt"
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
		Name        string `json:"name"`
	} `json:"user"`
	Reference string `json:"reference"`
	Amount    int    `json:"amount"`
	Currency  string `json:"currency"`
}

type GenerateCheckoutResponseData struct {
	Id        string `json:"id"`
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

// verify
type VerifyPaymentPayload struct {
	Reference string `json:"reference"`
}

type VerifyPaymentResponseData struct {
	ClientId            string `json:"clientId"`
	Status              string `json:"status"`
	Reference           string `json:"reference"`
	ExternalReference   string `json:"externalReference"`
	ProvidersReference  string `json:"providersReference"`
	Currency            string `json:"currency"`
	Amount              int    `json:"amount"`
	AotalAmount         int    `json:"totalAmount"`
	PaidAt              int    `json:"paidAt"`
	ChargeBearer        string `json:"chargeBearer"`
	CheckoutRedirectURL string `json:"checkoutRedirectURL"`
	Provider            string `json:"provider"`
	IntegrationId       string `json:"integrationId"`
	UserId              string `json:"userId"`
	Id                  string `json:"id"`
	AmountLabel         string `json:"amountLabel"`
	ChargeLabel         string `json:"chargeLabel"`
	User                struct {
		Email       string `json:"email"`
		PhoneNumber string `json:"phoneNumber"`
	} `json:"user"`
}
type VerifyPaymentResponse struct {
	Data       VerifyPaymentResponseData `json:"data"`
	StatusCode int                       `json:"statusCode"`
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

// Verify Payment
func (p *PayfonteInitValues) VerifyPayment(payload *VerifyPaymentPayload) (*VerifyPaymentResponse, error) {

	cl := httpclient.NewHttpClient(&httpclient.PayfonteInitValues{ClientId: p.ClientId, ClientSecret: p.ClientSecret, BaseURL: p.BaseURL})

	path := "/payments/v1/payments/verify/"

	fullUrlPath := path + payload.Reference

	fmt.Println(p.BaseURL)
	fmt.Println(fullUrlPath)

	response, err := cl.Get(fullUrlPath)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseErr := lib.CheckHttpError(response)
	if responseErr != nil {
		return nil, responseErr
	}
	data := &VerifyPaymentResponse{}

	jsonErr := json.NewDecoder(response.Body).Decode(&data)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return data, nil
}
