package listener

import (
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
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

	for {
		time.Sleep(2 * time.Second)
		log.Printf("Listening on the Endpoint %v\n", l.Endpoint)
		result, err := sqsSvc.ReceiveMessage(&sqs.ReceiveMessageInput{
			AttributeNames: []*string{
				aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
			},
			MessageAttributeNames: []*string{
				aws.String(sqs.QueueAttributeNameAll),
			},
			QueueUrl:            &l.Endpoint,
			MaxNumberOfMessages: aws.Int64(1),
			VisibilityTimeout:   aws.Int64(0),
			WaitTimeSeconds:     aws.Int64(0),
		})

		if err != nil {
			log.Printf("Error: %v\n", err)
			return
		}

		if len(result.Messages) == 0 {
			log.Printf("Received no messages\n")
			return
		}
		resultDelete, err := sqsSvc.DeleteMessage(&sqs.DeleteMessageInput{
			QueueUrl:      &l.Endpoint,
			ReceiptHandle: result.Messages[0].ReceiptHandle,
		})

		if err != nil {
			log.Printf("Delete Error: %v\n", err)
			return
		}

		log.Printf("Message Deleted %v\n", resultDelete)
	}
}
