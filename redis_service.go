package chatserver

import "github.com/go-redis/redis"

type RedisService struct {
	client *redis.Client
}

func (s *RedisService) Stop() error {
	return s.client.Close()
}

func NewRedisService() *RedisService {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &RedisService{
		client: client,
	}
}

func (s *RedisService) SendMessage(msg string) {
	s.client.Publish("chatroom 1", msg)
}
