package config

import (
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	APP_NAME string `env:"APP_NAME,required"`
	APP_KEY  string `env:"APP_KEY,required"`
	ENV      string `env:"ENV,required"`
	PORT     int    `env:"PORT,required"`
	PREFORK  bool   `env:"PREFORK,default=false"`

	LOG_LEVEL      string `env:"LOG_LEVEL,default=debug"`
	LOG_CONSOLE    bool   `env:"LOG_CONSOLE,default=true"`
	LOG_FILE       bool   `env:"LOG_FILE,default=true"`
	LOG_DIR        string `env:"LOG_DIR,default=./storage/log"`
	LOG_MAX_SIZE   int    `env:"LOG_MAX_SIZE,default=50"`
	LOG_MAX_AGE    int    `env:"LOG_MAX_AGE,default=7"`
	LOG_MAX_BACKUP int    `env:"LOG_MAX_BACKUP,default=20"`
	LOG_JSON       bool   `env:"LOG_JSON,default=true"`

	DB_DRIVER   string `env:"DB_DRIVER,default=sqlite"`
	DB_HOST     string `env:"DB_HOST,default=./storage/db.sqlite"`
	DB_PORT     int    `env:"DB_PORT,default=5432"`
	DB_NAME     string `env:"DB_NAME,default=boot"`
	DB_USERNAME string `env:"DB_USERNAME,default=boot"`
	DB_PASSWORD string `env:"DB_PASSWORD,default=boot"`
	DB_SSLMODE  string `env:"DB_SSLMODE,default=disable"`

	SESSION_DRIVER string `env:"SESSION_DRIVER,default=file"`
}

func New() (*Config, error) {
	config := new(Config)
	err := load(config)
	return config, err
}
