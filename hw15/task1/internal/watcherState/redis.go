package watcherstate

import (
	"context"
	"log"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type RedisState struct {
	redis *redis.Client
}

func NewRedisState() *RedisState {

	return &RedisState{
		redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
			DB:   0,
		})}
}

func (s *RedisState) Set(path string, id uint) {
	ctx := context.Background()

	err := s.redis.Set(ctx, path, id, 0).Err()
	if err != nil {
		log.Fatalln(err)
	}
}

func (s *RedisState) Get(path string) uint {
	ctx := context.Background()

	val, err := s.redis.Get(ctx, path).Result()
	if err != nil {
		return 0
	}

	u64, err := strconv.ParseUint(val, 10, 32)
	if err != nil {
		log.Fatalln(err)
	}
	return uint(u64)
}
