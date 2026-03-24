package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/hashicorp/yamux"
)

func main() {
	fmt.Println("Starting Local Agent")

	baseDelay := 1 * time.Second
	maxDelay := 30 * time.Second
	currentDelay := baseDelay

	for {
		err := runAgent(&currentDelay, baseDelay)
		if err != nil {
			fmt.Printf("Connection lost or failed:%v\n", err)
			fmt.Printf("Retrying in %v...\n", currentDelay)
			time.Sleep(currentDelay)

			currentDelay *= 2
			if currentDelay > maxDelay {
				currentDelay = maxDelay
			}
		}
	}

}

func runAgent(currentDelay *time.Duration, baseDelay time.Duration) error {
	relayConn, err := net.Dial("tcp", "localhost:7000")
	if err != nil {
		return err
	}
	defer relayConn.Close()

	fmt.Println("Connected to relay")
	*currentDelay = baseDelay

	session, err := yamux.Server(relayConn, nil)
	if err != nil {
		return fmt.Errorf("Yamux initialization failed: %v", err)
	}

	for {
		stream, err := session.Accept()
		if err != nil {
			return fmt.Errorf("Yamux session dropped: %v", err)
		}
		go handleStream(stream)
	}
}

func handleStream(stream net.Conn) {
	defer stream.Close()
	localConn, err := net.Dial("tcp", "localhost:5173")
	if err != nil {
		log.Println("Local app offline:", err)
		return
	}
	defer localConn.Close()

	go io.Copy(localConn, stream)
	io.Copy(stream, localConn)
}
