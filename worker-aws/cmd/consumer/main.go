package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/praiakov/worker-aws/internal/email"
	"github.com/praiakov/worker-aws/internal/pkg/cloud/aws"
)

func main() {
	http.ListenAndServe(":8010", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// This prints to STDOUT to show that processing has started
		fmt.Fprint(os.Stdout, "processing request\n")
		// We use `select` to execute a peice of code depending on which
		// channel receives a message first
		select {
		case <-time.After(0 * time.Second):
			// If we receive a message after 2 seconds
			// that means the request has been processed
			// Create a cancellable context.
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			// Create a session instance.
			ses, err := aws.New(aws.Config{
				Address: "http://localhost:4566",
				Region:  "eu-west-1",
				Profile: "localstack",
				ID:      "test",
				Secret:  "test",
			})
			if err != nil {
				log.Fatalln(err)
			}

			// Set queue URL.
			url := "http://localhost:4566/000000000000/welcome-email"

			// Instantiate client.
			client := aws.NewSQS(ses, time.Second*5)

			// Instantiate consumer and start consuming.
			a := email.NewConsumer(client, email.ConsumerConfig{
				Type:      email.AsyncConsumer,
				QueueURL:  url,
				MaxWorker: 1,
				MaxMsg:    10,
			}).Start(ctx)

			if a == "ok" {
				w.Write([]byte("request processed"))
			} else {
				w.Write([]byte("sqs failed"))
			}

		case <-ctx.Done():
			// If the request gets cancelled, log it
			// to STDERR
			fmt.Fprint(os.Stderr, "request cancelled\n")
		}
	}))

}
