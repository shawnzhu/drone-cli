package main

import (
	"fmt"
	"io/ioutil"

	"github.com/codegangsta/cli"
	"github.com/drone/drone-go/drone"
)

// NewEncryptCommand returns the CLI command for "encrypt".
func NewEncryptCommand() cli.Command {
	return cli.Command{
		Name:  "encrypt",
		Usage: "Encrypts parameter with public key of repository",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) {
			handle(c, EncryptCommandFunc)
		},
	}
}

// encryptCommandFunc executes the "encrypt" command.
func encryptCommandFunc(c *cli.Context, client *drone.Client) error {
	var host, owner, name, entry string
	var args = c.Args()

	if len(args) == 2 {
		host, owner, name = parseRepo(args[0])
	} else {
		return fmt.Errorf("Please specify repo string and k=v")
	}

	repo, err := client.Repos.Get(host, owner, name)

	if err != nil {
		return fmt.Errorf("Could not find public RSA key for %s. %s", args[0], err)
	} else if repo.PublicKey == "" {
		return fmt.Errorf("public RSA key for %s is empty. please run set-key command first", args[0])
	}

	fmt.Sprintf("- Secure: %s", entry)
}
