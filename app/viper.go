package app

import (
	"log"
	"github.com/spf13/viper"
)
type Config struct {
	AppName string
	AppPort int
	DBUser string
	DBPassword string
	DBHost string
	DBPort int
	DBName string
}

func NewViper() Config  {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.AddConfigPath("config/")
	viper.AutomaticEnv()

	
	if err := viper.ReadInConfig(); err != nil{
		log.Fatalf("Fatal error config file: %s \n", err)
	}

	config := Config{
		AppName: viper.GetString("app.name"),
		AppPort: viper.GetInt("app.port"),
		DBUser: viper.GetString("database.user"),
		DBPassword: viper.GetString("database.password"),
		DBHost: viper.GetString("database.host"),
		DBPort: viper.GetInt("database.port"),
		DBName: viper.GetString("database.name"),
	}
	return config
}