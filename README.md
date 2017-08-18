# Paypal Marketplace [![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/greater-commons/paypal-marketplace)

Marketplace is still in development, it is not ready for production use.

## Usage
All functions and types follow the REST API closely.

```go
c := market.NewClient(ctx, clientID, secret, market.Live)
params := &partner.CreatePartnerReferralParams{
	// ...
}
_, err := c.CreatePartnerReferral(ctx, params)
```

## Contributers
Pull requests are welcome.

- Code must have been run through `go fmt`.

Instructions on how to use Github pull requests can be found [here](https://youtu.be/iYIWwob0wKg)
