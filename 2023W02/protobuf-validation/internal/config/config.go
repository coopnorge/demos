package config

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	// EnvPrefix is the environment prefix for the config.
	EnvPrefix = "gosvc"
	// EnvPrefixU is the upper case environment prefix for the config.
	EnvPrefixU                     = strings.ToUpper(EnvPrefix)
	validate   *validator.Validate = validator.New()
)

// Loadable is a loadable config object.
type Loadable interface {
	LoadOptions() *LoadOptions
	Validate() (err error)
	SetDefaults()
	Process() (err error)
}

// Base is a base config object.
type Base struct{}

func (c *Base) Process() (err error) {
	return nil
}
