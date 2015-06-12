package ssh

import (
	"net"
	"time"

	"github.com/netantho/machine/log"
)

func WaitForTCP(addr string) error {
	for {
		log.Debugf("Testing TCP connection to: %s", addr)
		conn, err := net.DialTimeout("tcp", addr, 2*time.Second)

		if err != nil {
			continue
		}

		defer conn.Close()
		return nil
	}
}
