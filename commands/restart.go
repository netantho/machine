package commands

import (
	"github.com/netantho/machine/log"

	"github.com/codegangsta/cli"
)

func cmdRestart(c *cli.Context) {
	if err := runActionWithContext("restart", c); err != nil {
		log.Fatal(err)
	}
}
