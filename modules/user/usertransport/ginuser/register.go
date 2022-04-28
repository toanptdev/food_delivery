package ginuser

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/component"
	"rest-api/component/hasher"
	"rest-api/modules/user/userbusiness"
	"rest-api/modules/user/usermodel"
	"rest-api/modules/user/userstorage"
)

func Register(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var data usermodel.UserCreate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		db := appCtx.GetMainDBConnection()
		store := userstorage.NewSqlModel(db)
		md5 := hasher.NewMD5Hash()
		userBusiness := userbusiness.NewUserBusiness(store, md5)
		if err := userBusiness.Register(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, "message: ok")
	}
}
