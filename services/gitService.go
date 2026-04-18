package services

import (
	"encoding/json"
	"errors"
	"example/vinyl-web/models"
	"fmt"
	"net/http"
)

func FetchGitUser(name string) (models.GitUser, error) {
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

func FetchGitRepos(name string) ([]models.GitRepo, error) {
	resp, err := http.Get("https://api.github.com/users/" + name + "/repos")
	var gitRepos []models.GitRepo

	if resp.StatusCode == http.StatusNotFound {
		return []models.GitRepo{}, errors.New("user not found")
	}
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("Error fetching data:", err)
		return []models.GitRepo{}, errors.New("failed to fetch git repos")
	}

	err = json.NewDecoder(resp.Body).Decode(&gitRepos)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		panic(err)
	}

	defer resp.Body.Close()
	return gitRepos, nil

}
