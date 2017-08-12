package engine

import (
	"log"
	"os"
	"sync"

	"github.com/gurl/robin/listener"
)

var wg sync.WaitGroup
var queueURL = os.Getenv("QUEUE_URL")

func startListener(queueType string) {
	var queueListener listener.Listener
	if queueType == "rmq" {
		queueListener = listener.RMQListener{
			Endpoint:  queueURL,
			Host:      "localhost",
			Port:      "5672",
			QueueName: "robin-queue",
		}
	} else if queueType == "sqs" {
		queueListener = listener.SQSListener{
			Endpoint: queueURL,
		}
	} else {
		log.Fatalf("No listeners available for this time of queue\n")
	}
	listener.StartListening(queueListener)
}

// StartRobin is the main controller of the Robin Application
func StartRobin(queueType string) {
	go startListener(queueType)
	wg.Add(1)
	wg.Wait()
}
