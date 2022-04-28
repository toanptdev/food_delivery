package ginuser

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/component"
)

func Profile(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		user := c.MustGet("toanpt")

		c.JSON(http.StatusOK, user)
	}
}
