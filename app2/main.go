package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
)

// Server always prints Hello DevOps Playground!
func Server(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello DevOps Playground!")
}

func main() {
	http.HandleFunc("/", Server)
	appengine.Main()
}
