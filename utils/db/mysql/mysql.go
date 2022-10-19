package mysql

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type config struct {
	User     string
	Password string
	Port     string
	Database string
}

// read log config
func readConfig() config {
	conf := config{}
	conf.User = viper.Get("mysql.user").(string)
	conf.Password = viper.Get("mysql.password").(string)
	conf.Port = viper.Get("mysql.port").(string)
	conf.Database = viper.Get("mysql.database").(string)
	return conf
}

func InitMysql() {
	conf := readConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.User, conf.Password, conf.Port, conf.Database)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("Fatal error mysql init: %s \n", err))
	}
}
