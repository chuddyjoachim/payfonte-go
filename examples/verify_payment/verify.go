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

	rf := "YOUR_PAYMENT_REFERENCE"

	payload := &payfonte.NewPayfonte{ClientId: cl, ClientSecret: secret, IsProd: IsProd}
	api := payfonte.NewPayfonteApi(payload)

	fmt.Println("Verifying...")
	res, err := api.VerifyPayment(&payfonte.VerifyPaymentPayload{Reference: rf})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n AmountLabel: %s\n", res.Data.AmountLabel)
	fmt.Printf("Status: %s\n", res.Data.Status)

}
