package strategies

import (
	"fmt"
	"sync"

	"go.uber.org/zap"

	"github.com/Sackbuoy/gupTrade/pkg/brokers"
)

// every strategy needs to implement a simple Run() method.
type Strategy interface {
	Run(*sync.WaitGroup)
}

func NewStrategy(config StrategyConfiguration, client brokers.BrokerClient, logger *zap.SugaredLogger) (Strategy, error) {
	switch config.Name {
	case "optionSquat":
		return NewOptionSquat(config, client, logger)
	default:
		return nil, fmt.Errorf("strategy named %s is not supported", config.Name)
	}
}
