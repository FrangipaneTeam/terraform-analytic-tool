package clients

import (
	"crypto/tls"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Address      string `yaml:"address"`      // host:port address.
	Password     string `yaml:"password"`     // Optional password.
	MaxRetries   int    `yaml:"maxRetries"`   // Maximum number of retries before giving up.
	DialTimeout  int    `yaml:"dialTimeout"`  // in seconds
	ReadTimeout  int    `yaml:"readTimeout"`  // in seconds
	WriteTimeout int    `yaml:"writeTimeout"` // in seconds
	TLSConfig    *tls.Config
}

type RedisClient struct {
	*redis.Client
}

func NewRedisClient(r RedisConfig) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:        r.Address,
		Password:    r.Password,
		MaxRetries:  r.MaxRetries,
		DialTimeout: time.Duration(r.DialTimeout) * time.Second,
		ReadTimeout: time.Duration(r.ReadTimeout) * time.Second,
		TLSConfig:   r.TLSConfig,
	})

	return &RedisClient{client}
}
