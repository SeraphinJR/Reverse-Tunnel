package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/hashicorp/yamux"
)

func main() {
	tunnelListener, err := net.Listen("tcp", ":7000")
	if err != nil {
		log.Fatal(err)
	}

	agentConn, err := tunnelListener.Accept()
	if err != nil {
		log.Fatalf("Failed to bind to port 7000: %v", err)
	}
	fmt.Println("Agent connected...")

	session, err := yamux.Client(agentConn, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Waiting for traffic on http://localhost:8000")
	publicListener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		browserConn, err := publicListener.Accept()
		if err != nil {
			continue
		}
		go handleBrowser(browserConn, session)
	}
}

func handleBrowser(browserConn net.Conn, session *yamux.Session) {
	defer browserConn.Close()
	stream, err := session.Open()
	if err != nil {
		log.Println("Failed to open multiplexed stream:", err)
		return
	}
	defer stream.Close()

	go io.Copy(stream, browserConn)
	io.Copy(browserConn, stream)
}
