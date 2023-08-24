package db

import (
	"Task-Management/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DbClient *gorm.DB

func DbConnect() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.EnvironmentData.DbUser,
		config.EnvironmentData.DbPass,
		config.EnvironmentData.DbUrl,
		config.EnvironmentData.DbPort,
		config.EnvironmentData.DbName,
	)

	mysql := mysql.Open(dsn)
	db, err := gorm.Open(mysql, &gorm.Config{})

	if err != nil {
		fmt.Println("Error while connecting to DB")
	}
	DbClient = db
}
