package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Application struct {
		Port int `yaml:"port"`
	} `yaml:"application"`
	Logger struct {
		Level string `yaml:"level"`
		Path  string `yaml:"path"`
	} `yaml:"logger"`
	Database struct {
		Type string `yaml:"type"`
		File string `yaml:"file"`
	} `yaml:"database"`
}

func LoadDefaultConfig() Config {
	fp, err := ioutil.ReadFile("./config.yaml")
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}
	cf := Config{}
	if os.IsNotExist(err) {
		err := ioutil.WriteFile("config.yaml", defaultConfigSource, 0644)
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(defaultConfigSource, &cf)
		if err != nil {
			panic(err)
		}
		return cf
	}
	err = yaml.Unmarshal(fp, &cf)
	if err != nil {
		panic(err)
	}

	return cf
}
