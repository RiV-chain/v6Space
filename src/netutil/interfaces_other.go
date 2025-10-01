//go:build !android

package netutil

import "net"

func Interfaces() ([]net.Interface, error) {
	return net.Interfaces()
}

func InterfaceAddrsByInterface(intf *net.Interface) ([]net.Addr, error) {
	return intf.Addrs()
}
