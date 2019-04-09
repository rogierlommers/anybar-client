package anybar

import (
	"net"
)

const hostPort = "localhost:1738"

func Send(msg string) error {
	laddr, err := net.ResolveUDPAddr("udp", ":0")
	if err != nil {
		return err
	}

	maddr, err := net.ResolveUDPAddr("udp4", hostPort)
	if err != nil {
		return err
	}

	c, err := net.ListenUDP("udp4", laddr)
	if err != nil {
		return err
	}
	defer c.Close()

	_, err = c.WriteToUDP([]byte(msg), maddr)
	if err != nil {
		return err
	}

	return nil
}
