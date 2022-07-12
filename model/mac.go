package model

type (
	Mac struct {
		Mac string
	}
)

func (mac *Mac) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tmp string
	unmarshal(&tmp)
	mac.Mac = tmp
	return nil
}
