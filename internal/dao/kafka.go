package dao

import (
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/log"
	"github.com/znyh/library/pkg/kafka"
)

func NewKafkaProducer() (producer *kafka.Producer, err error) {
	cfg := &kafka.Config{}
	if err = paladin.Watch("kafka.txt", cfg); err != nil {
		return nil, err
	}

	producer = &kafka.Producer{}
	log.Info("[NewKafka] addrs = %v", cfg.Addr)
	*producer, err = kafka.NewProducer(cfg.Addr)
	if err != nil {
		log.Error("new kafka producer fail, msg: %s", err.Error())
	}
	return
}

// 推送 "publisher.push"
func (d *dao) PubKafka(topic string, data []byte) (err error) {
	err = d.producer.Producer(kafka.Message{
		Topic: topic,
		Value: data,
	})
	return
}
