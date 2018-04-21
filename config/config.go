package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const ConfigFile = "config.json"

var conf = &Config{}

type Config struct {
	Database struct {
		MongoDB bool
		MYSQL   bool
	}
	MongoDB struct {
		Host     string
		Password string
		Port     int
		DBName   string
		User     string
	}
	MYSQL struct {
		Host     string
		Password string
		Port     int
		DBName   string
		User     string
	}
}

func init() {
	fmt.Println(conf.GetPath() + ConfigFile)
	dat, err := ioutil.ReadFile(conf.GetPath() + ConfigFile)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(dat, conf)
	if err != nil {
		panic(err)
	}
}
func Get() *Config {
	return conf
}
func (conf *Config) GetPath() string {
	return os.Getenv("WEATHER_PATH")
}
