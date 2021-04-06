package main

import (
	"context"
	"github.com/aws"
	"github.com/pubsub"
	"log"
	"time"
)

func main() {
	// Create a session instance.
	ses, err := aws.New(aws.Config{
		Region:  "eu-central-1",
		Profile: "default",
		ID:      "secret",
		Secret:  "secret",
	})
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()

	var snsResource = aws.NewSNS(ses, time.Second*5)
	topicARN := pubsub.CreateNewTopic(ctx, snsResource, "newtopic")
	pubsub.ListTopics(ctx , snsResource)
	pubsub.Subscribe(ctx, snsResource, topicARN , "+380665065480", "sms")
	pubsub.Publish(ctx, snsResource, topicARN, "texttosend")
	pubsub.ListTopicSubscriptions(ctx, snsResource, topicARN)
}