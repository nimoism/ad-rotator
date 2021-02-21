package kafka

import (
	"fmt"
	"net"
	"strconv"

	"github.com/segmentio/kafka-go"

	"github.com/nimoism/ad-rotator/internal/log"
)

const TopicEvent = "events"

var topics = []kafka.TopicConfig{
	{Topic: TopicEvent, NumPartitions: 1, ReplicationFactor: 1},
}

type Stream struct {
	*BannerStream
}

func NewStream(log log.Logger, host string) (*Stream, error) {
	if err := createTopics(host); err != nil {
		return nil, err
	}
	return &Stream{
		BannerStream: NewBannerStream(log, host),
	}, nil
}

func createTopics(host string) error {
	conn, err := kafka.Dial("tcp", host)
	if err != nil {
		return fmt.Errorf("kafka stream connection error: %w", err)
	}
	controller, err := conn.Controller()
	if err != nil {
		return fmt.Errorf("kafka stream getting controller error: %w", err)
	}
	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		return fmt.Errorf("kafka stream controller connection error: %w", err)
	}
	for _, topicConf := range topics {
		if err = controllerConn.CreateTopics(topicConf); err != nil {
			return fmt.Errorf("kafka topic creating error: %w", err)
		}
	}
	return nil
}
