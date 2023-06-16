package brokers

import (
	"github.com/go-resty/resty/v2"
)

type TradeResponse struct{}

type Order struct {
	Symbol      string
	Type        string // enums....
	IsCompleted bool
	// TODO: other fields...
}

type BrokerClient interface {
	PlaceOrder() (*resty.Response, error)
	// SubmitOrder()
	// GetOrder()
	// CancelOrder()
	// TODO: other stuff...
}
