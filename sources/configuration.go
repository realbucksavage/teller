package sources

import (
	"fmt"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"

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
	Name   string                 `yaml:"name"`
	Type   string                 `yaml:"type"`
	Config map[string]interface{} `yaml:"config"`
}

func Init(cfg Configuration) (interface{}, error) {

	for _, sourceCfg := range cfg.Sources {

		sType := strings.ToLower(sourceCfg.Type)
		factory, ok := factories[sType]
		if !ok {
			return nil, errors.Errorf("unknown configuration source %q", sType)
		}

		fConfig := factory.DefaultConfig()
		if err := mapstructure.Decode(sourceCfg.Config, fConfig); err != nil {
			return nil, errors.Wrapf(err, "cannot configure %q", sourceCfg.Name)
		}

		if err := fConfig.Validate(); err != nil {
			return nil, errors.Wrapf(err, "invalid configuration in %q", sourceCfg.Name)
		}

		source, err := factory.New(sourceCfg.Name, fConfig)
		if err != nil {
			return nil, err
		}

		fmt.Printf("%s/%s created", sourceCfg.Type, source.Name())
	}

	return nil, nil
}
