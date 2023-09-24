package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var (
	config *viper.Viper
)

const (
	USERID        = "userId"
	AUTHORIZATION = "Authorization"
	XLENGTH       = "X-Length"
)

func Load(env string, configPaths ...string) {
	//set yaml file path
	var err error
	config = viper.New()
	config.SetConfigFile("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")
	config.AddConfigPath("../app/")
	config.AddConfigPath(".")
	if len(configPaths) != 0 {
		for _, path := range configPaths {
			config.AddConfigPath(path)
		}
	}
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file", err)
		return
	}
	if env == "server" {
		for _, value := range config.AllKeys() {
			if value == "version" {
				continue //skip the version

			}
			keys := config.GetString(value)
			keys = strings.ReplaceAll(keys, "$", "")
			if ev, ok := os.LookupEnv(keys); ok {
				config.Set(value, ev)
			} else {
				log.Fatal("Env value for key [", keys, "] is missing")
			}

		}
		log.Println("Application runnng on server ")

	} else {
		log.Println("Applicatin running locally ")
	}
}

func GetConfig() *viper.Viper {
	return config
}
