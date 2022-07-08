package config

import "os"

type Config struct {
	DOMAIN        string
	HOST          string
	PORT          string
	DB_USERNAME   string
	DB_PASSWORD   string
	DB_PORT       string
	DB_HOST       string
	DB_NAME       string
	JWT_SECRET    string
	MAIL_AT       string
	MAIL_RT       string
	MAIL_CLIENT   string
	MAIL_SECRET   string
	MAIL_REDIRECT string
	STATIC_ROOT   string
	STATIC_PATH   string
}

func InitializeConfig() Config {
	return Config{
		HOST:          os.Getenv("HOST"),
		PORT:          os.Getenv("PORT"),
		DB_USERNAME:   os.Getenv("DB_USERNAME"),
		DB_PASSWORD:   os.Getenv("DB_PASSWORD"),
		DB_PORT:       os.Getenv("DB_PORT"),
		DB_HOST:       os.Getenv("DB_HOST"),
		DB_NAME:       os.Getenv("DB_NAME"),
		JWT_SECRET:    os.Getenv("JWT_SECRET"),
		MAIL_AT:       os.Getenv("MAIL_AT"),
		MAIL_RT:       os.Getenv("MAIL_RT"),
		MAIL_CLIENT:   os.Getenv("MAIL_CLIENT"),
		MAIL_SECRET:   os.Getenv("MAIL_SECRET"),
		MAIL_REDIRECT: os.Getenv("MAIL_REDIRECT"),
		STATIC_ROOT:   os.Getenv("STATIC_ROOT"),
		STATIC_PATH:   os.Getenv("STATIC_PATH"),
	}
}
