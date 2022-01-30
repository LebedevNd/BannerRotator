package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Database DatabaseConf `json:"database"`
	Server   ServerConf   `json:"server"`
}

type DatabaseConf struct {
	Connection string `json:"dbConnection"`
	Host       string `json:"dbHost"`
	Port       int    `json:"dbPort"`
	Database   string `json:"dbDatabase"`
	Username   string `json:"dbUsername"`
	Password   string `json:"dbPassword"`
}

type ServerConf struct {
	Host string `json:"httpHost"`
	Port int    `json:"httpPort"`
}

func NewConfig(configPath string) (Config, error) {
	var config Config

	pwd, err := os.Getwd()
	if err != nil {
		return Config{}, err
	}

	fmt.Println("Reading from config file...")
	file, err := ioutil.ReadFile(pwd + configPath)
	if err != nil {
		return Config{}, err
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
