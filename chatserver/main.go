package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
	mu    sync.Mutex
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("new incoming connection from client:", ws.RemoteAddr())
	s.mu.Lock()
	s.conns[ws] = true
	s.mu.Unlock()

	// in case you have read and write separately, and you use goroutines you need to synchronize the access to the WS connection
	/*
		type SafeWebSocket struct {
			conn *websocket.Conn
			mux  sync.Mutex
		}

	*/
	// go s.readLoop(safeWs)
	// go s.writeLoop(safeWs)

	s.readLoop(ws)

	s.mu.Lock()
	delete(s.conns, ws)
	s.mu.Unlock()
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		//safeWs.mux.Lock()
		//n, err := safeWs.conn.Read(buf)
		//safeWs.mux.Unlock()
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read error:", err)
			continue
		}

		msg := buf[:n]

		s.broadcast(msg)
		fmt.Println(string(msg))
	}

}

// func (s *Server) writeLoop(safeWs *SafeWebSocket) {
// 	for {
// 		safeWs.mux.Lock()
// 		_, err := safeWs.conn.Write([]byte("Thank you for the message!"))
// 		safeWs.mux.Unlock()

// 		if err != nil {
// 			fmt.Println("Write error:", err)
// 			break
// 		}
// 	}
// }

func (s *Server) broadcast(b []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				fmt.Println("error on write:", err)
			}
		}(ws)
	}
}

func main() {
	server := NewServer()
	http.Handle("/ws", websocket.Handler(server.handleWS))
	http.ListenAndServe(":3001", nil)
}
