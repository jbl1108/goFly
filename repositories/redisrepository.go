package repositories

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jbl1108/goFly/util"
	"go.uber.org/multierr"
)

type RedisRepository struct {
	client *redis.Client
	config *util.Config
}

func NewRedisRepository(config *util.Config) *RedisRepository {
	restClient := new(RedisRepository)
	restClient.config = config
	restClient.client = redis.NewClient(&redis.Options{
		Addr:     config.RedisDBAddr(),
		Password: "", // no password setdo
		DB:       0,  // use default DB
	})
	return restClient
}

func (m *RedisRepository) StoreString(key string, value string) (err error) {
	ctx := context.Background()
	err = m.client.Set(ctx, key, value, 0).Err()
	return
}

func (m *RedisRepository) FetchString(key string) (value string, err error) {
	ctx := context.Background()
	valuecdm := m.client.Get(ctx, key)
	value, err = valuecdm.Result()
	return
}
func (m *RedisRepository) StoreList(key string, values []string) (err error) {
	ctx := context.Background()
	err = m.client.Del(ctx, key).Err()
	err = multierr.Combine(err, m.client.RPush(ctx, key, values).Err())
	return err
}

func (m *RedisRepository) AppendToList(key string, values []string) (err error) {
	ctx := context.Background()
	err = m.client.RPush(ctx, key, values).Err()
	return err
}

func (m *RedisRepository) FetchList(key string) (value []string, err error) {
	ctx := context.Background()
	value, err = m.client.LRange(ctx, key, 0, -1).Result()
	return
}
