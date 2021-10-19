package email

import (
	"context"
	"fmt"
	"log"

	"github.com/praiakov/worker-aws/internal/pkg/cloud"
)

type Producer struct {
	client cloud.MessageClient
}

func NewProducer(client cloud.MessageClient) Producer {
	return Producer{client: client}
}

func (p Producer) Produce(ctx context.Context, queueURL string) error {
	for i := 1; i <= 2200; i++ {
		if _, err := p.client.Send(ctx, &cloud.SendRequest{
			QueueURL: queueURL,
			Body:     fmt.Sprintf("Message : %d", i),
			Attributes: []cloud.Attribute{
				{
					Key:   "FromEmail",
					Value: "from@example.com",
					Type:  "String",
				},
				{
					Key:   "ToEmail",
					Value: "to@example.com",
					Type:  "String",
				},
				{
					Key:   "Subject",
					Value: "this is subject",
					Type:  "String",
				},
			},
		}); err != nil {
			return err
		}

		log.Printf("body: %d\n", i)
	}

	return nil
}
