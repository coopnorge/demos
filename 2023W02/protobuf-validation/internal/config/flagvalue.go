package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// FlagValue exposes pflag.Flag as a viper.FlagValue
type FlagValue struct {
	pflag.Flag
}

// Name returns the name of flag.
func (v *FlagValue) Name() string {
	return v.Flag.Name
}

// HasChanged returns true if the flag was changed from it's default/
func (v *FlagValue) HasChanged() bool {
	return v.Flag.Changed
}

// ValueString returns the value of the flag as a string.
func (v *FlagValue) ValueString() string {
	return v.Flag.Value.String()
}

// ValueType returns the type of the flag.
func (v *FlagValue) ValueType() string {
	return v.Flag.Value.Type()
}

var (
	_ = (*viper.FlagValue)(nil)
)
