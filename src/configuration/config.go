package configuration

import (
	"os"
	"encoding/json"
	"fmt"
)

type Config struct {
	Database struct{
		Name 	   string  `json:"dbname"`
		User       string  `json:"user"`
		Password   string  `json:"password"`
	} `json:"database"`

	Key string `json:"key"`
}

func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
