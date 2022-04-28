package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/modules/restaurantlike/restaurantlikestorage"
	"rest-api/modules/user/usermodel"

	"rest-api/common"
	"rest-api/component"
	"rest-api/modules/restaurant/restaurantbusiness"
	"rest-api/modules/restaurant/restaurantmodel"
	"rest-api/modules/restaurant/restaurantstore"
)

func ListRestaurant(appContext component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		user := c.MustGet("user").(*usermodel.User)

		var filter restaurantmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		paging.Fulfill()

		store := restaurantstore.NewSqlStore(appContext.GetMainDBConnection())
		likeStore := restaurantlikestorage.NewSqlStore(appContext.GetMainDBConnection())
		RestaurantBusiness := restaurantbusiness.NewListRestaurantBusiness(store, likeStore)
		restaurants, err := RestaurantBusiness.ListRestaurantByCondition(c.Request.Context(), map[string]interface{}{"owner_id": user.ID}, &filter, &paging)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(restaurants, paging, filter))
	}
}
