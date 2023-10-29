package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, resp interface{}) {
	c.JSON(http.StatusOK, gin.H{"retcode": 0, "resp": resp, "request": c.Request.URL.Path})
}
