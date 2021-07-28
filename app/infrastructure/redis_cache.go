package infrastructure

import (
	"context"
	"time"

	"github.com/go-redis/redis"
	"github.com/kyawmyintthein/rzlog"
)

type Cache interface {
	Get(context.Context, string) (string, bool, error)
	Set(context.Context, string, interface{}, time.Duration) error
	SetNX(context.Context, string, interface{}, time.Duration) (bool, error)
	Delete(context.Context, string) error
	Ping(context.Context) error
}

type RedisCfg struct {
	EnableCluster        bool     `json:"enable_cluster" envconfig:"ENABLE_CLUSTER" required:"true" default:"false"`
	Addrs                []string `json:"addrs" envconfig:"ADDRS" required:"true" default:"localhost:6379"`
	Password             string   `json:"-" envconfig:"PASSWORD" default:""`
	PoolSize             int      `mapstructure:"pool_size" json:"pool_size" envconfig:"POOL_SIZE" required:"true" default:"10"`
	PoolTimeoutInSeconds int64    `mapstructure:"pool_timeout_in_seconds" json:"pool_timeout_in_seconds" envconfig:"POOL_TIMEOUT_IN_SECONDS" default:"10"`
	MaxConnAgeInSeconds  int64    `mapstructure:"max_conn_age_in_seconds" json:"max_conn_age_in_seconds" envconfig:"MAX_CONN_AGE_IN_SECONDS"  default:"10"`
	MinIdleConns         int      `mapstructure:"min_idle_conns" json:"min_idle_conns" envconfig:"MIN_IDLE_CONNS"  default:"10"`
	IdleTimeoutInSeconds int64    `mapstructure:"idle_timeout_in_seconds" json:"idle_timeout_in_seconds" envconfig:"IDLE_TIMEOUT_IN_SECONDS" default:"30"`
}

type redisClusterCache struct {
	config *RedisCfg
	client redis.Cmdable
}

func NewRedisClient(config *RedisCfg) (redis.Cmdable, error) {
	if config.EnableCluster {
		redisClient := redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:        config.Addrs,
			Password:     config.Password,
			PoolSize:     config.PoolSize,
			MinIdleConns: config.MinIdleConns,
			PoolTimeout:  time.Duration(config.PoolTimeoutInSeconds) * time.Second,
			IdleTimeout:  time.Duration(config.IdleTimeoutInSeconds) * time.Second,
			MaxConnAge:   time.Duration(config.MaxConnAgeInSeconds) * time.Second,
		})

		err := redisClient.Ping().Err()
		if err != nil {
			return nil, err
		}

		return redisClient, nil
	}
	if len(config.Addrs) == 0 {
		config.Addrs = append(config.Addrs, "localhost:6379")
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Addrs[0],
		Password: config.Password,
	})

	err := redisClient.Ping().Err()
	if err != nil {
		return nil, err
	}
	return redisClient, nil
}

func NewRedisCache(config *RedisCfg, redisClient redis.Cmdable) (Cache, error) {
	return &redisClusterCache{
		config: config,
		client: redisClient,
	}, nil
}

func (rc *redisClusterCache) Get(ctx context.Context, key string) (string, bool, error) {
	val, err := rc.client.Get(key).Result()
	if err != nil {
		if err != redis.Nil {
			rzlog.Error(ctx, err, "error while retrieving from redis. Error: %v \n", err)
			return val, false, nil
		}
		return val, false, err
	}
	return val, true, nil
}

func (rc *redisClusterCache) Set(ctx context.Context, key string, val interface{}, expiry time.Duration) error {
	err := rc.client.Set(key, val, expiry).Err()
	if err != nil {
		rzlog.Error(ctx, err, "failed to set value to redis. Error :", err.Error())
		return err
	}
	return nil
}

func (rc *redisClusterCache) SetNX(ctx context.Context, key string, val interface{}, expiry time.Duration) (bool, error) {
	flag, err := rc.client.SetNX(key, val, expiry).Result()
	if err != nil {
		rzlog.Error(ctx, err, "failed to set value to redis. Error :", err.Error())
		return flag, err
	}
	return flag, nil
}

func (rc *redisClusterCache) Delete(ctx context.Context, key string) error {
	_, err := rc.client.Del(key).Result()
	if err != nil {
		rzlog.Error(ctx, err, "failed to delete from redis. Error :", err.Error())
		return err
	}
	return nil
}

func (rc *redisClusterCache) Ping(ctx context.Context) error {
	_, err := rc.client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
