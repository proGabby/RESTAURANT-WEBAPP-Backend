package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	routes "github.com/willie/restaurant-backend/routes"

	database "github.com/willie/restaurant-backend/database"
)

var foodMongoCollection = database.OpenCollection(database.Client, "Food-Items")

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loadin .env file")
	}

	Port, isAvail := os.LookupEnv("PORT")

	if !isAvail {
		log.Fatal("PORT key is not and env varialbe")
	}

	//create a gin router engine without middleware
	router := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter
	router.Use(gin.Logger())

	//Handling user routes
	routes.UserRoutes(router)

	//TODO: Handling authentication

	//Handling routes
	routes.FoodRoutes(router)
	routes.InvoiceRoutes(router)
	routes.OrderItemRoutes(router)
	routes.MenuRoutes(router)
	routes.OrderRoutes(router)
	routes.TableRoutes(router)

	//starts server
	router.Run(":" + Port)

}
