package key

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/Team-Work-Forever/FireWatchRest/config"
	"github.com/redis/go-redis/v9"
)

type (
	KeyValueService struct {
		client *redis.Client
		ctx    context.Context
	}
)

func NewKeyValueService() *KeyValueService {
	env := config.GetCofig()

	return &KeyValueService{
		client: redis.NewClient(&redis.Options{
			Addr:      fmt.Sprintf("%s:%s", env.REDIS_HOST, env.REDIS_PORT),
			Username:  env.REDIS_USER,
			Password:  env.REDIS_PASSWD,
			DB:        env.REDIS_DB,
			TLSConfig: &tls.Config{},
		}),
		ctx: context.Background(),
	}
}

func (kv *KeyValueService) Store(ve ValueEntity) error {
	ekv := ve.GetKV()

	return kv.client.Set(kv.ctx, ekv.Key.Build(), ekv.Value, ekv.Exp).Err()
}

func (kv *KeyValueService) Get(key Key, ve ValueEntity) error {
	value, err := kv.client.Get(kv.ctx, key.Build()).Result()

	if err != nil {
		return err
	}

	ve.Scan(key, value)
	return nil
}

func (kv *KeyValueService) Delete(ve ValueEntity) error {
	ekv := ve.GetKV()
	return kv.client.Del(kv.ctx, ekv.Key.Build()).Err()
}

func (kv *KeyValueService) Close() error {
	return kv.client.Close()
}
