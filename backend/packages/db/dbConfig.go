package db

import (
	"fmt"
	"go-todo/packages/config"
)

func DSN() string {
	appDbConfig := config.Get().Db
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		appDbConfig.User, appDbConfig.Password, appDbConfig.Host, appDbConfig.Port, appDbConfig.DbName)
}
