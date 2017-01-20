package main

import (
	"fmt"
	"time"

	"github.com/google/go-github/github"
)

func main() {
	client := github.NewClient(nil)
	owner := "eeddins"
	repo := "test"
	var now = time.Now().Unix()

	// pulls attributes from PullRequests.go
	prlist, _, err := client.PullRequests.List(owner, repo, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, prl := range prlist {
		if now-prl.UpdatedAt.Unix() > 200 {
			fmt.Printf("User who initiated pull %v\n", *prl.User.Login)
			fmt.Printf("Current state of pull %v\n", *prl.State)
			fmt.Printf("Pull request created at %v\n", prl.CreatedAt)
			fmt.Printf("Pull request updated at %v\n", prl.UpdatedAt)
		}
	}
}
