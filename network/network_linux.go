package network

import (
	"github.com/brodyxchen/aws-enclave-kms/log"
	"github.com/milosgajdos/tenus"
	"net"
)

// assignLoAddr assigns an IP address to the loopback interface, which is
// necessary because Nitro enclaves don't do that out-of-the-box.  We need the
// loopback interface because we run a simple TCP proxy that listens on
// 127.0.0.1:1080 and converts AF_INET to AF_VSOCK.
func AssignLoAddr() error {
	log.Info("AssignLoAddr() in linux")
	addrStr := "127.0.0.1/8"
	l, err := tenus.NewLinkFrom("lo")
	if err != nil {
		return err
	}
	addr, network, err := net.ParseCIDR(addrStr)
	if err != nil {
		return err
	}
	if err = l.SetLinkIp(addr, network); err != nil {
		return err
	}
	return l.SetLinkUp()
}
