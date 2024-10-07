package database

import (
	"github.com/centrex/webcore/core/env"
	"github.com/centrex/webcore/core/utils"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() {
	var err error

	dbType := env.GetEnv("DB_CONNECTION", "pgsql")

	dsn, err := utils.ConnectionURLBuilder(dbType)

	if err != nil {
		panic(err)
	}

	switch dbType {
	case "pgsql":
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	case "mysql":
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	default:
		panic("Unknown database type: " + dbType)
	}

	if err != nil {
		panic(err)
	}
}
