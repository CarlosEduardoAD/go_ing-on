package main

import (
	"log"

	"github.com/CarlosEduardoAD/Go-ing_on/app/routes"
	"github.com/CarlosEduardoAD/Go-ing_on/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	_, err := utils.ConnectToMongo()

	if err != nil {
		log.Println(err)
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))

	routes.Routes(app)

	app.Run(":9898")
}
