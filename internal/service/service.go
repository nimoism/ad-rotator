package service

import (
	"github.com/nimoism/ad-rotator/internal/log"
)

type Repo interface {
	BannerRepo
	SlotRepo
	UserGroupRepo
}

type Service struct {
	*BannerService
	*SlotService
	*UserGroupService
}

func NewService(log log.Logger, repo Repo) *Service {
	return &Service{
		BannerService:    NewBannerService(log, repo),
		SlotService:      NewSlotService(log, repo),
		UserGroupService: NewUserGroupService(log, repo),
	}
}
