package brokers

import "net/url"

type AuthConfiguration interface {
	Token()
}

type BrokerConfiguration struct {
	Name string         `yaml:"name"`
	Url  url.URL        `yaml:"url"`
	Auth map[string]any `yaml:"auth"`
}
