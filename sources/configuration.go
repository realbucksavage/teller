package sources

import (
	"fmt"
	"strings"

	"github.com/realbucksavage/teller"
)

var factories = map[string]teller.SourceFactory{}

func RegisterFactory(sType string, factory teller.SourceFactory) {
	sType = strings.ToLower(sType)
	fact, ok := factories[sType]
	if ok && fact != factory {
		panic(fmt.Sprintf("another %q factory is already registered", sType))
	}

	factories[sType] = factory
}

type Configuration struct {
	Sources []SourceDescriptor `yaml:"sources"`
}

type SourceDescriptor struct {
	Name     string                 `yaml:"name"`
	Type     string                 `yaml:"type"`
	Priority int                    `yaml:"priority"`
	Config   map[string]interface{} `yaml:"config"`
}
