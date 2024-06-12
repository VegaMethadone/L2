package functions

import (
	"fmt"
	"net"
)

func read(connection net.Conn, ch chan<- error) {
	input := make([]byte, 1024)
	for {
		for {
			n, err := connection.Read(input)
			if err != nil {
				ch <- fmt.Errorf("error server: %v", err)
				close(ch)
				return
			}
			fmt.Println(string(input[:n]))
		}
	}
}
