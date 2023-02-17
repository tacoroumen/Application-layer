package main

import (
    "fmt"
    "net"
    //"sort"
    "time"
)

func scanPort(targetIP string, port int, results chan int) {
    address := fmt.Sprintf("%s:%d", targetIP, port)

    conn, err := net.DialTimeout("tcp4", address, time.Second)

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
    targetIP := "31.201.56.59"
    startPort := 1
    endPort := 65535
    numWorkers := 5

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
        go scanPort(targetIP, port, results)
    }

    for i := 0; i < endPort-startPort+1; i++ {
        <-results
        elapsed := time.Since(start)
        fmt.Printf("Port scan took %s", elapsed)
        return
    }
}
