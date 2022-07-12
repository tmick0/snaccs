package model

import (
	"net"
)

type (
	Cidr struct {
		Address Addr
		Size    int
	}
)

func (cidr *Cidr) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tmp string
	unmarshal(&tmp)
	ip, net, err := net.ParseCIDR(tmp)
	cidr.Address.ParseIP(ip)
	cidr.Size, _ = net.Mask.Size()
	return err
}
