package routes

import (
	"example/vinyl-web/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var gitUserURL = "/user/:username"

func GitUserRoute(router *gin.Engine) {
	router.GET(gitUserURL, func(c *gin.Context) {
		name := c.Param("username")
		gituser, err := services.FetchGitUser(name)

		if err != nil && err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gituser)
	})
}
