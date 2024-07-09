package config

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

// Config holds the configuration values
type Config struct {
	RateLimit struct {
		RequestsPerSecond int `yaml:"requests_per_second"`
		Burst             int `yaml:"burst"`
	} `yaml:"rate_limit"`
}

var config Config

// LoadConfig loads configuration from a file
func LoadConfig(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		switch key {
		case "requests_per_second":
			config.RateLimit.RequestsPerSecond, err = strconv.Atoi(value)
			if err != nil {
				return err
			}
		case "burst":
			config.RateLimit.Burst, err = strconv.Atoi(value)
			if err != nil {
				return err
			}
		default:
			return errors.New("unknown configuration key")
		}
	}
	return nil
}

// GetConfig returns the loaded configuration
func GetConfig() Config {
	return config
}
