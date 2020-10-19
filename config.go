package obsws

import (
	"flag"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Addr           string `yaml:"addr"`
	Password       string `yaml:"password"`
	ReceiveTimeout string `yaml:"receiveTimeout"`
}

func NewConfig() (*Config, error) {
	path := "config/config.yaml"
	flag.StringVar(&path, "config file path", "config/config.yaml", "set config file path")
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	c := &Config{}
	yaml.Unmarshal(b, c)

	return c, nil
}
