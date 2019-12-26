package chatserver

import (
	"log"

	"github.com/go-redis/redis"
)

type RedisService struct {
	client *redis.Client
}

func (s *RedisService) Stop() error {
	return s.client.Close()
}

func NewRedisService(redisURL string) (*RedisService, error) {
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opt)
	return &RedisService{
		client: client,
	}, nil
}

func (s *RedisService) SendMessage(msg string) {
	resp := s.client.Publish("chatroom 1", msg)
	if resp.Err() != nil {
		log.Println("failed to publish", resp.Err())
	}
}
