package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/MovieStoreGuy/keyobtainer/engine"
)

var (
	githubOrg   string
	githubToken string
	user        string
)

func init() {
	const (
		base = ""
	)
	flag.StringVar(&user, "user", base, "The Github user to get their public ssh keys")
	flag.StringVar(&githubOrg, "org", base, "The Github org that want to fetch all public users's public ssh keys")
	flag.StringVar(&githubToken, "token", base, "A user's github token that can access the org's details")
}

func main() {
	flag.Parse()
	members, err := engine.CreateEngine(githubToken, githubOrg, user).GetUsers()
	if err != nil {
		log.Fatal("Unable to fetch users due to", err)
	}
	for _, member := range members {
		fmt.Printf("Member is: %+v\n", member)
	}
}
