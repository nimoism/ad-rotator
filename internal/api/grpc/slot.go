package grpc //nolint:dupl

import (
	"context"

	"github.com/nimoism/ad-rotator/internal/api/grpc/pb"
	"github.com/nimoism/ad-rotator/internal/entity"
	"github.com/nimoism/ad-rotator/internal/log"
)

type SlotService interface {
	Slots(ctx context.Context) ([]entity.Slot, error)
	CreateSlot(ctx context.Context, slot *entity.Slot) error
	UpdateSlot(ctx context.Context, slot *entity.Slot) error
	DeleteSlot(ctx context.Context, id int) error
}

type SlotServer struct {
	pb.UnimplementedSlotsServer
	log   log.Logger
	slots SlotService
}

func NewSlotServer(log log.Logger, slotService SlotService) *SlotServer {
	return &SlotServer{log: log, slots: slotService}
}

func (s *SlotServer) Slots(ctx context.Context, _ *pb.AllSlotsRequest) (*pb.AllSlotsResult, error) {
	slots, err := s.slots.Slots(ctx)
	if err != nil {
		s.log.Error(err)
		return nil, statusInternalError
	}
	apiSlots := make([]*pb.Slot, 0, len(slots))
	for _, slot := range slots {
		apiSlots = append(apiSlots, &pb.Slot{
			Id:   int64(slot.ID),
			Name: slot.Name,
		})
	}
	return &pb.AllSlotsResult{Slots: apiSlots}, nil
}

func (s *SlotServer) CreateSlot(ctx context.Context, req *pb.CreateSlotRequest) (*pb.CreateSlotResult, error) {
	slot := entity.Slot{
		Name: req.Slot.Name,
	}
	if err := s.slots.CreateSlot(ctx, &slot); err != nil {
		s.log.Error(err)
		return nil, statusInternalError
	}
	return &pb.CreateSlotResult{Slot: &pb.Slot{Id: int64(slot.ID), Name: slot.Name}}, nil
}

func (s *SlotServer) UpdateSlot(ctx context.Context, req *pb.UpdateSlotRequest) (*pb.UpdateSlotResult, error) {
	slot := entity.Slot{ID: int(req.Slot.Id), Name: req.Slot.Name}
	err := s.slots.UpdateSlot(ctx, &slot)
	if err != nil {
		s.log.Error(err)
		return nil, statusInternalError
	}
	return &pb.UpdateSlotResult{Slot: &pb.Slot{
		Id:   int64(slot.ID),
		Name: slot.Name,
	}}, nil
}

func (s *SlotServer) DeleteSlot(ctx context.Context, req *pb.DeleteSlotRequest) (*pb.DeleteSlotResult, error) {
	if err := s.slots.DeleteSlot(ctx, int(req.Id)); err != nil {
		s.log.Error(err)
		return nil, statusInternalError
	}
	return &pb.DeleteSlotResult{}, nil
}
