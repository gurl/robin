package engine

import (
	"log"
	"os"
	"sync"

	"github.com/alwindoss/robin/listener"
	"github.com/alwindoss/robin/web"
	// "github.com/docker/goamz/dynamodb"
)

var wg sync.WaitGroup
var queueURL = os.Getenv("QUEUE_URL")

func startConsole(port string) {
	// time.Sleep(1000 * time.Millisecond)
	log.Println("Starting the Console")
	web.StartServer(port)
}

func startListener() {
	sqsListener := listener.SQSListener{
		Endpoint: queueURL,
	}
	listener.StartListening(sqsListener)
}

// StartRobin is the main controller of the Robin Application
func StartRobin(port string) {
	go startConsole(port)
	wg.Add(1)
	go startListener()
	wg.Add(1)
	wg.Wait()
}
