package models

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	var err error
	var dsn = "%s:%s@(%s)/%s?charset=utf8mb4&parseTime=true"
	Db, err = gorm.Open(mysql.Open(fmt.Sprintf(dsn, user, password, host, dbname)), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("数据库连接失败：%s\n", err.Error()))
	}
}
