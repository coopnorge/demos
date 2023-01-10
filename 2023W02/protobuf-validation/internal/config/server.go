package config

import (
	"fmt"
	"time"

	"github.com/coopnorge/go-masker-lib"
	"github.com/spf13/viper"
)

// TLSConfig captures TLS configuration
type TLSConfig struct {
	Enabled  bool
	KeyFile  string `validate:"required,min=1" mapstructure:"key_file" yaml:"key_file"`
	CertFile string `validate:"required,min=1" mapstructure:"cert_file" yaml:"cert_file"`
}

func (c *TLSConfig) setDefaults() {
	c.Enabled = false
}

// ServerConfig contains configuration for the service
type ServerConfig struct {
	Base
	// The host to listen on, e.g. "1.2.3.4".
	Host string `validate:"hostname_rfc1123"`
	// The port to listen on, e.g. 1234.
	Port int `validate:"gt=1024,lt=65535"`
	// The timeout for graceful shutdown.
	ShutdownTimeout time.Duration `mapstructure:"shutdown_timeout" yaml:"shutdown_timeout"`

	TLS TLSConfig

	Secret masker.CensoredString
}

// Address returns the Address of the config as "{{ .Host }}:{{ .Port }".
func (c *ServerConfig) Address() string {
	return fmt.Sprintf("%s:%v", c.Host, c.Port)
}

func (c *ServerConfig) SetDefaults() {
	c.Host = "127.0.0.1"
	c.Port = 32604
	c.ShutdownTimeout = time.Duration(5) * time.Second
	c.TLS.setDefaults()
}

func (c *ServerConfig) Validate() (err error) {
	except := []string{}
	if !c.TLS.Enabled {
		except = append(except, "TLS.KeyFile")
		except = append(except, "TLS.CertFile")
	}
	err = validate.StructExcept(c, except...)
	if err != nil {
		return err
	}
	return nil
}

func (c *ServerConfig) LoadOptions() *LoadOptions {
	return &LoadOptions{SkipValidation: false, EnvPrefix: EnvPrefix, Viper: viper.GetViper()}
}
