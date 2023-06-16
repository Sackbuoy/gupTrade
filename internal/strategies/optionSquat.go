package strategies

import (
	"fmt"
	"sync"

	"go.uber.org/zap"

	"github.com/Sackbuoy/gupTrade/pkg/brokers"
)

type OptionSquat struct {
	client  brokers.BrokerClient
	logger  *zap.SugaredLogger
	symbols []string
}

func NewOptionSquat(config StrategyConfiguration, client brokers.BrokerClient, logger *zap.SugaredLogger) (OptionSquat, error) {
	return OptionSquat{
		client:  client,
		logger:  logger,
		symbols: config.Symbols,
	}, nil
}

func (s OptionSquat) Run(wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := s.client.PlaceOrder()
	if err != nil {
		s.logger.Infof("Failed to execute strategy: %s\n", err.Error())
	}

	fmt.Println(resp)
}
