package config

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

func InitConfig(name string) error {
	if name != "" {
		viper.SetConfigFile(name)
	} else {
		viper.SetConfigFile("conf/config.yaml")
	}
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APISERVER")

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		logrus.Error("加载日志异常，", err.Error())
		return err
	}
	logrus.Info("配置加载完成")
	if viper.GetString("run_mode") != "debug" {
		gin.SetMode(gin.ReleaseMode)
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetLevel(logrus.DebugLevel)
	}

	return nil
}
