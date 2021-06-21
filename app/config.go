package app

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
