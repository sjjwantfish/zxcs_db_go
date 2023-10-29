package router

import (
	"fmt"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/sjjwantfish/zxcs_db_go/handler/api"
	"github.com/sjjwantfish/zxcs_db_go/handler/resp"
	"github.com/sjjwantfish/zxcs_db_go/router/middleware"
)

func RecoverErr() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				const size = 4096
				buf := make([]byte, size)
				buf = buf[:runtime.Stack(buf, false)]

				logrus.WithFields(logrus.Fields{
					"url": c.Request.URL.Path,
				}).Error("发生错误:", err, "  方法: ", c.HandlerName(), " stack: ", string(buf))
				c.AbortWithStatusJSON(
					500,
					gin.H{
						"error":   "sorry, we made a mistake!",
						"retcode": 1001,
						"url":     c.Request.URL.Path,
					},
				)
			}
		}()
		c.Next()
	}
}

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	// g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	// g.Use(middleware.LogIp)

	g.Use(RecoverErr())
	g.RedirectTrailingSlash = false
	g.HandleMethodNotAllowed = true
	g.NoMethod(func(c *gin.Context) {
		c.JSON(
			404,
			gin.H{"error": "method not allowed", "retcode": 1005, "request": c.Request.URL.Path},
		)
	})
	g.NoRoute(func(c *gin.Context) {
		c.JSON(
			404,
			gin.H{"error": "route not found", "retcode": 1007, "request": c.Request.URL.Path},
		)
	})
	g.GET("/healthz", func(c *gin.Context) {
		resp.Success(c, gin.H{"status": "ok"})
		return
	})
	g.GET("/_count", func(c *gin.Context) {
		resp.Success(c, gin.H{"goroutine count": runtime.NumGoroutine()})
		return
	})

	// 注册路由
	book := g.Group("/api/v1/book")
	{
		book.GET("/titles", api.GetBookTitleByKindID)
	}
	// ec2 := g.Group("/api/v1/ec2")
	// {
	// 	ec2.GET("", api.AwsAssetsGet)
	// 	ec2.GET("/vpc", api.AwsAssetsGet)
	// 	ec2.GET("/subnet", api.AwsAssetsGet)
	// 	ec2.GET("/rtb", api.AwsAssetsGet)
	// 	ec2.GET("/igw", api.AwsAssetsGet)
	// 	ec2.GET("/ngw", api.AwsAssetsGet)
	// 	ec2.GET("/eip", api.AwsAssetsGet)
	// }
	// elb := g.Group("/api/v1/elb")
	// {
	// 	elb.GET("", api.AwsAssetsGet)
	// 	elb.GET("/listener", api.AwsAssetsGet)
	// 	elb.GET("/targetGroup", api.AwsAssetsGet)
	// }

	// s3 := g.Group("/api/v1/s3")
	// {
	// 	s3.GET("/bucket", api.AwsAssetsGet)
	// }
	// eks := g.Group("/api/v1/eks")
	// {
	// 	eks.GET("", api.AwsAssetsGet)
	// }
	// rds := g.Group("/api/v1/rds")
	// {
	// 	rds.GET("", api.AwsAssetsGet)
	// 	rds.GET("/instance", api.AwsAssetsGet)
	// }

	return g
}
