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
		t.Fatal(err)
	}
	fmt.Println(conn)

	message, err := bufio.NewReader(conn).ReadString('\n')

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(message)
}
