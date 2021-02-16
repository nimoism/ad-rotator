package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/nimoism/ad-rotator/internal/api/grpc/pb"
	"github.com/nimoism/ad-rotator/internal/log"
)

var statusInternalError = status.Error(codes.Internal, "internal error")

type Service interface {
	BannerService
	SlotService
	UserGroupService
}

func NewAPIServer(log log.Logger, service Service) *grpc.Server {
	api := grpc.NewServer()
	pb.RegisterBannersServer(api, NewBannerServer(log, service))
	pb.RegisterSlotsServer(api, NewSlotServer(log, service))
	pb.RegisterUserGroupsServer(api, NewUserGroupServer(log, service))
	return api
}
