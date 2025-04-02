package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Config holds application configuration
type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"database"`
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
}

// LoadConfig loads configuration from yaml file and environment variables
func LoadConfig() (*Config, error) {
	// Default configuration
	cfg := &Config{}
	cfg.Database.Host = "localhost"
	cfg.Database.Port = 5432
	cfg.Database.User = "postgres"
	cfg.Database.Password = "postgres"
	cfg.Database.DBName = "demo"
	cfg.Database.SSLMode = "disable"
	cfg.Server.Port = "8080"

	// Try to read from config file
	configFile := "config.yaml"
	if os.Getenv("CONFIG_FILE") != "" {
		configFile = os.Getenv("CONFIG_FILE")
	}

	// Check if config file exists
	if _, err := os.Stat(configFile); err == nil {
		yamlFile, err := ioutil.ReadFile(configFile)
		if err != nil {
			return nil, fmt.Errorf("error reading config file: %v", err)
		}

		err = yaml.Unmarshal(yamlFile, cfg)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling config: %v", err)
		}
	}

	// Override with environment variables if they exist
	if host := os.Getenv("DB_HOST"); host != "" {
		cfg.Database.Host = host
	}
	if port := os.Getenv("DB_PORT"); port != "" {
		// Simple override without error handling for brevity

		fmt.Sscanf(port, "%d", &cfg.Database.Port)
		fmt.Printf("Port: %d\n", cfg.Database.Port)
	}
	if user := os.Getenv("DB_USER"); user != "" {
		cfg.Database.User = user
	}
	if password := os.Getenv("DB_PASSWORD"); password != "" {
		cfg.Database.Password = password
	}
	if dbname := os.Getenv("DB_NAME"); dbname != "" {
		cfg.Database.DBName = dbname
	}
	if sslmode := os.Getenv("DB_SSLMODE"); sslmode != "" {
		cfg.Database.SSLMode = sslmode
	}
	if port := os.Getenv("SERVER_PORT"); port != "" {
		cfg.Server.Port = port
	}

	return cfg, nil
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host, c.Database.Port, c.Database.User, c.Database.Password, c.Database.DBName, c.Database.SSLMode)
}
