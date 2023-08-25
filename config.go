// Package main provides the main entry point for the Terraform Analytic Tool application.
// It reads the configuration from file or environment variables and starts the API server.
package main

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/FrangipaneTeam/terraform-analytic-tool/api"
	"github.com/FrangipaneTeam/terraform-analytic-tool/clients"
	"gopkg.in/yaml.v3"
)

// config represents the configuration for the application.
type config struct {
	Redis    clients.RedisConfig    `yaml:"redis"`
	InfluxDB clients.InfluxDBConfig `yaml:"influxdb"`
	API      api.Settings           `yaml:"api"`
}

// readConfig reads the configuration from file or environment variables.
// If a path is provided, it reads the configuration from the file at the given path.
// Otherwise, it reads the configuration from environment variables.
// Returns a pointer to the configuration and an error if the configuration is not valid.
func readConfig(path string) (*config, error) {
	c := new(config)

	// Read config from file if path is provided
	if path != "" {
		cfg, err := readConfigFromFile(path)
		if err != nil {
			return nil, err
		}

		c = cfg
	}

	// Read config from environment variables
	c.readConfigFromEnvVars()

	if !c.isValid() {
		return nil, errors.New("config is not valid")
	}

	return c, nil
}

// readConfigFromFile reads the configuration from file.
// Returns a pointer to the configuration and an error if the file does not exist or cannot be read.
func readConfigFromFile(path string) (*config, error) {
	c := new(config)

	// If file exists, read it
	if _, err := os.Stat(path); err == nil {
		// Read file
		dat, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}

		// Unmarshal yaml file
		if err := yaml.Unmarshal(dat, c); err != nil {
			return nil, err
		}

		return c, nil
	}

	return nil, errors.New("config file does not exist")
}

// envVarType represents the type of environment variables.
type envVarType interface {
	string | ~int | bool
}

// SetEnvVars sets the environment variables value if exists and not empty.
// Takes the name of the environment variable and a pointer to the value to set.
func SetEnvVars[T envVarType](envVarName string, value *T) {
	if v, ok := os.LookupEnv(envVarName); ok {
		if v != "" {
			*value = convertType(v, *value).(T)
		}
	}
}

// convertType converts the value from string to the specified type.
// Takes a string value and the type to convert to.
// Returns the converted value as an interface{}.
func convertType(from string, to any) interface{} {
	switch to.(type) {
	case string:
		return from
	case int:
		x, _ := strconv.Atoi(from)
		return x
	case bool:
		return from == "true"
	default:
		panic("unknown type")
	}
}

// readConfigFromEnvVars reads the configurations from environment variables.
// Sets the values of the configuration fields from the corresponding environment variables.
func (c *config) readConfigFromEnvVars() {
	// REDIS
	SetEnvVars("REDIS_ADDRESS", &c.Redis.Address)
	SetEnvVars("REDIS_PASSWORD", &c.Redis.Password)
	SetEnvVars("REDIS_MAX_RETRIES", &c.Redis.MaxRetries)
	SetEnvVars("REDIS_DIAl_TIMEOUT", &c.Redis.DialTimeout)
	SetEnvVars("REDIS_READ_TIMEOUT", &c.Redis.ReadTimeout)
	SetEnvVars("REDIS_WRITE_TIMEOUT", &c.Redis.WriteTimeout)

	// INFLUXDB
	SetEnvVars("INFLUXDB_ADDRESS", &c.InfluxDB.Address)
	SetEnvVars("INFLUXDB_TOKEN", &c.InfluxDB.Token)
	SetEnvVars("INFLUXDB_ORG", &c.InfluxDB.Org)
	SetEnvVars("INFLUXDB_BUCKET", &c.InfluxDB.Bucket)

	// API
	SetEnvVars("API_ADDRESS", &c.API.Address)
	SetEnvVars("API_PORT", &c.API.Port)
	SetEnvVars("API_TOKEN", &c.API.Token)
}

// isValid checks if the minimal config is valid.
// Returns true if the configuration is valid, false otherwise.
func (c *config) isValid() bool {
	return c.Redis.Address != "" &&
		c.InfluxDB.Address != "" &&
		(strings.HasPrefix(c.InfluxDB.Address, "http://") || strings.HasPrefix(c.InfluxDB.Address, "https://")) &&
		c.InfluxDB.Token != "" &&
		c.InfluxDB.Org != "" &&
		c.InfluxDB.Bucket != ""
}
