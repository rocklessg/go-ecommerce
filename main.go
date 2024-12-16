package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rocklessg/go-ecommerce/controllers"
	"github.com/rocklessg/go-ecommerce/database"
	"github.com/rocklessg/go-ecommerce/middleware"
	"github.com/rocklessg/go-ecommerce/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == ""{
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/listcart", controllers.GetItemFromCart())
	router.POST("/addaddress", controllers.AddAdress())
	router.PUT("/edithomeaddress", controllers.EditHomeAddAdress())
	router.PUT("/editworkaddress", controllers.EditWorkAddAdress())
	router.GET("/deleteaddresses", controllers.DeleteAdress())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}

//resource: https://www.youtube.com/watch?v=ry0uQ66n5aE&list=PL5dTjWUk_cPaf5uSEmr8ilR-GtO6s7QJJ