package main

import (
    "fmt"
    "net"
    "time"
    "os"
)

func scanPort(targetIP string, port int, protocol string, results chan int) {
    address := fmt.Sprintf("%s:%d", targetIP, port)

    conn, err := net.DialTimeout(protocol, address, time.Second)

    if err == nil {
        conn.Close()
        results <- port
    } else {
        //fmt.Printf("Port %d is closed\n", port)
        results <- 0
    }
}

func main() {
    start := time.Now()
    targetIP := os.Args[1]
    startPort := 1
    endPort := 65535
    numWorkers := 10

    results := make(chan int)

    for i := 0; i < numWorkers; i++ {
        go func() {
            for port := range results {
                if port != 0 {
                    fmt.Printf("Port %d is open\n", port)
                }
            }
        }()
    }

    for port := startPort; port <= endPort; port++ {
        go scanPort(targetIP, port, "tcp", results)
        //go scanPort(targetIP, port, "udp4", results)
    }

    for i := 0; i < endPort-startPort+1; i++ {
        <-results
        elapsed := time.Since(start)
        fmt.Printf("Port scan took %s", elapsed)
        return
    }
}
