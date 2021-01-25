package dao

import (
	"fmt"
	"tingfm/config"
	"tingfm/do"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

///初始化
func Init() {

	var mysqlConnectMsg = fmt.Sprintf("%s:%s@(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config.MySQL.Username, config.Config.MySQL.Password, config.Config.MySQL.Host, config.Config.MySQL.Port, config.Config.MySQL.DatabaseName)

	db, err := gorm.Open("mysql", mysqlConnectMsg)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&do.Albumn{})

	var albumn = do.Albumn{
		AlbumnName:     "三国演义",
		AlbumnImageUrl: "三国演义.png",
		AlbumnPlayer:   "袁阔成",
		Description:    "三国演义精品评书",
		ItemNumber:     365,
	}

	db.Debug().Create(&albumn)
}
