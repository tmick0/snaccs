package model

type (
	Host struct {
		Hostname string            `yaml:"hostname"`
		Mac      Mac               `yaml:"mac"`
		Network  NetworkIdentifier `yaml:"network"`
		Static   Addr              `yaml:"static"`
	}
)
