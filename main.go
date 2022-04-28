package main

import (
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"rest-api/middleware"
	"rest-api/modules/restaurantlike/restaurantliketransport/ginrestaurantlike"

	//"rest-api/middleware"
	"rest-api/modules/user/usertransport/ginuser"

	"rest-api/component"
	"rest-api/modules/restaurant/restauranttransport/ginrestaurant"
	//ginU "rest-api/modules/user/usertranport/ginuser"
	//"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:admin123@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln("connect to db failed")
	}

	err = runService(db)
	if err != nil {
		log.Fatalln("run func failed")
	}
}

func runService(db *gorm.DB) error {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	secretKey := "toanpt"

	appCtx := component.NewAppCtx(db, secretKey)

	r.POST("/register", ginuser.Register(appCtx))
	r.POST("/login", ginuser.Login(appCtx))
	r.GET("/profile", middleware.RequiredAuth(appCtx), ginuser.Profile(appCtx))
	restaurants := r.Group("/restaurants", middleware.RequiredAuth(appCtx))
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.GET(":id", ginrestaurant.GetRestaurantByID(appCtx))
		restaurants.PUT(":id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurants.DELETE(":id", ginrestaurant.DeleteRestaurant(appCtx))

		restaurants.POST(":id/like", ginrestaurantlike.LikeRestaurant(appCtx))
		restaurants.DELETE(":id/unlike", ginrestaurantlike.UnlikeRestaurant(appCtx))
	}

	return r.Run()
}
