package main

import (
	"fmt"
	"log"

	payfonte "github.com/chuddyjoachim/payfonte-go"
)

func main() {
	cl := "YOUR_CLIENTID"
	secret := "YOUR_CLIENTSECRET"
	IsProd := false //set isProd to true when in production environment

	payload := &payfonte.NewPayfonte{ClientId: cl, ClientSecret: secret, IsProd: IsProd}
	api := payfonte.NewPayfonteApi(payload)

	pl := &payfonte.GenerateCheckoutPayload{
		User: struct {
			Email       string "json:\"email\""
			Name        string "json:\"name\""
			PhoneNumber string "json:\"phoneNumber,omitempty\""
		}{Email: "example@example.com",
			Name: "Joachim test"},
		Reference: "ODXnncuence",
		Amount:    120,
		Currency:  "NGN",
	}

	fmt.Println("Generating checkout...")
	data, error := api.GenerateCheckoutUrl(pl)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Printf("Generated checkout Id: %s\n", data.Data.Id)
	fmt.Printf("Payment reference: %s\n", data.Data.Reference)
	fmt.Printf("Payment URL: %s\n", data.Data.ShortUrl)

}
