package main

import (
	"fmt"
	"net"
	"os"
	"sort"
	"strconv"
	"sync"
)

type Scanner struct {
	address  string
	ports    chan int
	results  chan int
	openports []int
}

func (s *Scanner) worker(wg *sync.WaitGroup) {
	defer wg.Done()

	for p := range s.ports {
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", s.address, p))
		if err != nil {
			continue
		}
		conn.Close()
		s.results <- p
	}
}

func (s *Scanner) scan(startPort, endPort int) {
	var wg sync.WaitGroup

	for i := 0; i < cap(s.ports); i++ {
		wg.Add(1)
		go s.worker(&wg)
	}

	go func() {
		wg.Wait()
		close(s.results)
	}()

	go func() {
		for i := startPort; i <= endPort; i++ {
			s.ports <- i
		}
		close(s.ports)
	}()

	for port := range s.results {
		s.openports = append(s.openports, port)
	}

	sort.Ints(s.openports)
}

func (s *Scanner) printResults() {
	for _, port := range s.openports {
		fmt.Printf("%d open\n", port)
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

	scanner := Scanner{
		address: address,
		ports:   make(chan int, 100),
		results: make(chan int),
	}

	scanner.scan(startPort, endPort)
	scanner.printResults()
}
