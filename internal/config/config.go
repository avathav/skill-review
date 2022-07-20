package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	Environment string
	Version     string
}

func LoadParameters() (c Config, err error) {
	viper.SetConfigName("parameters")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err = viper.ReadInConfig(); err != nil {
		return c, errors.Wrap(err, "error reading config file")
	}

	if err = viper.Unmarshal(&c); err != nil {
		errors.Wrap(err, "Unable to decode into struct")
	}

	return c, nil
}
