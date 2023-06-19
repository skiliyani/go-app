package messaging

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// SQSClient represents an AWS SQS client
type SQSClient struct {
	client *sqs.SQS
}

// NewSQSClient creates a new instance of SQSClient
func NewSQSClient(region string) *SQSClient {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))

	return &SQSClient{
		client: sqs.New(sess),
	}
}

// Queue represents an SQS queue
type Queue struct {
	url    string
	client *SQSClient
}

// NewQueue creates a new instance of Queue
func NewQueue(url string, client *SQSClient) *Queue {
	return &Queue{
		url:    url,
		client: client,
	}
}

// StartListening starts listening for incoming messages from the SQS queue
func (q *Queue) StartListening(handler func(string) error) error {
	// Implement your logic to start listening to the SQS queue
	// Call the handler function for each received message

	return nil
}
