package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"live-config/logger"
	"live-config/server/domain"
)

type MessageBroker struct {
	Redis *redis.Client
}

func getTopicTemplate(p *domain.Property) string {
	//TODO add support to enable custom topic creation
	return fmt.Sprintf(defaultTopicTemplate, p.Application, p.Profile, p.Label)
}

func (b *MessageBroker) PublishMessage(p *domain.Property) {
	t := getTopicTemplate(p)
	m := Redis.Publish(ctx, t, p)

	err := m.Err()

	if err != nil {
		logger.Instance.Error(fmt.Sprintf("Error trying to publish update message to subscribers of topic %s: %+v", t, p))
		return
	}
}
