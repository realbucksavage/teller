package main

import (
	"flag"
	"log"
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

	configFile, err := os.Open(configLocation)
	if err != nil {
		panic(err)
	}

	defer configFile.Close()

	var cfg Configuration
	if err := yaml.NewDecoder(configFile).Decode(&cfg); err != nil {
		panic(err)
	}

	src, err := sources.Init(cfg.Configuration)
	if err != nil {
		panic(err)
	}

	if err := src.Start(); err != nil {
		panic(err)
	}

	defer src.Close()

	errChan := make(chan error)
	src.StartRefresh(errChan)
	go func() {
		for refreshError := range errChan {
			log.Println(refreshError)
		}
	}()
}
