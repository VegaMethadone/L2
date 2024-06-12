package functions

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"telnet/structs"
)

func Telnet(args *structs.Args) error {
	address := fmt.Sprintf("%s:%s", args.Host, args.Port)

	connection, err := net.DialTimeout("tcp", address, args.Timeout)
	if err != nil {
		return err
	}
	fmt.Println("Connected to:", address)

	triger := make(chan os.Signal, 1)
	signal.Notify(triger, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	done := make(chan error, 1)

	go read(connection, done)
	go write(connection, done)

	select {
	case value := <-triger:
		fmt.Println("Connection is interrupted:", value)
	case value := <-done:
		fmt.Println("Connections error:", value)
	}

	connection.Close()

	return nil
}
