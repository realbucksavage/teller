package teller

import "io"

type Configuration interface {
	Validate() error
}

type ConfigSource interface {
	Name() string
	Refresh() error
	RefreshRate() int
	Priority() int
	Load(application, profile, label string) (map[string]interface{}, error)

	io.Closer
}
