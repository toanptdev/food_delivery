package ginrestaurantlike

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/common"
	"rest-api/component"
	"rest-api/modules/restaurantlike/restaurantlikebusiness"
	"rest-api/modules/restaurantlike/restaurantlikemodel"
	"rest-api/modules/restaurantlike/restaurantlikestorage"
	"rest-api/modules/user/usermodel"
	"strconv"
)

// Post /restaurants/:id/like
func LikeRestaurant(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		restaurantID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}

		user := c.MustGet("user").(*usermodel.User)

		store := restaurantlikestorage.NewSqlStore(appCtx.GetMainDBConnection())
		likeRestaurantBusiness := restaurantlikebusiness.NewUserLikeRestaurant(store)

		data := restaurantlikemodel.RestaurantLike{
			RestaurantID: restaurantID,
			UserID:       user.ID,
		}

		err = likeRestaurantBusiness.Like(c.Request.Context(), &data)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse("ok", nil, nil))
	}
}
