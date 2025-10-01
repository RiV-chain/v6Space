package netutil

import (
	"net"

	"github.com/wlynxg/anet"
)

func Interfaces() ([]net.Interface, error) {
	return anet.Interfaces()
}

func InterfaceAddrsByInterface(intf *net.Interface) ([]net.Addr, error) {
	return anet.InterfaceAddrsByInterface(intf)
}
