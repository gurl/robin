package listener

import (
	"log"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/sqs"

	"github.com/aws/aws-sdk-go/aws/session"
)

var sess = session.Must(session.NewSession())
var dynamoSvc = dynamodb.New(sess)
var sqsSvc = sqs.New(sess)

// SQSListener is the SQS implementation of the Listener interface
type SQSListener struct {
	Endpoint string
}

// Listen is an implementation of the Listen method of the Listener interface
func (l SQSListener) Listen() {
	// QUEUE_URL is the Endpoint
	log.Printf("Listening on the Endpoint %v\n", l.Endpoint)
	for {
		time.Sleep(1000 * time.Millisecond)
		log.Println("Listening...")
	}
}
