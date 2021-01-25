package main

import (
	"tingfm/config"
	"tingfm/dao"
)

func init() {

	///配置初始化
	config.Init()

	///数据库初始化
	dao.Init()

}

func main() {

}
