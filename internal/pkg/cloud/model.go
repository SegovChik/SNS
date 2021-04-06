package cloud

type Topic struct {
	ARN string
}

type Subscription struct {
	ARN      string
	TopicARN string
	Endpoint string
	Protocol string
}