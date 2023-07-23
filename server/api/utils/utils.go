package utils

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"github.com/syamsv/apollo/api/constants"
)

func ImportEnv() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetDefault("MIGRATE", false)
	viper.SetDefault("ENVIRONMENT", "development")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Panicln(fmt.Errorf("fatal error config file: %s", err))
		}
	}
	for _, element := range constants.ENV {
		if viper.GetString(element) == "" {
			log.Panicln(fmt.Errorf("env variables not present %s", element))
		}
	}
}
