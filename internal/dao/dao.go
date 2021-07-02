package dao

import (
	"context"

	"github.com/go-kratos/kratos/pkg/cache/redis"
	"github.com/google/wire"
	"github.com/znyh/middle-end/library/pkg/kafka"
)

var Provider = wire.NewSet(New, NewRedis, NewKafkaProducer)

type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
	RedisPublish(ctx context.Context, topic, data string) (err error)
	PubKafka(topic string, data []byte) (err error)
}

// dao dao.
type dao struct {
	redis    *redis.Redis
	producer *kafka.Producer
}

// New new a dao and return.
func New(r *redis.Redis, p *kafka.Producer) (d Dao, cf func(), err error) {
	d = &dao{
		redis:    r,
		producer: p,
	}
	cf = d.Close
	return
}

// Close close the resource.
func (d *dao) Close() {
	_ = d.redis.Close()
	_ = d.producer.Close()
}

// Ping ping the resource.
func (d *dao) Ping(ctx context.Context) (err error) {
	return nil
}
