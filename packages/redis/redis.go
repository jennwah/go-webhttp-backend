package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Host string
	Port string
	Pass string
}

func (c Config) addressString() string {
	return fmt.Sprintf("%v:%v", c.Host, c.Port)
}

// NewClient returns Redis Client connection.
func NewClient(c Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     c.addressString(),
		Password: c.Pass,
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}

	return client, nil
}
