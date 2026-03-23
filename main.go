package main

import(
	"github.com/tejutejas05/Ecommerce-Website-using-GO-/controllers"
	"github.com/tejutejas05/Ecommerce-Website-using-GO-/database"
	"github.com/tejutejas05/Ecommerce-Website-using-GO-/middleware"
	"github.com/tejutejas05/Ecommerce-Website-using-GO-/routes"
	//"github.com/tejutejas05/Ecommerce-Website-using-GO-/models"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getevn("PORT")
	if port == ""{
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))
}

