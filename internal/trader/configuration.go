package trader

import (
	"github.com/Sackbuoy/gupTrade/internal/strategies"
	"github.com/Sackbuoy/gupTrade/pkg/brokers"
)

type TraderConfiguration struct {
	Name       string                             `yaml:"name"`
	Broker     brokers.BrokerConfiguration        `yaml:"broker"`
	Strategies []strategies.StrategyConfiguration `yaml:"strategies"`
}
