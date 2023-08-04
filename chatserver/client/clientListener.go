package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/websocket"
)

func main() {
	origin := "http://localhost/"
	url := "ws://localhost:3001/ws"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		fmt.Println("Dial:", err)
		os.Exit(1)
	}

	message := []byte("hello from client")
	_, err = ws.Write(message)
	if err != nil {
		fmt.Println("Write:", err)
		os.Exit(1)
	}
	fmt.Println("Sent to server: ", string(message))

	received := make([]byte, 256)
	for {
		n, err := ws.Read(received)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Read: ", err)
			continue
		}
		fmt.Println("Received from server: ", string(received[:n]))
	}
}
