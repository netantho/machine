package commands

import (
	"fmt"

	"github.com/netantho/machine/log"

	"github.com/codegangsta/cli"
)

func cmdUrl(c *cli.Context) {
	url, err := getHost(c).GetURL()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(url)
}
