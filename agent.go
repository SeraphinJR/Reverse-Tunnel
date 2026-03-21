package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	fmt.Println("Starting Local Agent")
	conn, err := net.Dial("tcp", "localhost:7000")
	if err != nil {
		log.Fatalf("Could not connect to Relay: %v", err)
	}

	defer conn.Close()

	println("Connected successfully to Relay")

	go func() {
		networkScanner := bufio.NewScanner(conn)
		for networkScanner.Scan() {
			fmt.Println("->", networkScanner.Text())
		}
		fmt.Println("\nRelay disconnected. Ctrl+C to exit")
		os.Exit(0)
	}()

	keyboardScanner := bufio.NewScanner(os.Stdin)
	for keyboardScanner.Scan() {
		text := keyboardScanner.Text()
		_, err := conn.Write([]byte(text + "\n"))
		if err != nil {
			log.Fatalf("Failed to send message: %v", err)
		}
	}
}
