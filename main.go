package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
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
	var authClient *http.Client
	if githubToken != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: githubToken},
		)
		authClient = oauth2.NewClient(context.Background(), ts)
	}
	client := github.NewClient(authClient)

	if githubOrg != "" {
		for {
			members, resp, err := client.Organizations.ListMembers(context.Background(), githubOrg, &github.ListMembersOptions{})
			if err != nil {
				log.Fatal(err)
			}
			// Process data we have
			if resp.NextPage == 0 {
				break
			}
			for _, member := range members {
				fmt.Println("The user:", member.GetLogin())
			}
		}
	}
}
