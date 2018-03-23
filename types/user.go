package types

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	endpoint = "https://github.com/%s.keys"
)

// Users is the basic object to display the user details
type Users struct {
	Name string   `json:"Name" yaml:"Name"`
	Keys []string `json:"Keys" yaml:"Keys"`
}

// GetKeys will fetch the user's public ssh keys from Github
// You can limit how many keys are stored by the argument.
// If you set the limit to zero, it is considered to be unlimited.
func (u *Users) GetKeys(limit int) error {
	resp, err := http.Get(fmt.Sprintf(endpoint, u.Name))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("Invalid request was sent to Github")
	}
	for index, key := range strings.Split(string(body), "\n") {
		if len(key) == 0 {
			continue
		}
		if limit != 0 && limit == index {
			break
		}
		u.Keys = append(u.Keys, key)
	}
	return nil
}
