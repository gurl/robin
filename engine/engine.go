package engine

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/alwindoss/robin/listener"
	// "github.com/docker/goamz/dynamodb"
)

var wg sync.WaitGroup
var queueURL = os.Getenv("QUEUE_URL")

func startConsole(port string) {
	for {
		time.Sleep(1000 * time.Millisecond)
		log.Println("In the console")
	}
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
