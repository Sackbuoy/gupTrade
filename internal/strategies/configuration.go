package strategies

// TODO: different strategies will have different configuration options...
type StrategyConfiguration struct {
	Name    string   `yaml:"name"`
	Symbols []string `yaml:"tickers"`
}
