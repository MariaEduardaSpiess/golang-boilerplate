package database

import (
	"fmt"
	"golang-boilerplate/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := getDsn()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	return db
}

func getDsn() string {
	env := config.Env.Database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", env.Host, env.User, env.Password, env.DbName, env.Port, env.SslMode, env.TimeZone)
	return dsn
}
