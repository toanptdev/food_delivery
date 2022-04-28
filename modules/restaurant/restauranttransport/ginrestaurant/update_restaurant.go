package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/common"
	"rest-api/component"
	"rest-api/modules/restaurant/restaurantbusiness"
	"rest-api/modules/restaurant/restaurantmodel"
	"rest-api/modules/restaurant/restaurantstore"
	"strconv"
)

func UpdateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var restaurant restaurantmodel.RestaurantUpdate

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		if err := c.ShouldBind(&restaurant); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		store := restaurantstore.NewSqlStore(appCtx.GetMainDBConnection())
		updateRestaurantBusiness := restaurantbusiness.NewUpdateRestaurantBusiness(store)

		if err := updateRestaurantBusiness.UpdateRestaurant(c.Request.Context(), uint(id), &restaurant); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(gin.H{"message": "success"}, nil, nil))
	}
}
