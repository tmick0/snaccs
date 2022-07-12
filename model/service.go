package model

type (
	Service struct {
		Name       string              `yaml:"name"`
		Dns        string              `yaml:"dns"`
		Backend    string              `yaml:"backend"`
		SelfSigned bool                `yaml:"selfSigned"`
		Allow      []IdentifierWrapper `yaml:"allow"`
	}
)
