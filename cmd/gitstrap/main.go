package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	log.SetPrefix("")
	log.SetFlags(0)
	app := cli.App{
		Name:        "gitstrap",
		Description: "CLI tool to manage GitHub repositories",
		Usage:       "GitHub resource bootstrap",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "token",
				Usage: "GitHub API token with repo access",
			},
		},
		Commands: []*cli.Command{
			getCommand,
			listCommand,
			createCommand,
			deleteCommand,
			applyCommand,
			initCommand,
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func resolveToken(c *cli.Context) (string, error) {
	token := c.String("token")
	if token != "" {
		return token, nil
	}
	file := os.Getenv("HOME") + "/.config/gitstrap/github_token.txt"
	if bin, err := ioutil.ReadFile(file); err == nil {
		return strings.Trim(string(bin), "\n"), nil
	}
	return "", fmt.Errorf("GitHub token neither given as a flag, nor found in %s", file)
}
