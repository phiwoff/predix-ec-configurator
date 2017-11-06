package helpers

import (
	"encoding/json"
	"log"
	"os"

	"github.com/pkg/errors"
)

// LoadConfig ...
func LoadConfig(filename string, config *UserConfig) error {
	var err error
	if filename != "" {
		log.Println("-> Loading configuration file...")
		err = loadFile(filename, config)
		log.Println("-> DONE!")
	}
	return errors.Wrap(err, "filename cannot be an empty string")
}

func loadFile(filename string, config *UserConfig) error {
	configFile, err := os.Open(filename)
	if err != nil {
		return errors.Wrap(err, "failed to read config file")
	}
	defer configFile.Close()
	decoder := json.NewDecoder(configFile)

	if err = decoder.Decode(&config); err != nil {
		return errors.Wrap(err, "failed to decode config file")
	}
	return nil
}

// UserConfig ...
type UserConfig struct {
	Predix struct {
		Domain   string `json:"domain"`
		API      string `json:"api"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"predix"`
	WebServer struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"webserver"`
}
