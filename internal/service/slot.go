package service

import (
	"context"

	"github.com/nimoism/ad-rotator/internal/entity"
	"github.com/nimoism/ad-rotator/internal/log"
)

type SlotRepo interface {
	Slots(context.Context) ([]entity.Slot, error)
	CreateSlot(context.Context, *entity.Slot) error
	UpdateSlot(context.Context, *entity.Slot) error
	DeleteSlot(ctx context.Context, id int) error
}

type SlotService struct {
	log  log.Logger
	repo SlotRepo
}

func NewSlotService(log log.Logger, repo SlotRepo) *SlotService {
	return &SlotService{
		log:  log,
		repo: repo,
	}
}

func (s *SlotService) Slots(ctx context.Context) ([]entity.Slot, error) {
	return s.repo.Slots(ctx)
}

func (s *SlotService) CreateSlot(ctx context.Context, slot *entity.Slot) error {
	return s.repo.CreateSlot(ctx, slot)
}

func (s *SlotService) UpdateSlot(ctx context.Context, slot *entity.Slot) error {
	return s.repo.UpdateSlot(ctx, slot)
}

func (s *SlotService) DeleteSlot(ctx context.Context, id int) error {
	return s.repo.DeleteSlot(ctx, id)
}
