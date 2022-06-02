package utils

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)
var Db *gorm.DB
func Connect()  {
	config := Config
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",config.Mysql.Username,config.Mysql.Password,config.Mysql.Host,config.Mysql.Port,config.Mysql.Database)
	db, err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		log.Fatal("connect mysql fail",err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("connect pool fail:",err.Error())
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	Db = db
	fmt.Println(Db)
}