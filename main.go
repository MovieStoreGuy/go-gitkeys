package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"

	"github.com/MovieStoreGuy/keyobtainer/engine"
	"github.com/MovieStoreGuy/keyobtainer/output"
)

var (
	githubOrg, githubToken string
	user, team             string
	limit                  int

	out    io.Writer = os.Stdout
	outdir string
	format string = "raw"
)

func init() {
	const (
		base = ""
	)
	flag.StringVar(&user, "user", base, "The Github user to get their public ssh keys")
	flag.StringVar(&githubOrg, "org", base, "The Github org that want to fetch all public users's public ssh keys")
	flag.StringVar(&githubToken, "token", base, "A user's github token that can access the org's details")
	flag.StringVar(&format, "format", format, "The desired format for the output, they can be yaml, json or raw")
	flag.StringVar(&outdir, "output", base, "Define the path you wish to output the content to")
	flag.StringVar(&team, "team", base, "Define the team to filter results by")
	flag.IntVar(&limit, "limit", 0, "Sets the limit as to how many keys to store, zero is unlimited")
}

func main() {
	flag.Parse()
	members, err := engine.CreateEngine(githubToken, githubOrg, user, team).GetUsers(limit)
	if err != nil {
		log.Fatal("Unable to fetch users due to ", err)
	}
	if outdir != "" {
		if _, err := os.Stat(outdir); !os.IsNotExist(err) {
			log.Fatal("File already exists")
		}
		f, err := os.OpenFile(outdir, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}
		out = bufio.NewWriter(f)
	}
	printer, err := output.CreatePrinter(out)
	if err != nil {
		log.Fatal(err)
	}
	if err = printer.Print(format, members); err != nil {
		log.Fatal(err)
	}
}
