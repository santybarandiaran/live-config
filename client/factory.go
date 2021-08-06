package client

import (
	http2 "live-config/client/http"
	"live-config/server/redis"
	"net/http"
	"time"
)

func New() *http2.PropertyClient {
	c := createHttpClient()
	b := createRedisBroker()

	return &http2.PropertyClient{ H: c, B: b }
}

func createHttpClient() *http.Client {
	t := createTransportConfig()

	return  &http.Client{
		Timeout:   10 * time.Second,
		Transport: t,
	}
}

func createTransportConfig() *http.Transport {
	config := http2.New()

	t := http.DefaultTransport.(*http.Transport).Clone()

	t.MaxIdleConns = *config.MaxIdleConns
	t.MaxConnsPerHost = *config.MaxConnsPerHost
	t.MaxIdleConnsPerHost = *config.MaxIdleConnsPerHost
	return t
}

func createRedisBroker() *redis.MessageBroker {
	return &redis.MessageBroker{Redis: redis.Init()}
}
