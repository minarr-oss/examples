package main

import (
	"net/http"
	"fmt"

	"github.com/zenohq/example-apps/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Howdy)
	fmt.Println("Yep, I'm ready~ :8080")
	http.ListenAndServe(":8080", nil)
}
