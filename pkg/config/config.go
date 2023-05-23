package config

import (
	"errors"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Db DBConfig
}

type DBConfig struct {
	Type     string
	Username string
	Password string
	Host     string
	Port     int
	Database string
}

var config *Config

func GetConfig() *Config {
	if config == nil {
		config = &Config{}
	}
	return config
}

func (config *Config) ReadConfig(fileName string) error {
	log.Printf("reading config: %s\n", fileName)
	configFileBytes, err := os.ReadFile(fileName)
	if err != nil {
		return errors.New(fmt.Sprintf("error opening config file: %s\n", err))
	}

	err = yaml.Unmarshal(configFileBytes, config)
	if err != nil {
		return errors.New(fmt.Sprintf("error unmarshalling config file:%s\n", err))
	}
	return nil
}

func (config *Config) AssembleConnectString() string {

	if config.Db.Type == "postgres" {
		connectionString := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
			config.Db.Username,
			config.Db.Password,
			config.Db.Host,
			config.Db.Database)

		return connectionString
	}

	if config.Db.Type == "sqlite3" {
		return ":memory:"
	}

	return ""
}
