package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func echo(conn net.Conn) {
	defer conn.Close()

	// using bufio you avoid handling the buffer (eg. make([]byte, 1024) and the EOF and other low level staff)
	// the writer is similar to other langs, it outputs the data when the buffer is full or when you explicitly call it

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Unable to read data: %v", err)
			return
		}

		log.Printf("Read %d bytes: %s", len(s), s)

		log.Println("Writing data")
		
		if _, err := writer.WriteString(s); err != nil {
			log.Printf("Unable to write data: %v", err)
			return
		}

		writer.Flush()
	}
}

func main() {
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Failed to bind to port: %v", err)
	}

	log.Println("Socket created, binded, and is ready listening on 0.0.0.0:5000!")

	// Flag to signal when it's time to shut down
	// Common pattern in go, an empty struct dosen't take up space while an int will	
	shutdown := make(chan struct{})

	// Graceful shutdown
	sigCh := make(chan os.Signal, 1) //It's a common idiom in Go when dealing with os signals
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh // // value struct{}{}
		log.Println("Graceful shutdown initiated")
		close(shutdown) // Signal that it's time to shut down
		listener.Close()
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			select {
			case <-shutdown:
				log.Println("Listener closed, shutting down")
				return // Exit main shutdown signaled
			default:
				log.Printf("Failed to accept connection: %v", err)
				continue
			}
		}

		go echo(conn)
	}
}

