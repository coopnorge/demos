package config

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

/////////////////////////////////////////////////////////////////////////////////
// LoadOptions
/////////////////////////////////////////////////////////////////////////////////

// LoadOptions provide options which determine how loading should be performed.
type LoadOptions struct {
	SkipValidation bool
	EnvPrefix      string
	Viper          *viper.Viper
}

// WithEnvPrefix sets the environment prefix.
func (o LoadOptions) WithEnvPrefix(value string) *LoadOptions {
	o.EnvPrefix = value
	return &o
}

// LoadConfigInto populates the values of a config object.
func LoadConfigInto(config Loadable, opts *LoadOptions) error {
	vi := opts.Viper
	var err error
	// This is to deal with some quirks of viper
	// see https://github.com/spf13/viper/issues/188#issuecomment-413368673
	cYaml, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshall default config: %w", err)
	}
	cYamlBytes := bytes.NewReader(cYaml)
	vi.SetConfigType("yaml")
	if err := vi.MergeConfig(cYamlBytes); err != nil {
		return fmt.Errorf("failed to merge config: %w", err)
	}
	vi.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	vi.SetEnvPrefix(opts.EnvPrefix)

	vi.AutomaticEnv()

	err = vi.Unmarshal(&config,
		func(config *mapstructure.DecoderConfig) {
			config.WeaklyTypedInput = true
			config.DecodeHook = mapstructure.ComposeDecodeHookFunc(
				config.DecodeHook,
				mapstructure.TextUnmarshallerHookFunc(),
			)
		},
	)
	if err != nil {
		return fmt.Errorf("failed to unmarshall config: %w", err)
	}
	logrus.Debugf("config = %#v", config)
	if !opts.SkipValidation {
		err = config.Validate()
		if err != nil {
			return fmt.Errorf("failed to validate config: %w", err)
		}
	}
	err = config.Process()
	return err
}

// Load loads data into a config object with the supplied load options.
func Load(c Loadable, opts *LoadOptions) (err error) {
	if opts == nil {
		opts = c.LoadOptions()
	}
	c.SetDefaults()
	err = LoadConfigInto(c, opts)
	return err
}
