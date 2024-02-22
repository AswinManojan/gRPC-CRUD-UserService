package db

import (
	"fmt"
	"log"

	config "github.com/grpc/gobus/user/Config"
	model "github.com/grpc/gobus/user/pkg/Model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect_DB(cfg *config.Config) *gorm.DB {
	host := cfg.HOST
	username := cfg.USERNAME
	password := cfg.PASSWORD
	port := cfg.PORT
	sslmode := cfg.SSLMODE
	dbname := cfg.DBNAME
	dsn := fmt.Sprintf("host=%s user= %s password=%s dbname= %s port=%s sslmode= %s", host, username, password, dbname, port, sslmode)
	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to DB")
	}
	Db.AutoMigrate(&model.User{})
	return Db
}
