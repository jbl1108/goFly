package driver

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jbl1108/goFly/util"
)

type RedisDriver struct {
	client *redis.Client
	config *util.Config
}

func NewRedisDriver(config *util.Config) *RedisDriver {
	restClient := new(RedisDriver)
	restClient.config = config
	restClient.client = redis.NewClient(&redis.Options{
		Addr:     config.RedisDBAddr(),
		Password: "", // no password setdo
		DB:       0,  // use default DB
	})
	return restClient
}

func (m *RedisDriver) StoreString(key string, value string) (err error) {
	ctx := context.Background()
	err = m.client.Set(ctx, key, value, 0).Err()
	return
}

func (m *RedisDriver) FetchString(key string) (value string, err error) {
	ctx := context.Background()
	valuecdm := m.client.Get(ctx, key)
	value, err = valuecdm.Result()
	return
}
func (m *RedisDriver) StoreList(key string, values []string) (err error) {
	ctx := context.Background()
	err = m.client.Del(ctx, key).Err()
	err = m.client.RPush(ctx, key, values).Err()
	return
}

func (m *RedisDriver) FetchList(key string) (value []string, err error) {
	ctx := context.Background()
	value, err = m.client.LRange(ctx, key, 0, -1).Result()
	return
}
