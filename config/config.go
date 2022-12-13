package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type env struct {
	PORT              string
	DBUSER            string
	MSSQL_SA_PASSWORD string
	DBHOST            string
	MSSQL_TCP_PORT    string
	DBNAME            string
	JWT_SECRET        string
	JWT_EXPIRED       string
	USERINFO_ENDPOINT string
	REALMS            string
	CLIENT_ID         string
	CLIENT_SECRET     string
	CALLBACK_URL      string
	AUTH_URL          string
	LOGOUT_ENDPOINT   string
}

func InitConfig() {
	config := env{
		PORT:              os.Getenv("PORT"),
		DBUSER:            os.Getenv("DBUSER"),
		MSSQL_SA_PASSWORD: os.Getenv("MSSQL_SA_PASSWORD"),
		DBHOST:            os.Getenv("DBHOST"),
		MSSQL_TCP_PORT:    os.Getenv("MSSQL_TCP_PORT"),
		DBNAME:            os.Getenv("DBNAME"),
		JWT_EXPIRED:       os.Getenv("JWT_EXPIRED"),
		JWT_SECRET:        os.Getenv("JWT_SECRET"),
		USERINFO_ENDPOINT: os.Getenv("USERINFO_ENDPOINT"),
		REALMS:            os.Getenv("REALMS"),
		CLIENT_ID:         os.Getenv("CLIENT_ID"),
		CLIENT_SECRET:     os.Getenv("CLIENT_SECRET"),
		CALLBACK_URL:      os.Getenv("CALLBACK_URL"),
		AUTH_URL:          os.Getenv("AUTH_URL"),
		LOGOUT_ENDPOINT:   os.Getenv("LOGOUT_ENDPOINT"),
	}

	viper.SetDefault("app.port", config.PORT)
	viper.SetDefault("mssql1.db", config.DBNAME)
	viper.SetDefault("mssql1.host", fmt.Sprintf("%s%s%s", config.DBHOST, ":", config.MSSQL_TCP_PORT))
	viper.SetDefault("mssql1.user", config.DBUSER)
	viper.SetDefault("mssql1.pass", config.MSSQL_SA_PASSWORD)
	viper.SetDefault("jwt.secret", config.JWT_SECRET)
	viper.SetDefault("jwt.expired", config.JWT_EXPIRED)
	viper.SetDefault("userInfo.endpoint", config.USERINFO_ENDPOINT)
	viper.SetDefault("keycloak.realms", config.REALMS)
	viper.SetDefault("keycloak.clientId", config.CLIENT_ID)
	viper.SetDefault("keycloak.clientSecret", config.CLIENT_SECRET)
	viper.SetDefault("keycloak.callback", config.CALLBACK_URL)
	viper.SetDefault("keycloak.authUrl", config.AUTH_URL)
	viper.SetDefault("keycloak.logout", config.LOGOUT_ENDPOINT)
}
