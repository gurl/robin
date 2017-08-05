package engine

import (
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/sqs"

	"github.com/aws/aws-sdk-go/aws/session"
	// "github.com/docker/goamz/dynamodb"
)

var queueURL = os.Getenv("QUEUE_URL")
var sess = session.Must(session.NewSession())
var dynamoSvc = dynamodb.New(sess)
var sqsSvc = sqs.New(sess)
var wg sync.WaitGroup

func startConsole(port string) {

}

func startListener() {

}

// StartRobin is the main controller of the Robin Application
func StartRobin(port string) {
	go startConsole(port)
	wg.Add(1)
	go startListener()
	wg.Add(1)
	wg.Wait()
}
