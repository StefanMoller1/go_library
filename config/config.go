package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Server   Server   `toml:"api_server"`
	Database Database `toml:"database"`
}

func NewConfig(file, name string) (*Config, error) {
	c := Config{}

	if _, err := os.Stat(file); err != nil {
		file = "/etc/" + name + "/config.toml"
	}

	if _, err := toml.DecodeFile(file, &c); err != nil {
		return &c, err
	}

	return &c, nil
}

type Server struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
}

type Database struct {
	DNS string `toml:"dns"`
}
