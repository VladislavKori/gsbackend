package config

import (
	"github.com/spf13/viper"
)

type Env struct {
	SERVER_PORT         string `mapstructure:"SERVER_PORT"`
	POSTGRESQL_USERNAME string `mapstructure:"POSTGRESQL_USERNAME"`
	POSTGRESQL_PASSWORD string `mapstructure:"POSTGRESQL_PASSWORD"`
	POISTGRESQL_DB_NAME string `mapstructure:"POISTGRESQL_DB_NAME"`
	POSTGRESQL_PORT     string `mapstructure:"POSTGRESQL_PORT"`
	POSTGRESQL_HOST     string `mapstructure:"POSTGRESQL_HOST"`
	POSTGRESQL_SLLMODE  string `mapstructure:"POSTGRESQL_SLLMODE"`
	JWT_SECRET          string `mapstructure:"JWT_SECRET"`
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
