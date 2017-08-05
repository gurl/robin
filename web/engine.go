package web

import (
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func StartServer(port string) {
	startHttpServer(port)
}

func startHttpServer(port string) {
	router := NewRouter()
	log.Println("Server Started on port ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
