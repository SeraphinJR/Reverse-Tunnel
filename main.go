package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":7000")
	if err != nil {
		log.Fatalf("Failed to bind to port 7000: %v", err)
	}

	defer listener.Close()

	println("Listening on port 7000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to connect: %v", err)
		}
		fmt.Printf("Connection accepted from %s\n", conn.RemoteAddr())

		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte("Hello hi hellooooo\n"))
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Printf("Agent said: %s\n", text)
		echoMessage := fmt.Sprintf("Relay echo: %s\n", text)
		conn.Write([]byte(echoMessage))
	}
	fmt.Println("Disconnected")
}
