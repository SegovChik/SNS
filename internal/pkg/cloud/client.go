package cloud

import (
	"context"
	)

type PubSubClient interface {
	// Creates a new topic and returns its ARN.
	CreateNewTopic(ctx context.Context, topic string) (string, error)
	// Lists all topics.
	ListTopics(ctx context.Context) ([]*Topic, error)
	// Subscribes a user (e.g. email, phone) to a topic and returns subscription ARN.
	Subscribe(ctx context.Context, endpoint, protocol, topicARN string) (string, error)
	// Lists all subscriptions for a topic.
	ListTopicSubscriptions(ctx context.Context, topicARN string) ([]*Subscription, error)
	// Publishes a message to all subscribers of a topic and returns its message ID.
	Publish(ctx context.Context, message, topicARN string) (string, error)
	// Unsubscribes a topic subscription.
	Unsubscribe(ctx context.Context, subscriptionARN string) error
}