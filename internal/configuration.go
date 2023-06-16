package internal

import (
	"github.com/Sackbuoy/gupTrade/internal/trader"
)

// how to handle different auth types? I could just have a value for every possible auth lol.
type AuthConfiguration struct {
	AccountNumber string `yaml:"accountNumber"`
	AccessToken   string `yaml:"accessToken"`
}

type Configuration struct {
	Traders []trader.TraderConfiguration `yaml:"traders"`
}
