package tcp

import (
	"fmt"
	"net"
	"testing"
)

func TestServer(t *testing.T) {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(listener.Addr())

	for {
		conn, err := listener.Accept()

		if err != nil {
			t.Fatal(err)
		}
		conn.Write([]byte("Hei\n"))
		fmt.Println("Message sent")
	}

}
