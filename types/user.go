package types

import (
	"context"

	"github.com/google/go-github/github"
)

// Users is the basic object to display the user details
type Users struct {
	Name string   `json:"Name" yaml:"Name"`
	Keys []string `json:"Keys" yaml:"Keys"`
}

// GetKeys will fetch the user's public ssh keys from Github
// You can limit how many keys are stored by the argument.
// If you set the limit to zero, it is considered to be unlimited.
func (u *Users) GetKeys(client *github.Client, limit int) error {
	opt := &github.ListOptions{}
	for {
		keys, resp, err := client.Users.ListKeys(context.Background(), u.Name, opt)
		if err != nil {
			return err
		}
		for _, key := range keys {
			if limit != 0 && len(u.Keys) == limit {
				return nil
			}
			u.Keys = append(u.Keys, key.GetKey())
		}
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return nil
}
