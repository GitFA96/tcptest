package main

import (
	"crypto/rand"
	"fmt"
	"net"
)

func main() {
	payload := make([]byte, 1<<24) // allokerer plass i minnet - hvor mye og hvilken type --> intern MiB ikke MB
	_, err := rand.Read(payload)   // leser inn tilfeldige bytes i payload
	if err != nil {
		fmt.Println(err)

	}

	listener, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println(err)
	}
	// for å teste evt vente i golang rutine.

	fmt.Println("Hei, jeg venter!")

	// go func() {
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()

		_, err = conn.Write(payload)
		if err != nil {
			fmt.Println(err)
		}
	}
	//}()

}

// for å kjøre : go run read_server.go
// eller go build
