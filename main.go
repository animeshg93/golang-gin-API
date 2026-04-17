package main

import (
	"encoding/json"
	"errors"
	"example/vinyl-web/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/user/:username", func(c *gin.Context) {
		name := c.Param("username")
		gituser, err := fetchGitUser(name)

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

	router.Run(":8080")
}

func fetchGitUser(name string) (models.GitUser, error) {
	resp, err := http.Get("https://api.github.com/users/" + name)
	var gitUser models.GitUser

	if resp.StatusCode == http.StatusNotFound {
		return models.GitUser{}, errors.New("user not found")
	}
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("Error fetching data:", err)
		return models.GitUser{}, errors.New("failed to fetch git user")
	}

	err = json.NewDecoder(resp.Body).Decode(&gitUser)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		panic(err)
	}

	defer resp.Body.Close()
	return gitUser, nil

}
