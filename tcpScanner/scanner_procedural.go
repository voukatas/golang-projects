package main

import (
	"fmt"
	"net"
	"os"
	"sort"
	"strconv"
	"sync"
)

func worker(address string, ports, results chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for p := range ports {
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", address, p))
		if err != nil {
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: scanner <hostname> <start-port> <end-port>")
		os.Exit(1)
	}

	address := os.Args[1]
	startPort, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid start port: ", os.Args[2])
		os.Exit(1)
	}
	endPort, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Invalid end port: ", os.Args[3])
		os.Exit(1)
	}

	if startPort > endPort || startPort < 1 || endPort > 65535 {
		fmt.Println("Invalid port range. It must be within 1 - 65535 and start port should be less than end port.")
		os.Exit(1)
	}

	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int
	var wg sync.WaitGroup

	for i := 0; i < cap(ports); i++ {
		wg.Add(1)
		go worker(address, ports, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	go func() {
		for i := startPort; i <= endPort; i++ {
			ports <- i
		}
		close(ports)
	}()

	for port := range results {
		openports = append(openports, port)
	}

	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
