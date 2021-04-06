package pubsub

import (
	"context"
	"fmt"
	"log"

	"github.com/cloud"
)

/*func PubSub(client cloud.PubSubClient) {
	ctx := context.Background()

	tARN := create(ctx, client)
	listTopics(ctx, client)
	sARN := subscribe(ctx, client, tARN)
	listTopicSubscriptions(ctx, client, tARN)
	publish(ctx, client, tARN)
	unsubscribe(ctx, client, sARN)
}*/

func CreateNewTopic(ctx context.Context, client cloud.PubSubClient, myTopic string) string {
	arn, err := client.CreateNewTopic(ctx, myTopic)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("create: topic ARN:", arn)

	return arn
}

func ListTopics(ctx context.Context, client cloud.PubSubClient) {
	topics, err := client.ListTopics(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("list topics:")
	for _, topic := range topics {
		fmt.Printf("%+v\n", topic)
	}
}

func Subscribe(ctx context.Context, client cloud.PubSubClient, topicARN string, endpoint string, protocol string) string {
	arn, err := client.Subscribe(ctx, endpoint, protocol, topicARN)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("subscribe: subscription ARN:", arn)

	return arn
}


func ListTopicSubscriptions(ctx context.Context, client cloud.PubSubClient, topicARN string) {
	subs, err := client.ListTopicSubscriptions(ctx, topicARN)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("list topic subscriptions:")
	for _, sub := range subs {
		fmt.Printf("%+v\n", sub)
	}
}

func Publish(ctx context.Context, client cloud.PubSubClient, topicARN string, message string) {
	id, err := client.Publish(ctx, message, topicARN)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("publish: message ID:", id)
}

func Unsubscribe(ctx context.Context, client cloud.PubSubClient, subARN string) {
	if err := client.Unsubscribe(ctx, subARN); err != nil {
		log.Fatalln(err)
	}
	log.Println("unsubscribe: ok")
}
