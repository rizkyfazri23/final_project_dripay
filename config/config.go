package config

import "github.com/rizkyfazri23/dripay/utils"

type ServerConfig struct {
	ServerPort string
}

type DBConfig struct {
	DBHost, DBPort, DBUser, DBPass, DBName, DBSSLMode string
}

type AppConfig struct {
	ServerConfig
	DBConfig
}

func (c *AppConfig) readEnv() {
	api := utils.DotEnv("SERVER_PORT")

	c.DBConfig = DBConfig{
		DBHost:    utils.DotEnv("DB_HOST"),
		DBPort:    utils.DotEnv("DB_PORT"),
		DBUser:    utils.DotEnv("DB_USER"),
		DBPass:    utils.DotEnv("DB_PASSWORD"),
		DBName:    utils.DotEnv("DB_NAME"),
		DBSSLMode: utils.DotEnv("SSL_MODE"),
	}

	c.ServerConfig = ServerConfig{ServerPort: api}
}

func NewConfig() AppConfig {
	config := AppConfig{}
	config.readEnv()
	return config
}
