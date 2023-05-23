package config

import (
	"log"
	"os"
	"testing"

	"gopkg.in/yaml.v3"
)

func setup() {
	var config = Config{}
	config.Db.Database = "database"
	config.Db.Host = "localhost"
	config.Db.Port = 5432
	config.Db.Password = "pass"
	config.Db.Username = "user"

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

func TestApp(t *testing.T) {
	t.Run("ReadConfig loads the config", func(t *testing.T) {
		setup()
		defer tearDown()

		app := GetConfig()
		app.ReadConfig("config.yml")

		if config.Db.Database != "database" {
			t.Fatalf("Db config not loaded properly:%s", config.Db.Database)
		}

		if config.Db.Port != 5432 {
			t.Fatalf("Port config not loaded properly:%d", config.Db.Port)
		}
	})

	t.Run("AssembleConnectString creates the string from the config", func(t *testing.T) {
		setup()
		defer tearDown()

		config := GetConfig()
		err := config.ReadConfig("config.yml")
		if err != nil {
			t.Fatalf("error reading config file:%s", err)
		}
		connString := config.AssembleConnectString()
		if connString != "postgres://user:pass@localhost/database?sslmode=disable" {
			t.Fatalf("connect string did not come out correctly: %s", connString)
		}
	})

}
