package main

import (
	"flag"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/realbucksavage/teller/sources"
)

type Configuration struct {
	sources.Configuration `yaml:",inline"`
	HTTPAddr              string `yaml:"bind_http"`
}

func main() {

	var configLocation string
	flag.StringVar(&configLocation, "cfg", "config.yaml", "Program's configuration file")
	flag.Parse()

	configSource, err := os.Open(configLocation)
	if err != nil {
		panic(err)
	}

	defer configSource.Close()

	var cfg Configuration
	if err := yaml.NewDecoder(configSource).Decode(&cfg); err != nil {
		panic(err)
	}

	_, err = sources.Init(cfg.Configuration)
	if err != nil {
		panic(err)
	}
}
