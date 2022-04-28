package ginuser

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/component"
	"rest-api/component/hasher"
	"rest-api/component/tokenprovider/jwt"
	"rest-api/modules/user/userbusiness"
	"rest-api/modules/user/usermodel"
	"rest-api/modules/user/userstorage"
)

func Login(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var data usermodel.UserLogin
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		db := appCtx.GetMainDBConnection()
		store := userstorage.NewSqlModel(db)
		md5 := hasher.NewMD5Hash()
		tokenProvider := jwt.NewJwtTokenProvider(appCtx.SecretKey())
		loginBusiness := userbusiness.NewLoginBusiness(store, md5, tokenProvider, 60*60*24*30)

		token, err := loginBusiness.Login(c.Request.Context(), &data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, token)
	}
}
