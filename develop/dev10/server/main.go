package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const address = "127.0.0.1:8080"

func read(conn net.Conn, chanErr chan<- error) {
	incoming := make([]byte, 1024)
	for {
		amount, err := conn.Read(incoming)
		if err != nil {
			chanErr <- err
			close(chanErr)
			return
		}
		fmt.Println(string(incoming[:amount]))
	}
}

func write(conn net.Conn, chanErr chan<- error) {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			chanErr <- err
			close(chanErr)
			return
		}
		_, err = conn.Write([]byte(line))
		if err != nil {
			chanErr <- err
			close(chanErr)
			return
		}
	}
}

func handleConnection(conn net.Conn, chanErr chan<- error) {
	go read(conn, chanErr)
	go write(conn, chanErr)
}

func main() {
	fmt.Println("Starting...")

	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Printf("Host: %s\n", address)

	chanErr := make(chan error, 1)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		select {
		case err := <-chanErr:
			fmt.Println(err)
			os.Exit(1)
		case inter := <-sig:
			fmt.Println("\nServer stopped by signal:", inter)
			os.Exit(0)
		}
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			chanErr <- err
			continue
		}
		fmt.Printf("Connection with %s established\n", conn.RemoteAddr())
		go handleConnection(conn, chanErr)
	}
}
