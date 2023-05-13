package git

import (
	"github.com/go-playground/validator/v10"

	"github.com/realbucksavage/teller"
)

type Configuration struct {
	Repository string `mapsquash:"config" validate:"required"`
}

func (config *Configuration) Validate() error {
	return validator.New().Struct(config)
}

func defaultConfiguration() teller.Configuration {
	return new(Configuration)
}
