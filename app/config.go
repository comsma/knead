package app

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DB       string `yaml:"db"`
	} `yaml:"database"`
	Tables []string `yaml:"tables"`
}

func ReadConfig(path string) (*Config, error) {
	var config Config

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file:%v", err)
	}
	defer file.Close()

	dec := yaml.NewDecoder(file)

	if err := dec.Decode(&config); err != nil {
		return nil, fmt.Errorf("could not decode config:%v", err)
	}
	return &config, nil

}
