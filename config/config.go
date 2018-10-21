package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// Config is a global variables that stores the loaded configuration
var Config ConfigStruct

type ConfigStruct struct {
	Server serverConfig `json:"-"`
	Player playerConfig `json:"player"`
	Room   roomConfig   `json:"room"`
}

type serverConfig struct {
	IP   string
	Port int
	// TickRate int
	Slots int
	CORS  bool
	Debug bool
}

type playerConfig struct {
	MinNameLength int `json:"minNameLength"`
	MaxNameLength int `json:"maxNameLength"`
}

type roomConfig struct {
	MinNameLength int `json:"minNameLength"`
	MaxNameLength int `json:"maxNameLength"`
	MinSlots      int `json:"minSlots"`
	MaxSlots      int `json:"maxSlots"`
	MinPassLength int `json:"minPassLength"`
	MaxPassLength int `json:"maxPassLength"`
}

// ReadConfig reads the configuration from file and stores it in the global variable "Config"
func ReadConfig(configfile string) error {
	_, err := os.Open(configfile)
	if err != nil {
		return err
	}
	var config ConfigStruct
	_, err = toml.DecodeFile(configfile, &config)
	if err != nil {
		log.Fatal(err)
		return err
	}
	Config = config
	return err
}
