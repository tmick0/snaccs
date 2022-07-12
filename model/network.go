package model

type (
	Network struct {
		Name      string `yaml:"name"`
		Interface string `yaml:"interface"`
		Subnet    Cidr   `yaml:"cidr"`
		Gateway   Addr   `yaml:"gateway"`
		Dns       Addr   `yaml:"dns"`
		Suffix    string `yaml:"suffix"`
		Range     Range  `yaml:"range"`
	}
)
