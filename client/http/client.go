package http

import (
	"encoding/json"
	"errors"
	"fmt"
	redis2 "github.com/go-redis/redis/v8"
	"io/ioutil"
	"live-config/client/parser"
	"live-config/client/update"
	"live-config/logger"
	"live-config/server/domain"
	"live-config/server/redis"
	"live-config/server/redis/topic"
	"net/http"
)

type PropertyClient struct {
	H *http.Client
	B *redis.MessageBroker
	m map[string]bool
}

func (c *PropertyClient) GetConfig(application string, profile string, label string, out interface{}) error {
	return c.getConfig(application, profile, label, &out)
}

func (c *PropertyClient) GetConfigWithNotice(application string, profile string, label string, out interface{}, u update.Updater) error {
	t := topic.GetTopic(application, profile, label)

	if _, present := c.m[t]; present {
		return errors.New("you are trying to subscribe multiple times for the same application:profile:label")
	}

	if u == nil {
		return errors.New(fmt.Sprintf("Trying to get config from %s:%s:%s but without specifying an updater function", application, profile, label))
	}

	err := c.getConfig(application, profile, label, &out)

	if err != nil {
		return err
	}

	go c.handleSubscription(application, profile, label, &out, u)

	return nil
}

func (c *PropertyClient) handleSubscription(application string, profile string, label string, out *interface{}, u update.Updater) {
	s := c.B.Subscribe(application, profile, label)

	defer closePubSubConnection(application, profile, label, s)

	ch := s.Channel()

	for msg := range ch {
		err := c.getConfig(application, profile, label, out)

		if err != nil {
			logger.Instance.Error(fmt.Sprintf("Error getting config update %s. Message received: %s", err.Error(), msg.Payload))
		}

		var p domain.Property
		err = json.Unmarshal([]byte(msg.Payload), &p)

		if err != nil {
			logger.Instance.Error(fmt.Sprintf("Error parsing message update %s. Message received: %s", err.Error(), msg.Payload))
		}

		u.OnPropertyUpdate(out)
	}
}

func (c *PropertyClient) getConfig(application string, profile string, label string, out *interface{}) error {
	/*
		TODO Create a local cache to validate if there are multiple clients in the app checking for same a/p/l config
		 in that case, panic
	*/

	//TODO make this url configurable
	r, err := c.H.Get(fmt.Sprintf("http://localhost:8080/property/%s/%s/%s", application, profile, label))

	if err != nil {
		return err
	}

	body := r.Body
	defer body.Close()

	all, err := ioutil.ReadAll(body)

	if err != nil {
		return err
	}

	jsonResponse := map[string]interface{}{}
	err = json.Unmarshal(all, &jsonResponse)

	if err != nil {
		return err
	}

	err = parser.Map(jsonResponse, out)

	if err != nil {
		return err
	}

	return nil
}

func closePubSubConnection(application string, profile string, label string, s *redis2.PubSub) {
	func(s *redis2.PubSub) {
		err := s.Close()
		if err != nil {
			logger.Instance.Error(fmt.Sprintf("Error closing redis pub/sub for %s-%s-%s topic", application, profile, label))
		}
		logger.Instance.Info(fmt.Sprintf("Closing redis pub/sub for %s-%s-%s topic", application, profile, label))
	}(s)
}
