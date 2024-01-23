package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port int
	DB   DB
}

type DB struct {
	Host         string
	Port         int
	Username     string
	Password     string
	DatabaseName string
}

func GetConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return Config{
		Port: viper.GetInt("port"),
		DB: DB{
			Host:         viper.GetString("db.host"),
			Port:         viper.GetInt("db.port"),
			Username:     viper.GetString("db.username"),
			Password:     viper.GetString("db.password"),
			DatabaseName: viper.GetString("db.database_name"),
		},
	}
}
