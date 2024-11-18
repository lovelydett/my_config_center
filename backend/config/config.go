package config

import (
	"encoding/json"
	"fmt"
	"os"

	utils "wolf/utils"
)

type MySqlConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type PgConfig struct {
	Uri string `json:"uri"`
}

type OSSConfig struct {
	Region           string `json:"region"`
	BucketName       string `json:"bucket"`
	InternalEndpoint string `json:"internal_endpoint"`
	ExternalEndpoint string `json:"external_endpoint"`
}

type RedisConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password" `
	Database int    `json:"database"`
}

type DeployConfig struct {
	MySQL MySqlConfig `json:"mysql"`
	Pg    PgConfig    `json:"pg"`
	OSS   OSSConfig   `json:"oss"`
	Redis RedisConfig `json:"redis"`
}

var deployConfig DeployConfig

func init() {
	// First load ENV variable for the config file path
	config_file_path := os.Getenv("CONFIG")
	if config_file_path == "" {
		config_file_path = "deploy_config.json"
	}

	config_file_path = utils.GetProjectRoot() + "/config/" + config_file_path

	fmt.Println(os.Getwd())

	// Load the json file
	config_file, err := os.ReadFile(config_file_path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(config_file, &deployConfig)
	if err != nil {
		panic(err)
	}
}

func GetDeployConfig() DeployConfig {
	return deployConfig
}
