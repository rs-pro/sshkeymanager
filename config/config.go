package config

import (
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

type ConfigData struct {
	GinMode string `yaml:"gin_mode"`
	Listen  string `yaml:"listen"`
	ApiKey  string `yaml:"api_key"`
	Log     bool   `yaml:"log"`
}

var Config ConfigData

func init() {
	var err error

	dat, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal([]byte(dat), &Config)
	if err != nil {
		panic(err)
	}
}
