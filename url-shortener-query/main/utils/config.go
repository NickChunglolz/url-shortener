package utils

import (
	"os"

	"gopkg.in/yaml.v3"
)

const filename = "application.yml"

type Config struct {
	Upstream struct {
		Server struct {
			Host string `yaml:"host"`
			Port string    `yaml:"port"`
		} `yaml:"server"`
	} `yaml:"upstream"`
	Downstream struct {
		Db struct {
			Host string `yaml:"host"`
			Port string    `yaml:"port"`
			Database string    `yaml:"database"`
			User string    `yaml:"user"`
			Password string `yaml:"password"`
		} `yaml:"db"`
		CacheDB struct {
			Host string `yaml:"host"`
			Port string    `yaml:"port"`
		} `yaml:"cache_db"`
	} `yaml:"downstream"`
	Server struct {
		Host string `yaml:"host"`
		Port string    `yaml:"port"`
	} `yaml:"server"`
	Cache struct {
		KeyPrefix string `yaml:"key_prefix"`
		Ttl int    `yaml:"ttl"`
	} `yaml:"cache"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load() (*Config, error) {
    config := &Config{}
    
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    
    decoder := yaml.NewDecoder(file)
    err = decoder.Decode(config)
    if err != nil {
        return nil, err
    }
    
    return config, nil
}