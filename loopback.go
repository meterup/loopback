package loopback

import (
	"net"
)

func Loopback() ([]*net.IPAddr, error) {
	addr6, err6 := Loopback6()
	if err6 != nil {
		return nil, err6
	}
	addr4, err4 := Loopback4()
	if err4 != nil {
		return nil, err4
	}
	return []*net.IPAddr{addr4, addr6}, nil
}

func Loopback4() (*net.IPAddr, error) {
	return net.ResolveIPAddr("ip4", "localhost")
}

func Loopback6() (*net.IPAddr, error) {
	return net.ResolveIPAddr("ip6", "localhost")
}
