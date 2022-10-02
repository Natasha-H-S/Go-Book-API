package config

import "os"

type Config struct {
	User     string
	Password string
	Host     string
	Database string
}

func NewConfig() *Config {
	return &Config{
		User:     os.Getenv("MYSQL_ROOT_USER"),
		Password: os.Getenv("MYSQL_ROOT_PASSWORD"),
		Host:     os.Getenv("MYSQL_HOST"),
		Database: os.Getenv("MYSQL_DATABASE"),
	}
}
