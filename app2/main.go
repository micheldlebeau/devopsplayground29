package main

import (
	"fmt"
	"log"
	"net/http"
)

// Hello returns a nice hello world message
func Hello() string {
	return "Hello, world"
}

// Server always prints Hello DevOps Playground!
func Server(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello DevOps Playground!")
}

func main() {
	fmt.Println(Hello())
	handler := http.HandlerFunc(Server)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("could not listen on port 8080 %v", err)
	}

}
