package functions

import (
	"bufio"
	"net"
	"os"
)

func write(connection net.Conn, ch chan<- error) {
	reader := bufio.NewReader(os.Stdin)
	for {
		outcoming, err := reader.ReadBytes('\n')
		if err != nil {
			ch <- err
			close(ch)
			return
		}
		_, err = connection.Write(outcoming)
		if err != nil {
			ch <- err
			close(ch)
			return
		}

	}
}
