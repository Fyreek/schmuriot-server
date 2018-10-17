package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// Config is a global variables that stores the loaded configuration
var Config config

type config struct {
	Server serverConfig
	Game   gameConfig
}

type serverConfig struct {
	IP   string
	Port int
	// TickRate int
	Slots int
	CORS  bool
	Debug bool
}

type gameConfig struct {
	NameLength   int
	SlotsPerRoom int
}

// ReadConfig reads the configuration from file and stores it in the global variable "Config"
func ReadConfig(configfile string) error {
	_, err := os.Open(configfile)
	if err != nil {
		return err
	}
	var config config
	_, err = toml.DecodeFile(configfile, &config)
	if err != nil {
		log.Fatal(err)
		return err
	}
	Config = config
	return err
}
