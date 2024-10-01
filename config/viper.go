package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Endpoint     string `mapstructure:"ENDPOIND_SNOW"`
	SnowUsername string `mapstructure:"SNOW_USERNAME"`
	SnowPassword string `mapstructure:"SNOW_PASSWORD"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	if err = viper.ReadInConfig(); err != nil{
		return
	}
	
	if err = viper.Unmarshal(&config); err != nil{
		return
	}
	return config, nil
}