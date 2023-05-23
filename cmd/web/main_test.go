package main

import (
	"log"
	"os"
	"rentals-api/pkg/config"
	"testing"

	"gopkg.in/yaml.v3"
)

func setup() {
	var config = config.Config{}
	config.Db.Type = "sqlite3"

	configFile, err := os.Create("config.yml")
	if err != nil {
		log.Fatalf("could not mock config %s", err)
	}
	defer configFile.Close()

	configBytes, err := yaml.Marshal(config)
	if err != nil {
		log.Fatalf("failed to marshal dummy config to bytes:%s", err)
	}

	err = os.WriteFile("config.yml", configBytes, 0644)
	if err != nil {
		log.Fatalf("failed to write config to disk:%s", err)

	}
}

func tearDown() {
	os.Remove("config.yml")
}

func TestMain(t *testing.T) {
	t.Run("Test that setupDB connects", func(t *testing.T) {
		setup()
		defer tearDown()
		_, err := setupDB("config.yml")

		if err != nil {
			t.Errorf("failed to setup DB: %s", err)
		}

	})
}
