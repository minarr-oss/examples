package main

import (
	"net/http"

	"github.com/zenohq/example-apps/handlers"

	"github.com/sirupsen/logrus"
)

func main() {
	http.HandleFunc("/", handlers.Howdy)
	logrus.Println("Yep, I'm ready~ :8080")
	http.ListenAndServe(":8080", nil)
}
