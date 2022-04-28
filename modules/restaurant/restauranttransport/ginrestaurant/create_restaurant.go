package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/modules/user/usermodel"

	"rest-api/common"
	"rest-api/component"
	"rest-api/modules/restaurant/restaurantbusiness"
	"rest-api/modules/restaurant/restaurantmodel"
	"rest-api/modules/restaurant/restaurantstore"
)

func CreateRestaurant(appContext component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var restaurant restaurantmodel.Restaurant
		if err := c.ShouldBind(&restaurant); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		user := c.MustGet("user").(*usermodel.User)
		restaurant.UserID = user.ID
		store := restaurantstore.NewSqlStore(appContext.GetMainDBConnection())
		RestaurantBusiness := restaurantbusiness.NewCreateRestaurantBusiness(store)
		if err := RestaurantBusiness.CreateRestaurant(c.Request.Context(), &restaurant); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(restaurant, nil, nil))
	}
}
