package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/willie/restaurant-backend/controllers"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	//handling request on routes
	incomingRoutes.GET("/foods", controller.GetFoods())
	incomingRoutes.GET("/foods/:food_id", controller.GetFood())
	incomingRoutes.POST("/foods", controller.CreateFood())
	incomingRoutes.PATCH("/foods/:food_id", controller.UpdateFood())
}
