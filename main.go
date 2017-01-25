package main

import (
	"fmt"
	"log"
	"time"

	"github.com/eeddins/client/config"
	"github.com/google/go-github/github"
)

var conf config.Config

// TODO: move this out into its own package.
func listPulls(repo string) {
	log.Printf("Checking pulls in %v for %v\n", conf.Organization, repo)
	client := github.NewClient(nil)
	owner := conf.Organization
	var now = time.Now().Unix()

	// pulls attributes from PullRequests.go
	prlist, _, err := client.PullRequests.List(owner, repo, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, prl := range prlist {
		if age := now - prl.UpdatedAt.Unix(); age > int64(conf.MaxAge) {
			fmt.Printf("User who initiated pull %v\n", *prl.User.Login)
			fmt.Printf("Current state of pull %v\n", *prl.State)
			fmt.Printf("Pull request created at %v\n", prl.CreatedAt)
			fmt.Printf("Pull request updated at %v\n", prl.UpdatedAt)
		}
	}
}

func main() {
	conf = config.NewConfig()
	for {
		for _, r := range conf.Repos {
			listPulls(r)
		}
		time.Sleep(conf.CheckInterval)
	}
}
