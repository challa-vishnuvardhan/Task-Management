package config

import (
	"os"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
)

// struct to store all env variables
type EnvVariables struct {
	DbUrl  string
	DbUser string
	DbName string
	DbPass string
	DbPort string
}

var EnvironmentData *EnvVariables

func IntializeEnv() {
	_ = godotenv.Load("./Config.env")

	EnvironmentData = &EnvVariables{
		DbUrl:  os.Getenv("DB_URL"),
		DbUser: os.Getenv("DB_USER"),
		DbName: os.Getenv("DB_NAME"),
		DbPass: os.Getenv("DB_PASS"),
		DbPort: os.Getenv("DB_PORT"),
	}
	v := validator.New()
	err := v.Struct(*EnvironmentData)
	if err != nil {
		panic(err)
	}
}

func init() {
	IntializeEnv()
}
