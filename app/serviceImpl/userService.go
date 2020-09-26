package serviceImpl

import (
	"context"
	"odin/app/entity"
	"odin/app/repository"

	"odin/proto"
)

type UserService struct {
}

func NewUserService() *UserService {
	return new(UserService)
}

func (this *UserService) UserInfo(ctx context.Context, req *proto.UserInfoRequest, resp *proto.UserInfoResponse) error {
	userRepo, err := repository.GetUserRepo(ctx)
	if err != nil {
		return err
	}
	user, err := userRepo.GetUserInfo(ctx, req.Id)
	if err != nil {
		return err
	}
	resp.Name = user.Name
	resp.Age = user.Age
	resp.City = user.City
	return nil
}
func (this *UserService) AddUser(ctx context.Context, req *proto.AddUserRequest, resp *proto.AddUserResponse) error {
	userRepo, err := repository.GetUserRepo(ctx)
	if err != nil {
		return err
	}
	user := &entity.User{
		Name: req.Name,
		Age:  req.Age,
		City: req.City,
	}
	id, err := userRepo.AddUser(ctx, user)
	if err != nil {
		return err
	}
	resp.Id = id
	return nil
}

func (this *UserService) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest, resp *proto.UpdateUserResponse) error {
	userRepo, err := repository.GetUserRepo(ctx)
	if err != nil {
		return err
	}
	user := &entity.User{
		Id:   req.Id,
		Name: req.Name,
		Age:  req.Age,
		City: req.City,
	}
	err = userRepo.UpdateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
