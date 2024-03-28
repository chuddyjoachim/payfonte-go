package main

import (
	"fmt"
	"log"

	payfonte "github.com/chuddyjoachim/payfonte-go"
)

func main() {
	payload := &payfonte.NewPayfonte{ClientId: "textng", ClientSecret: "dev_1325ece6062f9c87550e89431fb8d53963e4c48c6e5fa8742a"}
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
