package teller

type configFn func() Configuration
type newSourceFn func(name string, priority int, cfg Configuration) (ConfigSource, error)

type SourceFactory interface {
	New(name string, priority int, cfg Configuration) (ConfigSource, error)
	DefaultConfig() Configuration
}

type defaultSourceFactory struct {
	newFn           newSourceFn
	defaultConfigFn configFn
}

func (factory *defaultSourceFactory) New(name string, priotity int, cfg Configuration) (ConfigSource, error) {
	return factory.newFn(name, priotity, cfg)
}

func (factory *defaultSourceFactory) DefaultConfig() Configuration {
	return factory.defaultConfigFn()
}

func NewFactory(newFactory newSourceFn, defaultConfigFn configFn) SourceFactory {
	return &defaultSourceFactory{
		newFn:           newFactory,
		defaultConfigFn: defaultConfigFn,
	}
}
