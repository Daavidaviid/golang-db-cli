package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type PostgresQLConfig struct {
	User     string
	Address  string
	Password string
	Port     string
	Database string
}

func (config PostgresQLConfig) Save() error {
	file, _ := json.MarshalIndent(config, "", " ")

	err := ioutil.WriteFile(".skyblogrc", file, 0644)

	return err
}

func (config *PostgresQLConfig) Load(file string) error {

	configFile, err := os.Open(file)
	defer configFile.Close()

	if err != nil {
		return err
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(config)

	return nil
}
