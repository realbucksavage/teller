package sources

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"

	"github.com/realbucksavage/teller"
)

func Init(cfg Configuration) (*Sources, error) {

	sources := make([]teller.ConfigSource, 0)
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

		source, err := factory.New(sourceCfg.Name, sourceCfg.Priority, fConfig)
		if err != nil {
			return nil, err
		}

		fmt.Printf("%s/%s created", sourceCfg.Type, source.Name())
		sources = append(sources, source)
	}

	sort.Slice(cfg.Sources, func(i, j int) bool { return sources[i].Priority() < sources[j].Priority() })
	return &Sources{sourceList: sources}, nil
}

type Sources struct {
	sourceList []teller.ConfigSource
}

func (src *Sources) Start() error {

	for _, source := range src.sourceList {
		name := source.Name()
		log.Printf("initiating source %q", name)
		if err := source.Refresh(); err != nil {
			return errors.Wrapf(err, "cannot start source %q", name)
		}
	}

	return nil
}

func (src *Sources) StartRefresh(errChan chan<- error) {
	for _, source := range src.sourceList {
		name := source.Name()
		log.Printf("refreshing %q every %d seconds", name, source.RefreshRate())
		go refreshLoop(source, errChan)
	}
}

func (src *Sources) Close() error {
	multiErr := new(multierror.Error)
	for _, source := range src.sourceList {
		if err := source.Close(); err != nil {
			multiErr = multierror.Append(multiErr, errors.Wrapf(err, "cannot close %q", source.Name()))
		}
	}

	return multiErr.ErrorOrNil()
}

func refreshLoop(source teller.ConfigSource, errChan chan<- error) {
	ticker := time.NewTicker(time.Duration(source.RefreshRate()) * time.Second)
	for {
		<-ticker.C
		log.Printf("refreshing %q", source.Name())
		if err := source.Refresh(); err != nil {
			errChan <- err
		}
	}
}
