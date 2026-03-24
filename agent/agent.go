package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/hashicorp/yamux"
)

func main() {
	fmt.Println("Starting Local Agent")
	relayConn, err := net.Dial("tcp", "localhost:7000")
	if err != nil {
		log.Fatalf("Could not connect to Relay: %v", err)
	}
	fmt.Println("Connected to relay")

	session, err := yamux.Server(relayConn, nil)
	if err != nil {
		log.Fatal(err)
	}

	for {
		stream, err := session.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleStream(stream)
	}
}

func handleStream(stream net.Conn) {
	defer stream.Close()
	localConn, err := net.Dial("tcp", "localhost:5173")
	if err != nil {
		log.Fatal(err)
	}
	defer localConn.Close()

	go io.Copy(localConn, stream)
	io.Copy(stream, localConn)
}
