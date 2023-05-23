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

	log.Printf("Config Values: %v\n", config)

	return nil
}

func (config *Config) AssembleConnectString() string {

	connectionString := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		config.Db.Username,
		config.Db.Password,
		config.Db.Host,
		config.Db.Database)

	return connectionString
}
