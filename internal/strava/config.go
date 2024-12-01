package strava

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Config struct {
	Email    string
	Password string
}

func (c *Config) Validate() error {
	if c.Email == "" {
		return fmt.Errorf("Email is required")
	}
	if c.Password == "" {
		return fmt.Errorf("Password is required")
	}

	return nil
}

func NewConfigFromFile(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Could not open config file '%s': %w", path, err)
	}
	defer file.Close()
	body, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("Could not read from config file '%s': %w", path, err)
	}
	var config Config
	err = json.Unmarshal(body, &config)
	if err != nil {
		return nil, fmt.Errorf("Could not unmarshal json from config file '%s': %w", path, err)
	}

	err = config.Validate()
	if err != nil {
		return nil, fmt.Errorf("Invalid config: %w", err)
	}

	return &config, nil
}

func NewConfig(email, password string) (*Config, error) {
	config := &Config{
		Email:    email,
		Password: password,
	}

	err := config.Validate()
	if err != nil {
		return nil, fmt.Errorf("Invalid config: %w", err)
	}

	return config, nil
}
