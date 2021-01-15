package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kembo91/linxdatacenter-test-task/handlers"
)

func main() {
	http.HandleFunc("/", handlers.EmailHandler)
	port := os.Getenv("LISTEN_PORT")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
