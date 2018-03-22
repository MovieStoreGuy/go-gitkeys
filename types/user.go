package types

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	endpoint = "https://github.com/%s.keys"
)

type Users struct {
	Name string   `json:"Name" yaml:"Name"`
	Keys []string `json:"Keys" yaml:"Keys"`
}

// GetKeys will fetch the user's public ssh keys from Github
func (u *Users) GetKeys() error {
	resp, err := http.Get(fmt.Sprintf(endpoint, u.Name))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	for _, key := range strings.Split(string(body), "\n") {
		if len(key) == 0 {
			continue
		}
		u.Keys = append(u.Keys, key)
	}
	return nil
}