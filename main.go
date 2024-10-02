package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const (
	getFollowers = "https://api.github.com/user/followers"
	getFollowing = "https://api.github.com/user/following"
)

type users struct {
	Login string `json:"Login"`
}

func fetchGithubData(url string) ([]users, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", getFollowers, nil)
	if err != nil {
		fmt.Println("Error creating request")
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+os.Getenv("TOKEN"))
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error sending request!")
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println("error !")
		return nil, fmt.Errorf("Error aaya!")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response!")
		return nil, err
	}
	var user []users
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, err
	}
	return user, nil
}

// main function
func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error loading the file")
	}
	// token := os.Getenv("TOKEN")
	followers, err := fetchGithubData(getFollowers)
	following, err := fetchGithubData(getFollowing)
	fmt.Printf("People you are following: %v\n", followers)
	fmt.Printf("People following you: %v\n", following)
	fmt.Println(followers)
}
