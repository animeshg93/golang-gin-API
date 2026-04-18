package main

import (
	"example/vinyl-web/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.GitUserRoute(router)

	router.Run(":8080")
}
