package api

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/sjjwantfish/zxcs_db_go/handler/resp"
	"github.com/sjjwantfish/zxcs_db_go/libs/db"
)

func GetBookTitleByKindID(c *gin.Context) {
	kindIDStr := c.DefaultQuery("kind_id", "0")
	kindID, err := strconv.Atoi(kindIDStr)
	if err != nil {
		logrus.Error(err)
	}
	logrus.Info(fmt.Sprintf("get kindId: %v", kindID))
	titles, err := db.GetBookTitleByKind(int64(kindID))
	if err != nil {
		logrus.Error(err)
	}
	resp.Success(c, gin.H{"data": titles})
}
