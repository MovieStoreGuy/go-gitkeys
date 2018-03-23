package engine

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/MovieStoreGuy/keyobtainer/types"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// GitHub stores all the required items
// to be able to fetch all users
type GitHub struct {
	organisation string
	team         string
	token        string
	user         string
	client       *github.Client
}

// CreateEngine makes the engine of the given settings
// If the token is defined, then it will create a secure connect
// to Github.
func CreateEngine(token, org, user, team string) *GitHub {
	var authClient *http.Client
	if token != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		authClient = oauth2.NewClient(context.Background(), ts)
	}
	return &GitHub{
		client:       github.NewClient(authClient),
		token:        token,
		user:         user,
		organisation: org,
		team:         team,
	}
}

// GetUsers will return all the users the engine is configured
// to fetch.
func (g *GitHub) GetUsers(limit int) ([]types.Users, error) {
	users := []types.Users{}
	switch {
	case g.organisation != "":
		opt := &github.ListMembersOptions{}
		for {
			members, resp, err := g.client.Organizations.ListMembers(context.Background(), g.organisation, opt)
			if err != nil {
				return nil, err
			}
			for _, member := range members {
				user := types.Users{
					Name: member.GetLogin(),
				}
				// Only output return users that have keys defined
				users = append(users, user)
			}
			// Process data we have
			if resp.NextPage == 0 {
				break
			}
			opt.Page = resp.NextPage
		}
	default:
		users = append(users, types.Users{
			Name: g.user,
		})
	}
	if g.team != "" {
		collection, err := g.filterBy(g.team, users)
		if err != nil {
			return nil, err
		}
		users = collection
	}
	collection := []types.Users{}
	for _, user := range users {
		if err := user.GetKeys(limit); err != nil {
			return nil, err
		}
		if len(user.Keys) != 0 {
			collection = append(collection, user)
		}
	}
	return collection, nil
}

func (g *GitHub) filterBy(teamName string, users []types.Users) ([]types.Users, error) {
	// Lets cache all the teams
	lookup := map[string]int64{}
	opt := &github.ListOptions{}
	for {
		teams, resp, err := g.client.Organizations.ListTeams(context.Background(), g.organisation, opt)
		if err != nil {
			return nil, err
		}
		for _, team := range teams {
			lookup[strings.ToLower(team.GetName())] = team.GetID()
		}
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	if _, exist := lookup[strings.ToLower(teamName)]; !exist {
		return nil, errors.New("The team doesn't exist inside the organisiation")
	}
	// Filter the teams
	collection := []types.Users{}
	for _, user := range users {
		member, _, err := g.client.Organizations.IsTeamMember(context.Background(), lookup[strings.ToLower(teamName)], user.Name)
		if err != nil {
			return nil, err
		}
		if member {
			collection = append(collection, user)
		}
	}
	return collection, nil
}
