package db

import (
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Engine *xorm.Engine

func InitDB() error {
	var err error

	dbName := viper.GetString("db.name")
	dbUser := viper.GetString("db.user")
	dbPassword := viper.GetString("db.password")
	dbHost := strings.Split(viper.GetString("db.addr"), ":")[0]
	dbPort := strings.Split(viper.GetString("db.addr"), ":")[1]

	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8", dbUser, dbPassword, dbHost, dbPort, dbName)
	Engine, err = xorm.NewEngine("mysql", dataSourceName)
	if err = Engine.DB().Ping(); err != nil {
		logrus.Errorf("数据库连接失败 " + err.Error())
		return err
	}
	Engine.DB().SetMaxIdleConns(10)
	Engine.DB().SetMaxOpenConns(10)
	logrus.Info("数据库连接成功")
	return nil
}
