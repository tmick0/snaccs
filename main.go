package main

import (
	"io/ioutil"
	"log"
	"os"

	model "github.com/tmick0/snaccs/model"
	yaml "gopkg.in/yaml.v2"
)

func main() {
	log.Println("hello world")
	var nets model.Configuration

	if len(os.Args) < 2 {
		log.Printf("usage: %s <config.yml>\n", os.Args[0])
		os.Exit(1)
	}

	bytes, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		log.Printf("failed to read config: %z\n", err)
	}

	err = yaml.Unmarshal(bytes, &nets)

	if err != nil {
		log.Printf("failed to deserialize yaml: %s\n", err)
		os.Exit(1)
	}

	log.Println(nets)

}
