package types

type NewPayfonte struct {
	// Payfonte client id
	ClientId string

	// Payfonte client id
	ClientSecret string

	// isProd: isProject a secrets and id from a production account
	IsProd bool
}
