package topic

import (
	"fmt"
	"live-config/server/domain"
)

const defaultTopicTemplate = "live-config_%s_%s_%s"

// TODO add support to enable custom topic creation

func GetTopicTemplate(p *domain.Property) string {
	return GetTopic(p.Application, p.Profile, p.Label)
}

func GetTopic(a string, p string, l string) string {
	return fmt.Sprintf(defaultTopicTemplate, a, p, l)
}
