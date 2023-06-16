package tradier

import (
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"
)

// must implement BrokerClient interface.
type TradierClient struct {
	Auth  TradierAuth
	resty *resty.Client
}

func NewClient(auth map[string]any, url *url.URL) TradierClient {
	accessToken := auth["accesstoken"].(string)
	accountNumber := auth["accountnumber"].(string)

	tradierAuth := TradierAuth{
		accessToken,
		accountNumber,
	}

	if url.String() == "" {
		url, _ = url.Parse("https://sandbox.tradier.com")
	}

	return TradierClient{
		Auth:  tradierAuth,
		resty: resty.New().SetBaseURL(url.String()),
	}
}

// WIP.
func (t TradierClient) PlaceOrder() (*resty.Response, error) {
	t.resty.SetAuthToken(t.Auth.accessToken)
	request := t.resty.R().
		SetHeader("Accept", "application/json").
		SetQueryParam("includeTags", "true")

	response, err := request.Get(fmt.Sprintf("/v1/accounts/%s/orders", t.Auth.accountNumber))
	if err != nil {
		return nil, err
	}

	return response, nil
}
