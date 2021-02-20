package kafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/nimoism/ad-rotator/internal/api/grpc/pb"
	"github.com/nimoism/ad-rotator/internal/entity"
	"github.com/nimoism/ad-rotator/internal/log"
)

type BannerStream struct {
	log    log.Logger
	writer *kafka.Writer
}

func NewBannerStream(log log.Logger, host string) *BannerStream {
	return &BannerStream{
		log: log,
		writer: &kafka.Writer{
			Addr:  kafka.TCP(host),
			Topic: TopicEvent,
			Async: true,
		},
	}
}

func (s *BannerStream) SendBannerClick(ctx context.Context, click entity.ClickEvent) error {
	err := s.sendBannerEvent(ctx, pb.BannerEvent_CLICK, &pb.BannerEvent{
		Type:     pb.BannerEvent_CLICK,
		BannerId: int64(click.BannerID),
		SlotId:   int64(click.SlotID),
		UgId:     int64(click.UserGroupID),
		Dt:       timestamppb.New(click.Created),
	})
	if err != nil {
		return fmt.Errorf("banner click event seinding to stream error: %w", err)
	}
	s.log.Debugf("banner click event sent to stream")
	return nil
}

func (s *BannerStream) SendBannerShow(ctx context.Context, show entity.ShowEvent) error {
	err := s.sendBannerEvent(ctx, pb.BannerEvent_CLICK, &pb.BannerEvent{
		Type:     pb.BannerEvent_CLICK,
		BannerId: int64(show.BannerID),
		SlotId:   int64(show.SlotID),
		UgId:     int64(show.UserGroupID),
		Dt:       timestamppb.New(show.Created),
	})
	if err != nil {
		return fmt.Errorf("banner show event seinding to stream error: %w", err)
	}
	s.log.Debugf("banner show event sent to stream")
	return nil
}

func (s *BannerStream) sendBannerEvent(ctx context.Context, typ pb.BannerEvent_Type, event *pb.BannerEvent) error {
	event.Type = typ
	serialized, err := proto.Marshal(event)
	if err != nil {
		return err
	}
	err = s.writer.WriteMessages(ctx, kafka.Message{
		Value:     serialized,
		Partition: 0,
	})
	return err
}
