package config

import (
	"github.com/joho/godotenv"
)

func LoadEnv() {
	godotenv.Load()
	appConfig.Init()
}
