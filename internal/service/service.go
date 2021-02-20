package service

import (
	"github.com/nimoism/ad-rotator/internal/log"
)

type Repo interface {
	BannerRepo
	SlotRepo
	UserGroupRepo
}

type Stream interface {
	BannerStream
}

type Service struct {
	*BannerService
	*SlotService
	*UserGroupService
}

func NewService(log log.Logger, repo Repo, stream Stream) *Service {
	return &Service{
		BannerService:    NewBannerService(log, repo, stream),
		SlotService:      NewSlotService(log, repo),
		UserGroupService: NewUserGroupService(log, repo),
	}
}
