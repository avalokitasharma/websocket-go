package main

import (
	"log"
	"net/http"

	"github.com/avalokitasharma/websocket-go/internal/handlers"
)

const port = ":8080"

func main() {

	mux := routes()

	log.Println("Starting websocket:")
	go handlers.ListenToWsChannel()

	log.Println("Starting application on port:", port)
	_ = http.ListenAndServe(port, mux)
}
