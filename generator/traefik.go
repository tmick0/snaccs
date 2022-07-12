package generator

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"gopkg.in/yaml.v2"

	"github.com/tmick0/snaccs/model"
	"github.com/tmick0/snaccs/utils"
)

type (
	TraefikConfigGenerator struct {
	}

	traefikConfigRouter struct {
		Rule        string            `yaml:"rule"`
		Service     string            `yaml:"service"`
		EntryPoints []string          `yaml:"entryPoints"`
		Tls         map[string]string `yaml:"tls"`
		Middlewares []string          `yaml:"middlewares"`
	}

	traefikConfigServer struct {
		Url string `yaml:"url"`
	}

	traefikConfigLoadBalancer struct {
		Servers          []traefikConfigServer `yaml:"servers"`
		ServersTransport *string               `yaml:"serversTransport,omitempty"`
	}

	traefikConfigService struct {
		LoadBalancer traefikConfigLoadBalancer `yaml:"loadBalancer"`
	}

	traefikConfigServersTransport struct {
		InsecureSkipVerify bool `yaml:"insecureSkipVerify"`
	}

	traefikConfigWhitelist struct {
		SourceRanges []string
	}

	traefikConfigMiddleware struct {
		IpWhitelist traefikConfigWhitelist `yaml:"ipWhitelist"`
	}

	traefikConfigRoot struct {
		Http struct {
			Routers           map[string]*traefikConfigRouter           `yaml:"routers"`
			Services          map[string]*traefikConfigService          `yaml:"services"`
			ServersTransports map[string]*traefikConfigServersTransport `yaml:"serversTransports"`
			Middlewares       map[string]*traefikConfigMiddleware       `yaml:"middlewares"`
		} `yaml:"http"`
	}
)

func (c *traefikConfigRoot) init() {
	c.Http.Routers = make(map[string]*traefikConfigRouter)
	c.Http.Services = make(map[string]*traefikConfigService)
	c.Http.ServersTransports = make(map[string]*traefikConfigServersTransport)
	c.Http.Middlewares = make(map[string]*traefikConfigMiddleware)
}

func (c *traefikConfigLoadBalancer) addServer(url string) {
	c.Servers = append(c.Servers, traefikConfigServer{url})
}

func (c *traefikConfigRouter) addEntryPoint(entrypoint string) {
	c.EntryPoints = append(c.EntryPoints, entrypoint)
}

func (c *traefikConfigRouter) addMiddleware(middleware string) {
	c.Middlewares = append(c.Middlewares, middleware)
}

func (c *traefikConfigWhitelist) addRange(ipRange string) {
	c.SourceRanges = append(c.SourceRanges, ipRange)
}

func (c *traefikConfigRoot) addRouter(name string) *traefikConfigRouter {
	res := &traefikConfigRouter{}
	c.Http.Routers[name] = res
	res.Tls = make(map[string]string)
	return res
}

func (c *traefikConfigRoot) addService(name string) *traefikConfigService {
	res := &traefikConfigService{}
	c.Http.Services[name] = res
	return res
}

func (c *traefikConfigRoot) addMiddleware(name string) *traefikConfigMiddleware {
	res := &traefikConfigMiddleware{}
	c.Http.Middlewares[name] = res
	return res
}

func (c *traefikConfigRoot) addTransport(name string) *traefikConfigServersTransport {
	res := &traefikConfigServersTransport{}
	c.Http.ServersTransports[name] = res
	return res
}

func (_ TraefikConfigGenerator) Generate(cfg model.Configuration) ([]byte, error) {
	var traefikCfg traefikConfigRoot
	traefikCfg.init()

	for _, svcDef := range cfg.Services {
		uuid, err := uuid.NewRandom()
		if err != nil {
			return nil, err
		}
		key := strings.Replace(uuid.String(), "-", "", -1)

		routerCfg := traefikCfg.addRouter(key)
		routerCfg.Rule = fmt.Sprintf("Host(`%s`)", svcDef.Dns)
		routerCfg.Service = key
		routerCfg.addEntryPoint("websecure") // todo: get this from config
		routerCfg.addMiddleware(key)

		serviceCfg := traefikCfg.addService(key)
		serviceCfg.LoadBalancer.addServer(svcDef.Backend)

		middlewareCfg := traefikCfg.addMiddleware(key)
		for _, identifier := range svcDef.Allow {
			ipRange, err := utils.ResolveIdentifierToIpRange(cfg, identifier)
			if err != nil {
				return nil, err
			}
			middlewareCfg.IpWhitelist.addRange(ipRange)
		}

		if svcDef.SelfSigned {
			transportCfg := traefikCfg.addTransport(key)
			transportCfg.InsecureSkipVerify = true
			serviceCfg.LoadBalancer.ServersTransport = &key
		}
	}

	res, err := yaml.Marshal(traefikCfg)
	return res, err
}
