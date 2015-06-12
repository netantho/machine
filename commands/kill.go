package commands

import (
	"github.com/netantho/machine/log"

	"github.com/codegangsta/cli"
)

func cmdKill(c *cli.Context) {
	if err := runActionWithContext("kill", c); err != nil {
		log.Fatal(err)
	}
}
