package engine

import (
	"context"
	"net/http"

	"github.com/MovieStoreGuy/keyobtainer/types"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type GitHub struct {
	organisation string
	token        string
	user         string
	client       *github.Client
}

func CreateEngine(token, org, user string) *GitHub {
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
	}
}

func (g *GitHub) GetUsers() ([]types.Users, error) {
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
				if err := user.GetKeys(); err != nil {
					return nil, err
				}
				// Only output return users that have keys defined
				if len(user.Keys) != 0 {
					users = append(users, user)
				}
			}
			// Process data we have
			if resp.NextPage == 0 {
				break
			}
			opt.Page = resp.NextPage
		}
	default:
		user := types.Users{
			Name: g.user,
		}
		if err := user.GetKeys(); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
