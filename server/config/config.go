package config

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/syamsv/apollo/api/constants"
	"github.com/syamsv/apollo/api/utils"
)

var (
	SERVER_PORT            = ""
	MIGRATE                = false
	POSTGRES_USER          = ""
	POSTGRES_PASS          = ""
	POSTGRES_DB            = ""
	POSTGRES_HOST          = ""
	POSTGRES_PORT          = ""
	CORS_ORIGIN            = ""
	REDIS_HOST             = ""
	REDIS_PORT             = ""
	REDIS_PASSWORD         = ""
	JWT_ACCESS_KEY_SECRET  = ""
	JWT_REFRESH_KEY_SECRET = ""
)

func LoadConfig() {
	utils.ImportEnv()
	SERVER_PORT = fmt.Sprintf(":%s", viper.GetString("SERVER_PORT"))
	MIGRATE = viper.GetBool("MIGRATE")
	POSTGRES_USER = viper.GetString("POSTGRES_USER")
	POSTGRES_PASS = viper.GetString("POSTGRES_PASS")
	POSTGRES_DB = viper.GetString("POSTGRES_DB")
	POSTGRES_HOST = viper.GetString("POSTGRES_HOST")
	POSTGRES_PORT = viper.GetString("POSTGRES_PORT")
	if viper.GetString("ENVIRONMENT") == "production" {
		CORS_ORIGIN = constants.PROD_CORS_ALLOWED_ORIGINS
	} else {
		CORS_ORIGIN = constants.DEV_CORS_ALLOWED_ORIGINS
	}
	REDIS_HOST = viper.GetString("REDIS_HOST")
	REDIS_PORT = viper.GetString("REDIS_PORT")
	REDIS_PASSWORD = viper.GetString("REDIS_PASSWORD")
	JWT_ACCESS_KEY_SECRET = viper.GetString("JWT_ACCESS_KEY_SECRET")
	JWT_REFRESH_KEY_SECRET = viper.GetString("JWT_REFRESH_KEY_SECRET")
}
