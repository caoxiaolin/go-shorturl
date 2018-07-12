package config

import (
	"github.com/BurntSushi/toml"
	"os"
	"path/filepath"
)

type TomlConfig struct {
	Server   server
	Database database
	Redis    redis
	Log      log
}

type database struct {
	Host     string
	Port     int
	UserName string
	PassWord string
	DbName   string
	MaxConn  int
}

type redis struct {
	Host     string
	Port     int
	Password string
	Database int
}

type server struct {
	Host string
	Port int
}

type log struct {
	Logpath string
	Logfile string
}

var Cfg *TomlConfig

func init() {
	filePath, _ := filepath.Abs("./config/config.toml")
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		filePath, _ = filepath.Abs("../config/config.toml")
	}
	Cfg = new(TomlConfig)
	if _, err := toml.DecodeFile(filePath, &Cfg); err != nil {
		panic(err)
	}
}
