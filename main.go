package main

import (
	"encoding/json"
	"example/vinyl-web/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	gituser := fetchGitUser()

	router.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gituser)
	})

	router.Run(":8080")
}

func fetchGitUser() models.GitUser {
	resp, err := http.Get("https://api.github.com/users/animeshg93")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		panic("ERRORRRRR")
	}
	defer resp.Body.Close()

	var gitUser models.GitUser
	err = json.NewDecoder(resp.Body).Decode(&gitUser)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		panic(err)
	}
	return gitUser
}
