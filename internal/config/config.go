package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() Config {
	file_path := filepath.Clean(getHomeDir() + "/.gatorconfig.json")
	file, err := os.Open(file_path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}
	return config
}

func (config *Config) SetUser(name string) error {
	config.CurrentUserName = name
	file_path := filepath.Clean(getHomeDir() + "/.gatorconfig.json")
	file, err := os.Create(file_path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(config)
	return err
}

func getHomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return home
}
