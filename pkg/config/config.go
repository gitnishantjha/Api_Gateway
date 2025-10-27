package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Port      int        `yaml:"port"`
	Services  []Service  `yaml:"services"`
	EndPoints []EndPoint `yaml:"endpoints"`
}

type Service struct {
	Name    string   `yaml:"name"`
	Servers []string `yaml:"servers"`
}

type EndPoint struct {
	Path        string   `yaml:"path"`
	Service     string   `yaml:"service"`
	Middlewares []string `yaml:"middlewares"`
}

func LoadConfig(filename string) (*Config, error) {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	var c Config
	err = yaml.Unmarshal(bytes, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil

}
