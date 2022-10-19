package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("conf")     // 文件名
	viper.SetConfigType("yaml")     // 文件类型
	viper.AddConfigPath("./config") // 文件地址 ( 项目地址 + 后缀)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	//fmt.Println(viper.Get("log"))
	//fmt.Println(viper.Get("log.age"))
}
