package kafka

import (
	"github.com/nimoism/ad-rotator/internal/log"
)

const TopicEvent = "events"

type Stream struct {
	*BannerStream
}

func NewStream(log log.Logger, host string) (*Stream, error) {
	// _, err := kafka.DialLeader(context.Background(), "tcp", host, TopicEvent, 0)
	// if err != nil {
	// 	return nil, err
	// }
	// conn, err := kafka.Dial("tcp", host)
	// if err != nil {
	// 	return nil, fmt.Errorf("kafka stream create error: %w", err)
	// }
	// controller, err := conn.Controller()
	// if err != nil {
	// 	panic(err.Error())
	// }
	// var controllerConn *kafka.Conn
	// controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	// if err != nil {
	// 	panic(err.Error())
	// }
	// err = controllerConn.CreateTopics(kafka.TopicConfig{
	// 	Topic: TopicEvent,
	// 	NumPartitions: 1,
	// 	ReplicationFactor: 1,
	// })
	// if err != nil {
	// 	return nil, err
	// }

	return &Stream{
		BannerStream: NewBannerStream(log, host),
	}, nil
}
