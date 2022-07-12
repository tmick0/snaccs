package model

import (
	"errors"
)

type IdentifierClass int64

const (
	NetworkIdentifierClass IdentifierClass = 0
	HostIdentifierClass    IdentifierClass = 1
)

type (
	Identifier interface {
		Class() IdentifierClass
		Reference() string
	}

	HostIdentifier struct {
		Hostname string `yaml:"hostname"`
	}

	NetworkIdentifier struct {
		Network string `yaml:"network"`
	}

	IdentifierWrapper struct {
		Identifier Identifier
	}
)

func (id *HostIdentifier) UnmarshalYAML(unmarshal func(interface{}) error) error {
	unmarshal(&id.Hostname)
	return nil
}

func (id *NetworkIdentifier) UnmarshalYAML(unmarshal func(interface{}) error) error {
	unmarshal(&id.Network)
	return nil
}

func (id *IdentifierWrapper) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tmp map[string]string
	unmarshal(&tmp)
	for k := range tmp {
		switch k {
		case "hostname":
			id.Identifier = HostIdentifier{tmp["hostname"]}
			return nil
		case "network":
			id.Identifier = NetworkIdentifier{tmp["network"]}
			return nil
		default:
			goto err
		}
	}
err:
	return errors.New("Invalid network/host identifier")
}

func (c IdentifierClass) String() string {
	switch c {
	case HostIdentifierClass:
		return "host"
	case NetworkIdentifierClass:
		return "network"
	default:
		return "unknown"
	}
}

func (_ HostIdentifier) Class() IdentifierClass {
	return HostIdentifierClass
}

func (_ NetworkIdentifier) Class() IdentifierClass {
	return NetworkIdentifierClass
}

func (id IdentifierWrapper) Class() IdentifierClass {
	return id.Identifier.Class()
}

func (id HostIdentifier) Reference() string {
	return id.Hostname
}

func (id NetworkIdentifier) Reference() string {
	return id.Network
}

func (id IdentifierWrapper) Reference() string {
	return id.Identifier.Reference()
}
