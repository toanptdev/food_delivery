package ginrestaurantlike

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/common"
	"rest-api/component"
	"rest-api/modules/restaurant/restaurantstore"
	"rest-api/modules/restaurantlike/restaurantlikebusiness"
	"rest-api/modules/restaurantlike/restaurantlikestorage"
	"rest-api/modules/user/usermodel"
	"strconv"
)

func UnlikeRestaurant(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		restaurantID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		user := c.MustGet("user").(*usermodel.User)

		store := restaurantlikestorage.NewSqlStore(appCtx.GetMainDBConnection())
		decreaseStore := restaurantstore.NewSqlStore(appCtx.GetMainDBConnection())
		unlikeRestaurantBusiness := restaurantlikebusiness.NewUserUnlikeRestaurant(store, decreaseStore)

		err = unlikeRestaurantBusiness.Unlike(c.Request.Context(), user.ID, restaurantID)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse("ok", nil, nil))
	}
}
