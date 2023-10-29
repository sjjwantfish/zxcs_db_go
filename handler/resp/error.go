package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServerError(c *gin.Context, resp interface{}) {
	c.JSON(http.StatusInternalServerError, gin.H{"retcode": 1001, "request": c.Request.URL.Path, "error": resp})
}

func InvalidParams(c *gin.Context, resp interface{}) {
	c.JSON(http.StatusBadRequest, gin.H{"retcode": 1004, "error": "missing parameters", "resp": resp, "request": c.Request.URL.Path})
}
