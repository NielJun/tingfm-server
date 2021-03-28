package main

import (
	"fmt"
	"tingfm/config"
	"tingfm/dao"
	"tingfm/routers"
)

func init() {
	///配置初始化
	config.Init()
	///数据库初始化
	dao.Init()
}

func main() {
	r := routers.SetupRouter()
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v \n", err)
	}

}
