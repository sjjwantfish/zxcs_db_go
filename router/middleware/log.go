package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 为访问日志加ip。
// 修改logrus为修改的Logger

var myLog *logrus.Entry

func LogIp(c *gin.Context) {
	myLog = logrus.WithFields(logrus.Fields{
		//"time1": time.Now(),
		"url": c.Request.URL.Path,
		"ip":  strings.Split(c.ClientIP(), ":")[0],
	})
}

func Info(rgs ...interface{}) {
	myLog.Info(rgs...)
}

func Warn(rgs ...interface{}) {
	myLog.Warn(rgs...)
}

func Error(rgs ...interface{}) {
	myLog.Error(rgs...)
}
