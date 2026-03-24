package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received request: %s %s\n", r.Method, r.URL.Path)
		fmt.Fprintf(w, "<h1>YAYYYYYY!!!!</h1></p>WE went through the tunnel yippee</p>")
	})
	fmt.Println("Local app on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
