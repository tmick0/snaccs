package generator

import "github.com/tmick0/snaccs/model"

type (
	ConfigGenerator interface {
		Generate(model.Configuration) ([]byte, error)
	}
)
