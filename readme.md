## Основные сущности AWS SNS:
- Topic - это коммуникационный канал для отправки сообщений и подписки на уведомления. 
Для создания/изменения нам нужен только его ARN. 
```
type Topic struct {
	ARN string 
 }
```

- Subscription - подписка на тему. Состоит из своего ARN'a, ARN'a Topic'a c которым связана, поля Endpoint(мыло/номер юзера) и поля Protool("email"/"sms")
  Момент который меня смущает - юзер на email которого создаётся Subscription получает письмо в котором он должен дать согласие на подписку, если он этого не делает то все 
  уведомления которые попадают в тему к нему не доходят:(
```
type Subscription struct {
	ARN      string
	TopicARN string
	Endpoint string
	Protocol string
}
```
- Publish - публицация сообщения в Topic. После публикации cообщения оно отправляется на все Subscription, который связаны с темой.

##Устройство интерфейса:
 ```
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
```
## Последовательность действий для отправки пользователю сообщения(много деталей пропущено, полноценный пример использования в main.go)
 1. Создать Topic
 ```
 topicARN := pubsub.CreateNewTopic(ctx, snsResource, "topicname")
 ```
 2. Создать Subscription передав ARN связанной темы, Endpoint(мыло/номер юзера) Protool("email"/"sms")
```
subARN := pubsub.Subscribe(ctx, snsResource, topicARN , "+380665065480", "sms")
```
3. Опубликовать сообщение в Тему на которую создана подписка с номером/мылом
```
pubsub.Publish(ctx, snsResource, topicARN, "texttosend")
```
