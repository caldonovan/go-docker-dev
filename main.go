package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	sess, err := CreateSession(os.Getenv("AWS_REGION"))
	if err != nil {
		panic(err)
	}

	sqsSvc := sqs.New(sess)

	queueName := os.Getenv("SQS_QUEUE_NAME")

	// wait for queue to be ready
	for {
		result, err := sqsSvc.GetQueueUrl(&sqs.GetQueueUrlInput{
			QueueName: &queueName,
		})
		if err != nil {
			log.Printf("Error getting Queue URL: %v", err)
		}

		if result != nil {
			fmt.Println("Got Queue URL: " + *result.QueueUrl)
			break
		}
	}

	for {
		result, err := sqsSvc.ReceiveMessage(
			&sqs.ReceiveMessageInput{
				QueueUrl:            aws.String("http://host.docker.internal:4566/000000000000/queue1"),
				MaxNumberOfMessages: aws.Int64(1),
				WaitTimeSeconds:     aws.Int64(3),
			},
		)

		if err != nil {
			fmt.Printf("Failed to receive message with error %v", err)
			time.Sleep(time.Second * 1)
			continue
		}

		if len(result.Messages) == 0 {
			fmt.Printf("No messages found in queue\n")
			time.Sleep(time.Second * 1)
			continue
		}

		fmt.Println(*result.Messages[0].Body)
	}
}

func CreateSession(region string) (*session.Session, error) {
	awsConfig := &aws.Config{
		Region: aws.String(region),
	}

	if localstackEndpoint := os.Getenv("LOCALSTACK_ENDPOINT"); localstackEndpoint != "" {
		awsConfig.Endpoint = aws.String(localstackEndpoint)
	}

	return session.NewSession(awsConfig)
}
