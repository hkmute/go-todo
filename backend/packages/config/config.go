package config

import "os"

type appConfigStruct struct {
	Port      string
	JwtSecret string
	Db        db
}

type db struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

var appConfig appConfigStruct

func (c appConfigStruct) Init() {
	appConfig = appConfigStruct{
		Port:      os.Getenv("PORT"),
		JwtSecret: os.Getenv("JWT_SECRET"),
		Db: db{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DbName:   os.Getenv("DB_NAME"),
		},
	}
}

func Get() appConfigStruct {
	return appConfig
}
