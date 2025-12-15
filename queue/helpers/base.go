package helpers

import (
	"net"
)

// GetOutboundIP returns the preferred outbound IP of this machine.
// Works reliably on macOS, Linux, Docker, WSL.
func GetOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}
