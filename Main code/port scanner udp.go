package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// Set timeout for the connection
	timeout := time.Second

	// Set target IP address
	targetIP := net.ParseIP("31.201.56.59")

	// Loop through ports 1-65535 and try to connect via UDP
	for port := 1; port <= 65535; port++ {
		conn, err := net.DialTimeout("udp", fmt.Sprintf("%s:%d", targetIP.String(), port), timeout)
		if err == nil {
			fmt.Printf("Port %d is open\n", port)
			conn.Close()
		}
	}
}
