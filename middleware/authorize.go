package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/component"
	"rest-api/component/tokenprovider/jwt"
	"rest-api/modules/user/userstorage"
	"strings"
)

func extractTokenFromHeaderAuthorization(token string) (string, error) {
	parts := strings.Split(token, " ")
	fmt.Println(token)
	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", errors.New("wrong authorization")
	}
	return parts[1], nil
}

func RequiredAuth(appCtx component.AppContext) func(ctx *gin.Context) {
	tokenProvider := jwt.NewJwtTokenProvider(appCtx.SecretKey())
	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderAuthorization(c.GetHeader("Authorization"))

		if err != nil {
			fmt.Println()
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		fmt.Println(c.GetHeader("Authorization"))

		db := appCtx.GetMainDBConnection()
		store := userstorage.NewSqlModel(db)

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		user, err := store.GetUser(c.Request.Context(), map[string]interface{}{"id": payload.UserID})
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if user.Status == 0 {
			c.JSON(http.StatusBadRequest, errors.New("user has been deleted").Error())
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
