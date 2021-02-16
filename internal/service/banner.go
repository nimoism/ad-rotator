package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/nimoism/ad-rotator/internal/entity"
	"github.com/nimoism/ad-rotator/internal/log"
	"github.com/nimoism/ad-rotator/pkg/ucb1"
)

type BannerRepo interface {
	BannerStats(ctx context.Context, slotID, userGroupID int) ([]entity.BannerStat, error)
	CreateBannerClick(ctx context.Context, click entity.ClickEvent) error
	CreateBannerShow(ctx context.Context, show entity.ShowEvent) error
	Banners(ctx context.Context) ([]entity.Banner, error)
	CreateBanner(ctx context.Context, banner *entity.Banner) error
	UpdateBanner(ctx context.Context, banner *entity.Banner) error
	DeleteBanner(ctx context.Context, id int) error
	CreateBannerSlot(ctx context.Context, bannerID, slotID int) error
	RemoveBannerSlot(ctx context.Context, bannerID, slotID int) error
	SlotsByBanner(ctx context.Context, bannerID int) ([]entity.Slot, error)
}

type BannerService struct {
	log  log.Logger
	repo BannerRepo
}

func NewBannerService(log log.Logger, repo BannerRepo) *BannerService {
	return &BannerService{log: log, repo: repo}
}

func (s *BannerService) ChooseBanner(ctx context.Context, slotID, userGroupID int) (entity.Banner, error) {
	stats, err := s.repo.BannerStats(ctx, slotID, userGroupID)
	if err != nil {
		return entity.Banner{}, fmt.Errorf("banner getting error: %w", err)
	}
	if len(stats) < 1 {
		return entity.Banner{}, errors.New("no next banner found")
	}

	ucbArms := make([]ucb1.Arm, 0, len(stats))
	for _, stat := range stats {
		ucbArms = append(ucbArms, ucb1.NewArm(stat.ShowCount, stat.ClickCount))
	}
	index, _ := ucb1.NextArm(ucbArms)

	banner := stats[index].Banner
	showEvent := entity.ShowEvent{BannerID: banner.ID, SlotID: slotID, UserGroupID: userGroupID}
	if err = s.repo.CreateBannerShow(ctx, showEvent); err != nil {
		return banner, fmt.Errorf("show event registration error: %w", err)
	}
	// TODO: send to kafka
	return banner, nil
}

func (s *BannerService) Click(ctx context.Context, click entity.ClickEvent) error {
	if err := s.repo.CreateBannerClick(ctx, click); err != nil {
		return err
	}
	// TODO: send to kafka
	return nil
}

func (s *BannerService) Banners(ctx context.Context) ([]entity.Banner, error) {
	return s.repo.Banners(ctx)
}

func (s *BannerService) CreateBanner(ctx context.Context, banner *entity.Banner) error {
	return s.repo.CreateBanner(ctx, banner)
}

func (s *BannerService) UpdateBanner(ctx context.Context, banner *entity.Banner) error {
	return s.repo.UpdateBanner(ctx, banner)
}

func (s *BannerService) DeleteBanner(ctx context.Context, id int) error {
	return s.repo.DeleteBanner(ctx, id)
}

func (s *BannerService) BindBannerToSlot(ctx context.Context, bannerID, slotID int) error {
	return s.repo.CreateBannerSlot(ctx, bannerID, slotID)
}

func (s *BannerService) UnbindBannerFromSlot(ctx context.Context, bannerID, slotID int) error {
	return s.repo.RemoveBannerSlot(ctx, bannerID, slotID)
}

func (s *BannerService) BoundSlots(ctx context.Context, bannerID int) ([]entity.Slot, error) {
	return s.repo.SlotsByBanner(ctx, bannerID)
}
