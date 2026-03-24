package main

import(
	"github.com/tejutejas05/Ecommerce-Website-using-GO-/controllers"
	"github.com/tejutejas05/Ecommerce-Website-using-GO-/database"
	"github.com/tejutejas05/Ecommerce-Website-using-GO-/middleware"
	"github.com/tejutejas05/Ecommerce-Website-using-GO-/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getevn("PORT")
	if port == ""{
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(':' + port))

}
