package main

import (
	"log"

	"github.com/gin-gonic/gin"
	config "github.com/rizperdana/prototype-golang-rest/configs"
	routes "github.com/rizperdana/prototype-golang-rest/routes"
)

func main() {
	// init database
	config.Connect()

	// init router
	router := gin.Default()
	routes.Routes(router)
	log.Fatal(router.Run(":8800"))
}
