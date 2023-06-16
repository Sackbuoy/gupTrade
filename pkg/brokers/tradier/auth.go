package tradier

type TradierAuth struct {
	accessToken   string
	accountNumber string
}

// right now they just straight up gave me an access Token which is all i need for the api...so idk.
func (t *TradierAuth) Token() {}
