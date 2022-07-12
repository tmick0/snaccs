package model

type (
	Configuration struct {
		Snaccs   SnaccsConfig `yaml:"snaccs"`
		Networks []Network    `yaml:"networks"`
		Hosts    []Host       `yaml:"hosts"`
		Services []Service    `yaml:"services"`
	}
)
