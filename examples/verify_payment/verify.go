package main

import (
	"fmt"
	"log"

	payfonte "github.com/chuddyjoachim/payfonte-go"
)

func main() {
	payload := &payfonte.NewPayfonte{ClientId: "textng", ClientSecret: "dev_1325ece6062f9c87550e89431fb8d53963e4c48c6e5fa8742a"}
	api := payfonte.NewPayfonteApi(payload)

	fmt.Println("Verifying...")
	res, err := api.VerifyPayment(&payfonte.VerifyPaymentPayload{Reference: "D20240207174304EKWUG"})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n AmountLabel: %s\n", res.Data.AmountLabel)
	fmt.Printf("Status: %s\n", res.Data.Status)

}
