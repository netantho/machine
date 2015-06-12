package commands

import (
	"github.com/netantho/machine/log"

	"github.com/codegangsta/cli"
)

func cmdStop(c *cli.Context) {
	if err := runActionWithContext("stop", c); err != nil {
		log.Fatal(err)
	}
}
