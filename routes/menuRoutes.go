package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/willie/restaurant-backend/controller"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menus", controller.GetMenus())
	incomingRoutes.GET("/menus/:menu_id", controller.GetMenu())
	incomingRoutes.POST("/menus", controller.CreateMenu())
	incomingRoutes.PATCH("/menus/:menu_id", controller.UpdateMenu())
}