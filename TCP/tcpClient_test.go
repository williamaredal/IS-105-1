package tcp

import (
	"bufio"
	"fmt"
	"net"
	"testing"
)

func TestClient(t *testing.T) {
	conn, err := net.Dial("tcp", ":8080")

	if err != nil {
		panic("Dialer failed!")
	}
	fmt.Println(conn)

	message, err := bufio.NewReader(conn).ReadString('\n')

	if err != nil {
		fmt.Println("Message error!")
	}

	fmt.Println(message)
}
