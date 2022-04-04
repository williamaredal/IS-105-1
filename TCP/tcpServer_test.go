package tcp

import (
	"fmt"
	"net"
	"testing"
)

func TestServer(t *testing.T) {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic("Listener error!")
	}
	fmt.Println(listener.Addr())

	for {
		conn, err := listener.Accept()

		if err != nil {
			panic("Connection error!")
		}
		conn.Write([]byte("Hei\n"))
		fmt.Println("Message sent")
	}

}
