package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/tmick0/snaccs/generator"
	"github.com/tmick0/snaccs/model"
	"gopkg.in/yaml.v2"
)

func main() {
	var cfg model.Configuration

	if len(os.Args) < 2 {
		log.Printf("usage: %s <config.yml>\n", os.Args[0])
		os.Exit(1)
	}

	bytes, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		log.Printf("failed to read config: %z\n", err)
	}

	err = yaml.Unmarshal(bytes, &cfg)

	if err != nil {
		log.Printf("failed to deserialize yaml: %s\n", err)
		os.Exit(1)
	}

	log.Println(cfg)

	res, err := generator.TraefikConfigGenerator{}.Generate(cfg)

	if err != nil {
		log.Printf("failed to create traefik config: %s\n", err)
		os.Exit(1)
	}

	log.Println(string(res))

}
