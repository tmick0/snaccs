package model

import (
	"net"
)

type (
	Addr struct {
		Ip net.IP
	}
)

func (addr *Addr) ParseIP(ip net.IP) {
	addr.Ip = ip
}

func (addr *Addr) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tmp string
	unmarshal(&tmp)
	addr.Ip = net.ParseIP(tmp)
	return nil
}
