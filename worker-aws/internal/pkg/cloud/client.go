package cloud

import (
	"context"

	"github.com/aws/aws-sdk-go/service/sqs"
)

type MessageClient interface {
	// Send a message to queue and returns its message ID.
	Send(ctx context.Context, req *SendRequest) (string, error)
	// Long polls given amount of messages from a queue.
	Receive(ctx context.Context, queueURL string, maxMsg int64) ([]*sqs.Message, error)
	// Deletes a message from a queue.
	Delete(ctx context.Context, queueURL, rcvHandle string) error
}
