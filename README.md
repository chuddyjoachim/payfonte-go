# Payfonte Go

payfonte-go contains the official API bindings for generating **Checkout URL** and **Verifying payment** for [payfonte.com](https://payfonte.com/).

## Installation

```sh
go get github.com/chuddyjoachim/payfonte-go@latest
```

## Configuration

- `ClientId`: Your **Payfonte** client ID.
- `ClientSecret`: Your **Payfonte** API secret key.
- `IsProd (optional) - defaults to false`: If you'd want to interact with the production or sandbox environment.

## Usage

This API contains 2 interactions:

1. [Generate checkout URL](#generate-checkout)
2. [Verify payment](#verify-payment)

### Generate checkout URL<a id="generate-checkout"></a>

To Generate checkout, use the `GenerateCheckoutUrl` function.

```go
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

```

### Verify payment<a id="verify-payment"></a>

Verify payment, use the `VerifyPayment` function.

```go
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

```

## API Documentation

[payfonte.readme.io](https://payfonte.readme.io/). Contains a full up-to-date API for interacting with payfontes service.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

