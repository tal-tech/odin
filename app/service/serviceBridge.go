package service

import (
	"context"
	"odin/proto"
)

type SayHello interface {
	SayHello(context.Context, *proto.SayHelloRequest, *proto.SayHelloResponse) error
}

type UserInfo interface {
	UserInfo(context.Context, *proto.UserInfoRequest, *proto.UserInfoResponse) error
}

type AddUser interface {
	AddUser(context.Context, *proto.AddUserRequest, *proto.AddUserResponse) error
}

type UpdateUser interface {
	UpdateUser(context.Context, *proto.UpdateUserRequest, *proto.UpdateUserResponse) error
}

type RedisSet interface {
	RedisSet(context.Context, *proto.RedisSetRequest, *proto.RedisSetResponse) error
}

type RedisGet interface {
	RedisGet(context.Context, *proto.RedisGetRequest, *proto.RedisGetResponse) error
}

type MysqlInsert interface {
	MysqlInsert(context.Context, *proto.MysqlInsertRequest, *proto.MysqlInsertResponse) error
}

type MysqlSelect interface {
	MysqlSelect(context.Context, *proto.MysqlSelectRequest, *proto.MysqlSelectResponse) error
}

type MysqlDelete interface {
	MysqlDelete(context.Context, *proto.MysqlDeleteRequest, *proto.MysqlDeleteResponse) error
}

type serviceBridge struct {
	SayHelloImpl    SayHello
	UserInfoImpl    UserInfo
	AddUserImpl    AddUser
	UpdateUserImpl    UpdateUser
	RedisSetImpl    RedisSet
	RedisGetImpl    RedisGet
	MysqlInsertImpl    MysqlInsert
	MysqlSelectImpl    MysqlSelect
	MysqlDeleteImpl    MysqlDelete
}

func NewServiceBridge() *serviceBridge {
	return new(serviceBridge)
}

func (s *serviceBridge) SayHello(ctx context.Context, req *proto.SayHelloRequest, resp *proto.SayHelloResponse) error {
	return s.SayHelloImpl.SayHello(ctx, req, resp)
}

func (s *serviceBridge) UserInfo(ctx context.Context, req *proto.UserInfoRequest, resp *proto.UserInfoResponse) error {
	return s.UserInfoImpl.UserInfo(ctx, req, resp)
}

func (s *serviceBridge) AddUser(ctx context.Context, req *proto.AddUserRequest, resp *proto.AddUserResponse) error {
	return s.AddUserImpl.AddUser(ctx, req, resp)
}

func (s *serviceBridge) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest, resp *proto.UpdateUserResponse) error {
	return s.UpdateUserImpl.UpdateUser(ctx, req, resp)
}

func (s *serviceBridge) RedisSet(ctx context.Context, req *proto.RedisSetRequest, resp *proto.RedisSetResponse) error {
	return s.RedisSetImpl.RedisSet(ctx, req, resp)
}

func (s *serviceBridge) RedisGet(ctx context.Context, req *proto.RedisGetRequest, resp *proto.RedisGetResponse) error {
	return s.RedisGetImpl.RedisGet(ctx, req, resp)
}

func (s *serviceBridge) MysqlInsert(ctx context.Context, req *proto.MysqlInsertRequest, resp *proto.MysqlInsertResponse) error {
	return s.MysqlInsertImpl.MysqlInsert(ctx, req, resp)
}

func (s *serviceBridge) MysqlSelect(ctx context.Context, req *proto.MysqlSelectRequest, resp *proto.MysqlSelectResponse) error {
	return s.MysqlSelectImpl.MysqlSelect(ctx, req, resp)
}

func (s *serviceBridge) MysqlDelete(ctx context.Context, req *proto.MysqlDeleteRequest, resp *proto.MysqlDeleteResponse) error {
	return s.MysqlDeleteImpl.MysqlDelete(ctx, req, resp)
}

