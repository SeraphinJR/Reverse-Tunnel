package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	fmt.Println("Starting Local Agent")
	relayConn, err := net.Dial("tcp", "localhost:7000")
	if err != nil {
		log.Fatalf("Could not connect to Relay: %v", err)
	}
	fmt.Println("Connected to relay")

	localConn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Local app connected")

	go io.Copy(localConn, relayConn)

	io.Copy(relayConn, localConn)
}
