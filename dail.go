package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080") // _ = conn og uavklarte metodikk

	if err != nil {
		fmt.Println(err)
	}

	defer conn.close()

	fmt.Fprintf(conn, "hei\n")
	echo, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(echo)
}
