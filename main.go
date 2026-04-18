package main

import (
	"example/vinyl-web/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/user/:username", routes.GetGitUser)
	router.GET("/repos/:username", routes.GetGitRepos)

	router.Run(":8080")
}
