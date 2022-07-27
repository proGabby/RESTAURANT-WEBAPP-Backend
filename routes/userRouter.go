package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/willie/restaurant-backend/controller"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	//gin.Engine is an instance of the gin Engine. it contains the muxer, middleware and configuration settings.

	//handling request on routes
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
	incomingRoutes.POST("/users/signup", controller.SignUp())
	incomingRoutes.POST("/users/login", controller.Login())
}
