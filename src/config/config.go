package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// 读取配置文件config
type TingFMConfig struct {
	Redis string
	MySQL MySQLConfig
}

type MySQLConfig struct {
	Host         string
	Port         int
	Username     string
	Password     string
	DatabaseName string
}

var Config *TingFMConfig

func Init() {
	// 把配置文件读取到结构体上

	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	viper.Unmarshal(&Config) //将配置文件绑定到config上
	fmt.Println("config: ", Config, "redis: ", Config.Redis, "mysql remote: ", Config.MySQL)
}
