package service

import (
	"context"

	"github.com/nimoism/ad-rotator/internal/api/grpc/pb"
	"github.com/nimoism/ad-rotator/internal/entity"
	"github.com/nimoism/ad-rotator/internal/log"
)

type UserGroupRepo interface {
	UsersGroups(ctx context.Context) ([]entity.UserGroup, error)
	CreateUserGroup(context.Context, *entity.UserGroup) error
	UpdateUserGroup(context.Context, *entity.UserGroup) error
	DeleteUserGroup(ctx context.Context, ID int) error
}

type UserGroupService struct {
	pb.UnimplementedUserGroupsServer
	log  log.Logger
	repo UserGroupRepo
}

func NewUserGroupService(log log.Logger, repo UserGroupRepo) *UserGroupService {
	return &UserGroupService{
		log:  log,
		repo: repo,
	}
}

func (s *UserGroupService) UsersGroups(ctx context.Context) ([]entity.UserGroup, error) {
	return s.repo.UsersGroups(ctx)
}

func (s *UserGroupService) CreateUserGroup(ctx context.Context, ug *entity.UserGroup) error {
	return s.repo.CreateUserGroup(ctx, ug)
}

func (s *UserGroupService) UpdateUserGroup(ctx context.Context, ug *entity.UserGroup) error {
	return s.repo.UpdateUserGroup(ctx, ug)
}

func (s *UserGroupService) DeleteUserGroup(ctx context.Context, id int) error {
	return s.repo.DeleteUserGroup(ctx, id)
}
