package commands

import (
	"github.com/codegangsta/cli"
	"github.com/netantho/machine/log"
)

func cmdIp(c *cli.Context) {
	if err := runActionWithContext("ip", c); err != nil {
		log.Fatal(err)
	}
}
