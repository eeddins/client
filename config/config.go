package config

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// Config represents a configuration with overrides from environment variables
// merged in.
type Config struct {
	CheckInterval time.Duration // unit in hours
	MaxAge        int           // unit in seconds
	GitHub        string        // the URL to github
	Organization  string        // the organization in GitHub to scan
	Repos         []string      // a comma-separated list of repos in Organization to scan
}

// NewConfig returns a new Config struct after looking for any overrides from
// envrionment variables.
func NewConfig() Config {
	c := Config{}
	c.setDefaults()
	return c
}

func (conf *Config) setDefaults() {
	var err error // is it bad to reuse this?

	checkInterval := os.Getenv("CHECK_INTERVAL")
	if checkInterval == "" {
		checkInterval = "5"
	}
	duration, err2 := strconv.Atoi(checkInterval)
	if err2 != nil {
		log.Print("Error setting CHECK_INTERVAL: ", err)
	}
	conf.CheckInterval = time.Hour * time.Duration(duration)

	maxAge := os.Getenv("MAX_AGE")
	if maxAge == "" {
		maxAge = "86400"
	}
	conf.MaxAge, err = strconv.Atoi(maxAge)
	if err != nil {
		log.Print("Error setting MAX_AGE: ", err)
	}

	gitHub := os.Getenv("GITHUB_URL")
	if gitHub == "" {
		gitHub = "https://github.com"
	}
	conf.GitHub = gitHub

	organization := os.Getenv("GITHUB_ORG")
	if organization == "" {
		organization = "eeddins"
	}
	conf.Organization = organization

	repos := os.Getenv("GITHUB_REPOS")
	if repos == "" {
		repos = "client"
	}
	conf.Repos = strings.Split(repos, ",")
}
