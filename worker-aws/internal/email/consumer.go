package email

import (
	"context"
	"errors"
	"log"
	"sync"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/praiakov/worker-aws/internal/pkg/cloud"
)

type ConsumerType string

const (
	// Consumers to consume messages one by one.
	// A single goroutine handles all messages.
	// Progression is slower and requires less system resource.
	// Ideal for quiet/non-critical queues.
	SyncConsumer ConsumerType = "blocking"
	// Consumers to consume messages at the same time.
	// Runs an individual goroutine per message.
	// Progression is faster and requires more system resource.
	// Ideal for busy/critical queues.
	AsyncConsumer ConsumerType = "non-blocking"
)

type ConsumerConfig struct {
	// Instructs whether to consume messages come from a worker synchronously or asynchronous.
	Type ConsumerType
	// Queue URL to receive messages from.
	QueueURL string
	// Maximum workers that will independently receive messages from a queue.
	MaxWorker int
	// Maximum messages that will be picked up by a worker in one-go.
	MaxMsg int
}

type Consumer struct {
	client cloud.MessageClient
	config ConsumerConfig
}

func NewConsumer(client cloud.MessageClient, config ConsumerConfig) Consumer {
	return Consumer{
		client: client,
		config: config,
	}
}

func (c Consumer) Start(ctx context.Context) string {
	wg := &sync.WaitGroup{}
	wg.Add(c.config.MaxWorker)

	for i := 1; i <= c.config.MaxWorker; i++ {
		ok, err := c.worker(ctx, wg, i)

		if err != nil {
			return "creu"
		}

		return ok
	}

	wg.Wait()

	return "ok"

}

func (c Consumer) worker(ctx context.Context, wg *sync.WaitGroup, id int) (string, error) {
	defer wg.Done()

	log.Printf("worker %d: started\n", id)

	for {
		select {
		case <-ctx.Done():
			log.Printf("worker %d: stopped\n", id)
			return "", errors.New("failed")
		default:
		}

		msgs, err := c.client.Receive(ctx, c.config.QueueURL, int64(c.config.MaxMsg))
		if err != nil {
			// Critical error!
			log.Printf("worker %d: receive error: %s\n", id, err.Error())
			return "", errors.New("failed receive message")
		}

		if len(msgs) == 0 {
			return "ok", nil
		}

		if c.config.Type == SyncConsumer {
			c.sync(ctx, msgs)
		} else {
			c.async(ctx, msgs)
		}
	}
}

func (c Consumer) sync(ctx context.Context, msgs []*sqs.Message) {
	for _, msg := range msgs {
		c.consume(ctx, msg)
	}
}

func (c Consumer) async(ctx context.Context, msgs []*sqs.Message) {
	wg := &sync.WaitGroup{}
	wg.Add(len(msgs))

	for _, msg := range msgs {
		go func(msg *sqs.Message) {
			defer wg.Done()

			c.consume(ctx, msg)
		}(msg)
	}

	wg.Wait()
}

func (c Consumer) consume(ctx context.Context, msg *sqs.Message) {
	log.Println(*msg.Body)

	if err := c.client.Delete(ctx, c.config.QueueURL, *msg.ReceiptHandle); err != nil {
		// Critical error!
		log.Printf("delete error: %s\n", err.Error())
	}
}
