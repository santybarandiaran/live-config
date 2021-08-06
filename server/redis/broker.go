package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"live-config/logger"
	"live-config/server/domain"
	"live-config/server/redis/topic"
)

type MessageBroker struct {
	Redis *redis.Client
}

func (b *MessageBroker) PublishMessage(p *domain.Property) {
	//TODO I don't want the property itself, just key value
	t := topic.GetTopicTemplate(p)
	m := Redis.Publish(ctx, t, p)

	err := m.Err()

	if err != nil {
		logger.Instance.Error(fmt.Sprintf("Error trying to publish update message to subscribers of topic %s: %+v.  Error: %s", t, p, err.Error()))
		return
	}
}

func (b *MessageBroker) Subscribe(application string, profile string, label string) *redis.PubSub {
	t := topic.GetTopic(application, profile, label)
	return b.Redis.Subscribe(ctx, t)
}
