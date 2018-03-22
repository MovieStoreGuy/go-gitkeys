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
