package repository

import (
	"context"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"internal_api/global"
	"strings"
	"time"
)

var ctx = context.Background()

type Repo interface {
	i()
	Set(key, value string, ttl time.Duration) error
	Get(key string) (string, error)
	TTL(key string) (time.Duration, error)
	Expire(key string, ttl time.Duration) bool
	ExpireAt(key string, ttl time.Time) bool
	Del(key string) bool
	Exists(keys ...string) bool
	Incr(key string) int64
	Close() error
	Version() string
}

type cacheRepo struct {
	client *redis.Client
}

func New(client *redis.Client) Repo {
	//client, err := redisConnect()
	//if err != nil {
	//	return nil, err
	//}

	return &cacheRepo{
		client: client,
	}
}

func (c *cacheRepo) i() {}

func RedisConnect() (*redis.Client, error) {
	cfg := global.CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		Password:     cfg.Password,
		DB:           cfg.DB,
		MaxRetries:   cfg.MaxRetries,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConn,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, errors.Wrap(err, "ping redis err")
	}

	return client, nil
}

// Set set some <key,value> into redis
func (c *cacheRepo) Set(key, value string, ttl time.Duration) error {
	if err := c.client.Set(context.Background(), key, value, ttl).Err(); err != nil {
		return errors.Wrapf(err, "redis set key: %s err", key)
	}

	return nil
}

// Get get some key from redis
func (c *cacheRepo) Get(key string) (string, error) {

	value, err := c.client.Get(context.Background(), key).Result()
	if err != nil {
		return "", errors.Wrapf(err, "redis get key: %s err", key)
	}

	return value, nil
}

// TTL get some key from redis
func (c *cacheRepo) TTL(key string) (time.Duration, error) {
	ttl, err := c.client.TTL(context.Background(), key).Result()
	if err != nil {
		return -1, errors.Wrapf(err, "redis get key: %s err", key)
	}

	return ttl, nil
}

// Expire expire some key
func (c *cacheRepo) Expire(key string, ttl time.Duration) bool {
	ok, _ := c.client.Expire(ctx, key, ttl).Result()
	return ok
}

// ExpireAt expire some key at some time
func (c *cacheRepo) ExpireAt(key string, ttl time.Time) bool {
	ok, _ := c.client.ExpireAt(ctx, key, ttl).Result()
	return ok
}

func (c *cacheRepo) Exists(keys ...string) bool {
	if len(keys) == 0 {
		return true
	}
	value, _ := c.client.Exists(ctx, keys...).Result()
	return value > 0
}

func (c *cacheRepo) Del(key string) bool {
	if key == "" {
		return true
	}

	value, _ := c.client.Del(ctx, key).Result()
	return value > 0
}

func (c *cacheRepo) Incr(key string) int64 {
	value, _ := c.client.Incr(ctx, key).Result()
	return value
}

// Close close redis client
func (c *cacheRepo) Close() error {
	return c.client.Close()
}

// Version redis server version
func (c *cacheRepo) Version() string {
	server := c.client.Info(ctx, "server").Val()
	spl1 := strings.Split(server, "# Server")
	spl2 := strings.Split(spl1[1], "redis_version:")
	spl3 := strings.Split(spl2[1], "redis_git_sha1:")
	return spl3[0]
}
