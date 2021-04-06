package aws

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/cloud"
)

var _ cloud.PubSubClient = SNS{}

type SNS struct {
	timeout time.Duration
	client  *sns.SNS
}

func (s SNS) CreateNewTopic(ctx context.Context, topic string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	res, err := s.client.CreateTopicWithContext(ctx, &sns.CreateTopicInput{
		Name: aws.String(topic),
	})
	if err != nil {
		return "", fmt.Errorf("create: %w", err)
	}

	return *res.TopicArn, nil
}

func NewSNS(session *session.Session, timeout time.Duration) SNS {
	return SNS{
		timeout: timeout,
		client:  sns.New(session),
	}
}


func (s SNS) ListTopics(ctx context.Context) ([]*cloud.Topic, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	res, err := s.client.ListTopicsWithContext(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("list topics: %w", err)
	}

	topics := make([]*cloud.Topic, len(res.Topics))

	for i, topic := range res.Topics {
		topics[i] = &cloud.Topic{
			ARN: *topic.TopicArn,
		}
	}

	return topics, nil
}

func (s SNS) Subscribe(ctx context.Context, endpoint, protocol, topicARN string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	res, err := s.client.SubscribeWithContext(ctx, &sns.SubscribeInput{
		Endpoint:              aws.String(endpoint),
		Protocol:              aws.String(protocol),
		ReturnSubscriptionArn: aws.Bool(true),
		TopicArn:              aws.String(topicARN),
	})
	if err != nil {
		return "", fmt.Errorf("subscribe: %w", err)
	}

	return *res.SubscriptionArn, nil
}

func (s SNS) ListTopicSubscriptions(ctx context.Context, topicARN string) ([]*cloud.Subscription, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	res, err := s.client.ListSubscriptionsByTopicWithContext(ctx, &sns.ListSubscriptionsByTopicInput{
		NextToken: nil,
		TopicArn:  aws.String(topicARN),
	})
	if err != nil {
		return nil, fmt.Errorf("list topic subscriptions: %w", err)
	}

	subs := make([]*cloud.Subscription, len(res.Subscriptions))

	for i, sub := range res.Subscriptions {
		subs[i] = &cloud.Subscription{
			ARN:      *sub.SubscriptionArn,
			TopicARN: *sub.TopicArn,
			Endpoint: *sub.Endpoint,
			Protocol: *sub.Protocol,
		}
	}

	return subs, nil
}

func (s SNS) Publish(ctx context.Context, message, topicARN string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	res, err := s.client.PublishWithContext(ctx, &sns.PublishInput{
		Message:  &message,
		TopicArn: aws.String(topicARN),
	})
	if err != nil {
		return "", fmt.Errorf("publish: %w", err)
	}

	return *res.MessageId, nil
}

func (s SNS) Unsubscribe(ctx context.Context, subscriptionARN string) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	if _, err := s.client.UnsubscribeWithContext(ctx, &sns.UnsubscribeInput{
		SubscriptionArn: aws.String(subscriptionARN),
	}); err != nil {
		return fmt.Errorf("unsubscribe: %w", err)
	}

	return nil
}