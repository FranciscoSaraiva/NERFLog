package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"nerflog/errors"
)

type ApiConfig struct {
	BaseUrl string `json:"base_url"`
	Realm   string `json:"realm"`
	Region  string `json:"region"`
	ApiKey  string `json:"api_key"`
}

var (
	Config *ApiConfig
)

// Reads the json configuration file and setups the config variable to access the fields.
func ReadConfig() error {
	fmt.Println("Reading configuration file...")

	file, err := ioutil.ReadFile("./config/config.json")
	errors.CheckCommonError(err)

	err = json.Unmarshal(file, &Config)
	errors.CheckCommonError(err)

	log.Printf("Base URL: %v\n", Config.BaseUrl)
	log.Printf("Realm: %v\n", Config.Realm)
	log.Printf("Region: %v\n", Config.Region)
	log.Printf("API Key: %v\n", Config.ApiKey)

	return nil
}
