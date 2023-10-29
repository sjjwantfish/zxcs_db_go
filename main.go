package main

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/sjjwantfish/zxcs_db_go/config"
	"github.com/sjjwantfish/zxcs_db_go/libs/db"
	"github.com/sjjwantfish/zxcs_db_go/router"
)

var confFileName = pflag.StringP("config", "f", "", "config file path.")

func main() {
	pflag.Parse()
	filePath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	// 初始化
	if err := config.InitLog(filePath); err != nil {
		panic(err)
	}
	if err := config.InitConfig(*confFileName); err != nil {
		panic(err)
	}

	//if err := redis.Init(); err != nil {
	//	panic(err)
	//	return
	//}

	if err := db.InitDB(); err != nil {
		panic(err)
	}

	// libs.DescribeEc2Instances()

	// 加载转发路径
	g := gin.New()
	router.Load(
		g,
	)

	_ = g.Run(viper.GetString("addr"))
}
