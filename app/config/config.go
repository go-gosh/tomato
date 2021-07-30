package config

import (
	"io/ioutil"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

var (
	_defaultConfig Config
	_once          sync.Once
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
	Runtime struct {
		DefaultUser string `yaml:"default_user"`
	} `yaml:"runtime"`
}

func LoadDefaultConfig() Config {
	_once.Do(func() {
		fp, err := ioutil.ReadFile("./config.yaml")
		if err != nil && !os.IsNotExist(err) {
			panic(err)
		}
		if os.IsNotExist(err) {
			err := ioutil.WriteFile("config.yaml", defaultConfigSource, 0644)
			if err != nil {
				panic(err)
			}
			err = yaml.Unmarshal(defaultConfigSource, &_defaultConfig)
			if err != nil {
				panic(err)
			}
		}
		err = yaml.Unmarshal(fp, &_defaultConfig)
		if err != nil {
			panic(err)
		}
	})

	return _defaultConfig
}
