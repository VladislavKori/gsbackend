package config

import (
	"github.com/spf13/viper"
)

type Env struct {
	SERVER_PORT string `mapstructure:"SERVER_PORT"`
}

func NewEnv() (*Env, error) {
	env := Env{}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&env); err != nil {
		return nil, err
	}

	return &env, nil
}
