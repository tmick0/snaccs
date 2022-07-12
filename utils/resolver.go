package utils

import (
	"errors"
	"fmt"

	"github.com/tmick0/snaccs/model"
)

func ResolveIdentifierToIpRange(cfg model.Configuration, id model.Identifier) (string, error) {
	switch id.Class() {
	case model.HostIdentifierClass:
		for _, host := range cfg.Hosts {
			if host.Hostname == id.Reference() {
				res := host.Static.Ip.String()
				return res, nil
			}
		}
	case model.NetworkIdentifierClass:
		for _, net := range cfg.Networks {
			if net.Name == id.Reference() {
				res := net.Subnet.String()
				return res, nil
			}
		}
	}
	return "", errors.New(fmt.Sprintf("Could not resolve %s key %s", id.Class().String(), id.Reference()))
}
