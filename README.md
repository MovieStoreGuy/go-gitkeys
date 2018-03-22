# Keyobtainer
[![Build Status](https://travis-ci.org/MovieStoreGuy/keyobtainer.svg?branch=master)](https://travis-ci.org/MovieStoreGuy/keyobtainer)
[![Maintainability](https://api.codeclimate.com/v1/badges/bb13016d1510af20b550/maintainability)](https://codeclimate.com/github/MovieStoreGuy/keyobtainer/maintainability)
[![Go Report Card](https://goreportcard.com/badge/github.com/MovieStoreGuy/keyobtainer)](https://goreportcard.com/report/github.com/MovieStoreGuy/keyobtainer)  
Imagine you are trying to provide ssh access to your internal development team and you require to obtain
their public ssh key at varying times.
Now, imagine if you could automatically grab their most recent public ssh keys without having to ask them!  

THAT WOULD BE AMAZING!

Well, go no further, we have made exactly that!
This application fetches public ssh keys from Github to make it easier on your DevOps team.

# Installation
```sh
go get -u github.com/MovieStoreGuy/keyobtainer
```

# Usage
```
Usage of ./keyobtainer:
  -format string
    	The desired format for the output, they can be yaml, json or raw (default "raw")
  -limit int
    	Sets the limit as to how many keys to store, zero is unlimited
  -org string
    	The Github org that want to fetch all public users's public ssh keys
  -output string
    	Define the path you wish to output the content to
  -token string
    	A user's github token that can access the org's details
  -user string
    	The Github user to get their public ssh keys
```
