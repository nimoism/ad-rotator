package grpc //nolint:dupl

import (
	"context"

	"github.com/nimoism/ad-rotator/internal/api/grpc/pb"
	"github.com/nimoism/ad-rotator/internal/entity"
	"github.com/nimoism/ad-rotator/internal/log"
)

type UserGroupService interface {
	UsersGroups(ctx context.Context) ([]entity.UserGroup, error)
	CreateUserGroup(ctx context.Context, slot *entity.UserGroup) error
	UpdateUserGroup(ctx context.Context, slot *entity.UserGroup) error
	DeleteUserGroup(ctx context.Context, id int) error
}

type UserGroupServer struct {
	pb.UnimplementedUserGroupsServer
	log log.Logger
	ugs UserGroupService
}

func NewUserGroupServer(log log.Logger, ugService UserGroupService) *UserGroupServer {
	return &UserGroupServer{log: log, ugs: ugService}
}

func (s *UserGroupServer) UserGroups(ctx context.Context, _ *pb.AllUGsRequest) (*pb.AllUGsResult, error) {
	ugs, err := s.ugs.UsersGroups(ctx)
	if err != nil {
		s.log.Error(err)
		return nil, statusInternalError
	}
	apiUGs := make([]*pb.UserGroup, 0, len(ugs))
	for _, ug := range ugs {
		apiUGs = append(apiUGs, &pb.UserGroup{
			Id:   int64(ug.ID),
			Name: ug.Name,
		})
	}
	return &pb.AllUGsResult{Ugs: apiUGs}, nil
}

func (s *UserGroupServer) CreateUserGroup(ctx context.Context, req *pb.CreateUGRequest) (*pb.CreateUGResult, error) {
	ug := entity.UserGroup{
		Name: req.Ug.Name,
	}
	if err := s.ugs.CreateUserGroup(ctx, &ug); err != nil {
		s.log.Error(err)
		return nil, statusInternalError
	}
	return &pb.CreateUGResult{Ug: &pb.UserGroup{Id: int64(ug.ID), Name: ug.Name}}, nil
}

func (s *UserGroupServer) UpdateUserGroup(ctx context.Context, req *pb.UpdateUGRequest) (*pb.UpdateUGResult, error) {
	ug := entity.UserGroup{ID: int(req.Ug.Id), Name: req.Ug.Name}
	err := s.ugs.UpdateUserGroup(ctx, &ug)
	if err != nil {
		s.log.Error(err)
		return nil, statusInternalError
	}
	return &pb.UpdateUGResult{Ug: &pb.UserGroup{
		Id:   int64(ug.ID),
		Name: ug.Name,
	}}, nil
}

func (s *UserGroupServer) DeleteUserGroup(ctx context.Context, req *pb.DeleteUGRequest) (*pb.DeleteUGResult, error) {
	if err := s.ugs.DeleteUserGroup(ctx, int(req.Id)); err != nil {
		s.log.Error(err)
		return nil, statusInternalError
	}
	return &pb.DeleteUGResult{}, nil
}
