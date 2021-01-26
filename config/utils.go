package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func GetEnv(KEY string) string {

	viper.SetConfigName("config")

	viper.AddConfigPath(".")

	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {

		logrus.Error("Error reading config file, %s", err)

	}

	ENV := viper.GetString(KEY)

	return ENV
}