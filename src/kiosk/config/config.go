package config

import (
	"../system"
	"strconv"
)

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() *Config {
	dbPort, err := strconv.Atoi(system.GetEnv("DB_PORT", "3306"))
	if err != nil {
		dbPort = 3306
	}

	return &Config{
		DB: &DBConfig{
			Dialect:  system.GetEnv("DB_DIALECT", "msql"),
			Host:     system.GetEnv("DB_HOST", "127.0.0.1"),
			Port:     dbPort,
			Name:     system.GetEnv("DB_NAME", "be_kiosk"),
			Username: system.GetEnv("DB_USERNAME", "root"),
			Password: system.GetEnv("DB_PASSWORD", "123456"),
			Charset:  system.GetEnv("DB_CHARSET", "utf8"),
		},
	}
}
