package grpc

import (
	"context"
	"time"

	"github.com/nimoism/ad-rotator/internal/api/grpc/pb"
	"github.com/nimoism/ad-rotator/internal/entity"
	"github.com/nimoism/ad-rotator/internal/log"
)

type BannerService interface {
	ChooseBanner(ctx context.Context, slotID, userGroupID int) (entity.Banner, error)
	Click(ctx context.Context, event entity.ClickEvent) error
	Banners(ctx context.Context) ([]entity.Banner, error)
	CreateBanner(ctx context.Context, slot *entity.Banner) error
	UpdateBanner(ctx context.Context, slot *entity.Banner) error
	DeleteBanner(ctx context.Context, id int) error
	BindBannerToSlot(ctx context.Context, bannerID, slotID int) error
	UnbindBannerFromSlot(ctx context.Context, bannerID, slotID int) error
	BoundSlots(ctx context.Context, bannerID int) ([]entity.Slot, error)
}

type BannerServer struct {
	pb.UnimplementedBannersServer
	log     log.Logger
	banners BannerService
}

func NewBannerServer(log log.Logger, banners BannerService) *BannerServer {
	return &BannerServer{log: log, banners: banners}
}

func (s *BannerServer) Banner(ctx context.Context, req *pb.BannerRequest) (*pb.BannerResult, error) {
	banner, err := s.banners.ChooseBanner(ctx, int(req.SlotId), int(req.UserGroupId))
	if err != nil {
		s.log.Error(err)
		return nil, statusInternalError
	}
	s.log.Debugf("got banner %v", banner.ID)
	return &pb.BannerResult{Banner: &pb.Banner{
		Id:   int64(banner.ID),
		Name: banner.Name,
	}}, nil
}

func (s *BannerServer) Click(ctx context.Context, req *pb.ClickRequest) (*pb.ClickResult, error) {
	click := entity.ClickEvent{
		BannerID:    int(req.BannerId),
		SlotID:      int(req.SlotId),
		UserGroupID: int(req.UserGroupId),
		Created:     time.Now(),
	}
	if err := s.banners.Click(ctx, click); err != nil {
		s.log.Error(err)
		return nil, statusInternalError
	}
	return &pb.ClickResult{}, nil
}

func (s *BannerServer) Banners(ctx context.Context, _ *pb.AllBannersRequest) (*pb.AllBannersResult, error) {
	banners, err := s.banners.Banners(ctx)
	if err != nil {
		s.log.Error(err)
		return nil, statusInternalError
	}
	apiBanners := make([]*pb.Banner, 0, len(banners))
	for _, banner := range banners {
		apiBanners = append(apiBanners, &pb.Banner{
			Id:   int64(banner.ID),
			Name: banner.Name,
		})
	}
	return &pb.AllBannersResult{Banners: apiBanners}, nil
}

func (s *BannerServer) CreateBanner(ctx context.Context, req *pb.CreateBannerRequest) (*pb.CreateBannerResult, error) {
	banner := entity.Banner{Name: req.Banner.Name}
	err := s.banners.CreateBanner(ctx, &banner)
	if err != nil {
		s.log.Error(err)
		return nil, statusInternalError
	}
	return &pb.CreateBannerResult{Banner: &pb.Banner{
		Id:   int64(banner.ID),
		Name: banner.Name,
	}}, nil
}

func (s *BannerServer) UpdateBanner(ctx context.Context, req *pb.UpdateBannerRequest) (*pb.UpdateBannerResult, error) {
	banner := entity.Banner{ID: int(req.Banner.Id), Name: req.Banner.Name}
	err := s.banners.UpdateBanner(ctx, &banner)
	if err != nil {
		s.log.Error(err)
		return nil, statusInternalError
	}
	return &pb.UpdateBannerResult{Banner: &pb.Banner{
		Id:   int64(banner.ID),
		Name: banner.Name,
	}}, nil
}

func (s *BannerServer) DeleteBanner(ctx context.Context, req *pb.DeleteBannerRequest) (*pb.DeleteBannerResult, error) {
	err := s.banners.DeleteBanner(ctx, int(req.Id))
	return &pb.DeleteBannerResult{}, err
}

func (s *BannerServer) BoundSlots(ctx context.Context, req *pb.BoundSlotsRequest) (*pb.BoundSlotsResult, error) {
	slots, err := s.banners.BoundSlots(ctx, int(req.BannerId))
	if err != nil {
		s.log.Error(statusInternalError)
		return nil, statusInternalError
	}
	apiSlots := make([]*pb.Slot, 0, len(slots))
	for _, slot := range slots {
		apiSlots = append(apiSlots, &pb.Slot{
			Id:   int64(slot.ID),
			Name: slot.Name,
		})
	}
	return &pb.BoundSlotsResult{Slots: apiSlots}, nil
}

func (s *BannerServer) BindSlot(ctx context.Context, req *pb.BindSlotRequest) (*pb.BindSlotResult, error) {
	err := s.banners.BindBannerToSlot(ctx, int(req.BannerId), int(req.SlotId))
	if err != nil {
		s.log.Error(err)
		return nil, statusInternalError
	}
	return &pb.BindSlotResult{}, nil
}

func (s *BannerServer) UnbindSlot(ctx context.Context, req *pb.UnbindSlotRequest) (*pb.UnbindSlotResult, error) {
	err := s.banners.UnbindBannerFromSlot(ctx, int(req.BannerId), int(req.SlotId))
	if err != nil {
		s.log.Error(err)
		return nil, statusInternalError
	}
	return &pb.UnbindSlotResult{}, nil
}
