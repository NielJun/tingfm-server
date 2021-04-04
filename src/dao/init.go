package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"tingfm/config"
	"tingfm/do"
)

var DB *gorm.DB

///初始化
func Init() {

	var mysqlConnectMsg = fmt.Sprintf("%s:%s@(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config.MySQL.Username, config.Config.MySQL.Password, config.Config.MySQL.Host, config.Config.MySQL.Port, config.Config.MySQL.DatabaseName)

	db, err := gorm.Open("mysql", mysqlConnectMsg)

	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	db.AutoMigrate(&do.Album{})
	db.AutoMigrate(&do.Category{})
	db.AutoMigrate(&do.Money{})
	DB = db

}
