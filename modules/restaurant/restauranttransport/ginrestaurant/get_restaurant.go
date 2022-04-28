package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/common"
	"rest-api/component"
	"rest-api/modules/restaurant/restaurantbusiness"
	"rest-api/modules/restaurant/restaurantstore"
	"strconv"
)

func GetRestaurantByID(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		store := restaurantstore.NewSqlStore(appCtx.GetMainDBConnection())
		getRestaurantBusiness := restaurantbusiness.NewGetRestaurantBusiness(store)
		restaurant, err := getRestaurantBusiness.GetRestaurantByID(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(restaurant, nil, nil))
	}
}
