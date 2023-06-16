package internal

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/Sackbuoy/gupTrade/internal/trader"
)

/* Builds configuration object from given yaml. */
func BuildConfiguration(viperRef *viper.Viper, configuration *Configuration) error {
	viperRef.SetConfigName("configuration")
	viperRef.SetConfigType("yaml")
	viperRef.AddConfigPath(".")

	viperRef.AutomaticEnv()

	err := viperRef.ReadInConfig()
	if err != nil {
		return err
	}

	err = viperRef.Unmarshal(configuration)
	if err != nil {
		return err
	}

	return nil
}

/*
Creates a Trader for each broker in the configuration.
*/
func SetupTraders(traderConfigs []trader.TraderConfiguration, logger *zap.SugaredLogger) ([]trader.Trader, error) {
	var traders []trader.Trader

	for i := range traderConfigs {
		t, err := trader.NewTrader(traderConfigs[i], logger)
		if err != nil {
			return nil, err
		}
		traders = append(traders, *t)
	}

	return traders, nil
}
