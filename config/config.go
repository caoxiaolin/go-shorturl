package config

import (
	"github.com/BurntSushi/toml"
	"path/filepath"
)

type TomlConfig struct {
	Server   server
	Database database
}

type database struct {
	Host     string
	Port     int
	UserName string
	PassWord string
	DbName   string
	MaxConn  int
}

type server struct {
	Host string
	Port int
}

func Load() *TomlConfig {
	filePath, err := filepath.Abs("config/config.toml")
	if err != nil {
		panic(err)
	}
	var config = new(TomlConfig)
	if _, err := toml.DecodeFile(filePath, &config); err != nil {
		panic(err)
	}
	return config
}
