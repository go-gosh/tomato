package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"gopkg.in/yaml.v3"
)

// RootPathMark this char represent executable dir
const RootPathMark = "$(root)"

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
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		appRoot := filepath.Dir(ex)
		configPath := filepath.Join(appRoot, "config.yaml")
		fp, err := ioutil.ReadFile(configPath)
		if err != nil && !os.IsNotExist(err) {
			panic(err)
		}
		if os.IsNotExist(err) {
			err := ioutil.WriteFile(configPath, defaultConfigSource, 0644)
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
		_defaultConfig.Database.File = strings.ReplaceAll(_defaultConfig.Database.File, RootPathMark, appRoot)
	})

	return _defaultConfig
}
