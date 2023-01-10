package config

import (
	"encoding"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"
)

type envMap map[string]*string

func (m envMap) apply() (backupEnv envMap) {
	var err error
	backupEnv = envMap{}
	for key, value := range m {
		val, valSet := os.LookupEnv(key)
		if valSet {
			backupEnv[key] = &val
		} else {
			backupEnv[key] = nil
		}
		if value != nil {
			log.Printf("Setting %s=%s\n", key, *value)
			err = os.Setenv(key, *value)
			if err != nil {
				panic(err)
			}
		} else {
			log.Printf("Clearing %s\n", key)
			err = os.Unsetenv(key)
			if err != nil {
				panic(err)
			}
		}
	}
	return backupEnv
}

func (m envMap) clear() {
	for k := range m {
		delete(m, k)
	}
}

type envKey int

const (
	_ envKey = iota
	envHost
	envPort
	envShutdownTimeout
	envTLSEnabled
	envTLSCertFile
	envTLSKeyFile
	envSecret
)

var (
	envKeyStringMap = map[envKey]string{
		envHost:            "host",
		envPort:            "port",
		envShutdownTimeout: "shutdown_timeout",
		envTLSEnabled:      "tls_enabled",
		envTLSCertFile:     "tls_cert_file",
		envTLSKeyFile:      "tls_key_file",
		envSecret:          "secret",
	}
)

func (k envKey) Key() (result string) {
	if result, ok := envKeyStringMap[k]; ok {
		return fmt.Sprintf("%s_%s", EnvPrefixU, strings.ToUpper(result))
	}
	panic(fmt.Sprintf("invalid location type %v", k))
}

type envKeyMap map[envKey]interface{}

func (m envKeyMap) toEnvMap() (result envMap, err error) {
	result = envMap{}
	for key, value := range m {
		if value == nil {
			result[key.Key()] = nil
			continue
		}
		var stringValue string
		switch valueTyped := value.(type) {
		case string:
			stringValue = valueTyped
		case bool:
			if valueTyped {
				stringValue = "true"
			} else {
				stringValue = "false"
			}
		case int:
			stringValue = strconv.Itoa(valueTyped)
		case time.Duration:
			stringValue = valueTyped.String()
		case encoding.TextMarshaler:
			textValue, err := valueTyped.MarshalText()
			if err != nil {
				return nil, err
			}
			stringValue = string(textValue)
		default:
			return nil, fmt.Errorf("unsupported type %T", value)
		}
		result[key.Key()] = &stringValue
	}
	return result, nil
}

func (m envKeyMap) apply() (backupEnv envMap, err error) {
	envMap, err := m.toEnvMap()
	if err != nil {
		return nil, fmt.Errorf("failed to convert to env map: %w", err)
	}
	backupEnv = envMap.apply()
	return backupEnv, nil
}

func (m envKeyMap) update(updates envKeyMap) {
	for key, value := range updates {
		m[key] = value
	}
}

func assertError(t *testing.T, err error, regex string) {
	assert.Error(t, err)
	assert.Regexp(t, regex, err.Error())
}

func TestDefaultConfig(t *testing.T) {
	var err error
	config := ServerConfig{}
	config.SetDefaults()
	err = config.Validate()
	assert.NoError(t, err)
}

func TestLoadConfigFromEnv(t *testing.T) {
	var err error
	okayEnv := envKeyMap{
		envHost:            "1.2.3.4",
		envPort:            1234,
		envShutdownTimeout: time.Duration(100) * time.Second,
		envSecret:          "Secret",

		envTLSEnabled:  true,
		envTLSCertFile: "/etc/cert.pem",
		envTLSKeyFile:  "/etc/key.pem",
	}

	backupEnv, err := okayEnv.apply()
	if err != nil {
		t.Fatal(err)
	}
	dto := &ServerConfig{}
	err = Load(dto, nil)
	assert.NoError(t, err)
	_ = backupEnv.apply()

	type TestCase struct {
		name       string
		envUpdate  envKeyMap
		altEnv     envKeyMap
		errorRegex string
		validate   func(t *testing.T, testCase *TestCase, env envKeyMap, config *ServerConfig, err error)
	}
	type validateFunc func(t *testing.T, testCase *TestCase, env envKeyMap, config *ServerConfig, err error)

	var checkTest validateFunc = func(t *testing.T, testCase *TestCase, env envKeyMap, config *ServerConfig, err error) {
		if len(testCase.errorRegex) > 0 {
			// assert.Nil(t, config)
			assertError(t, err, testCase.errorRegex)
		} else {
			if err != nil {
				t.Fatal(err)
			}
			// assert.NotNil(t, config)
			assert.Equal(t, env[envHost], config.Host)
			assert.Equal(t, env[envPort], config.Port)
			assert.Equal(t, env[envShutdownTimeout], config.ShutdownTimeout)
			assert.Equal(t, env[envSecret], config.Secret.UnmaskString())

			assert.Equal(t, env[envTLSEnabled], config.TLS.Enabled)
			assert.Equal(t, env[envTLSCertFile], config.TLS.CertFile)
			assert.Equal(t, env[envTLSKeyFile], config.TLS.KeyFile)
			err = config.Validate()
			assert.NoError(t, err)
		}
	}

	testCases := []TestCase{
		{
			name: "everything okay",
		},
		{
			name:      "hostnames",
			envUpdate: envKeyMap{envHost: "main.example.com"},
		},
		{
			name:       "invalid host",
			envUpdate:  envKeyMap{envHost: "#@!!!@@@-_.11"},
			errorRegex: "validation.*Host",
		},
		{
			name:       "invalid port",
			envUpdate:  envKeyMap{envPort: "ABCD"},
			errorRegex: "parse.*Port",
		},
		{
			name:       "invalid port",
			envUpdate:  envKeyMap{envPort: "ABCD"},
			errorRegex: "parse.*Port",
		},
		{
			name:      "censor",
			envUpdate: envKeyMap{envSecret: "SECRET_VALUE"},
		},
		{
			name: "no key file",
			envUpdate: envKeyMap{
				envTLSEnabled:  true,
				envTLSCertFile: nil,
			},
			errorRegex: "TLS.CertFile",
		},
		{
			name: "no cert file",
			envUpdate: envKeyMap{
				envTLSEnabled: true,
				envTLSKeyFile: nil,
			},
			errorRegex: "TLS.KeyFile",
		},
		{
			name: "tls disabled",
			envUpdate: envKeyMap{
				envTLSEnabled:  false,
				envTLSKeyFile:  "",
				envTLSCertFile: "",
			},
		},
	}
	for _, testCase := range testCases {
		testCase := testCase
		t.Run(fmt.Sprintf("%v", testCase.name), func(t *testing.T) {
			testEnv := envKeyMap{}
			var err error
			if testCase.altEnv != nil {
				err = copier.Copy(&testEnv, &testCase.altEnv)
			} else {
				err = copier.Copy(&testEnv, &okayEnv)
			}
			if err != nil {
				t.Fatal(err)
			}
			if testCase.envUpdate != nil {
				t.Logf("Updating %v with %v", testEnv, testCase.envUpdate)
				testEnv.update(testCase.envUpdate)
			}
			t.Logf("testEnv = %v", testEnv)
			backupEnv, err := testEnv.apply()
			if err != nil {
				t.Fatal(err)
			}
			defer backupEnv.apply()
			config := &ServerConfig{}
			err = Load(config, nil)
			backupEnv.apply()
			backupEnv.clear()
			if testCase.validate != nil {
				testCase.validate(t, &testCase, testEnv, config, err)
			} else {
				checkTest(t, &testCase, testEnv, config, err)
			}
		})
	}
}
