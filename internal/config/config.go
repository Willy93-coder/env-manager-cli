// Package config handles loading and parsing the configuration file
// located at ~/.env-manager/config.yaml
package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config holds all configuration loaded from ~/.env-manager/config.yaml
type Config struct {
	Database DatabaseConfig `yaml:"database"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

// DSN returns a PostgreSQL connection string
func (c *Config) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Password,
		c.Database.DBName,
		c.Database.SSLMode,
	)
}

// Validate checks that all required fields are set
func (c *Config) Validate() error {
	if c.Database.Host == "" {
		return fmt.Errorf("database.host is required in config file")
	}
	if c.Database.Port == 0 {
		return fmt.Errorf("database.port is required in config file")
	}
	if c.Database.User == "" {
		return fmt.Errorf("database.user is required in config file")
	}
	if c.Database.Password == "" {
		return fmt.Errorf("database.password is required in config file")
	}
	if c.Database.DBName == "" {
		return fmt.Errorf("database.dbname is required in config file")
	}
	return nil
}

// Load reads the config file from the default location
func Load() (*Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("could not find home directory: %w", err)
	}

	path := fmt.Sprintf("%s/.env-manager/config.yaml", home)

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read config file at %s: %w", path, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("could not parse config file: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}
