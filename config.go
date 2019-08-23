package main

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

type Config struct {
	Port       int
	genericCfg map[string]interface{} `toml:"-"`
}

var DefaultConfig = Config{Port: 8080}

func ParseConfig(file string) (*Config, error) {
	cfg := Config{}

	//needed so users can use whatever key they want in the templates
	var genericCfg map[string]interface{}

	bs, err := ioutil.ReadFile(file)

	if err != nil {
		return nil, err
	}

	_, err = toml.Decode(string(bs), &cfg)

	if err != nil {
		return nil, err
	}

	_, err = toml.Decode(string(bs), &genericCfg)

	cfg.genericCfg = genericCfg

	return &cfg, err
}

func (c *Config) Get(name string) interface{} {
	if val, ok := c.genericCfg[name]; ok {
		return val
	}

	return ""
}
