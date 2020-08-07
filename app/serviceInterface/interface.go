package serviceInterface

import (
	"context"

	"odin/proto"
)

type Service interface {
	//HelloService
	SayHello(ctx context.Context, req *proto.SayHelloRequest, resp *proto.SayHelloResponse) error

	//UserService
	UserInfo(ctx context.Context, req *proto.UserInfoRequest, resp *proto.UserInfoResponse) error
	AddUser(ctx context.Context, req *proto.AddUserRequest, resp *proto.AddUserResponse) error
	UpdateUser(ctx context.Context, req *proto.UpdateUserRequest, resp *proto.UpdateUserResponse) error

	//DemoService
	RedisSet(ctx context.Context, req *proto.RedisSetRequest, resp *proto.RedisSetResponse) error
	RedisGet(ctx context.Context, req *proto.RedisGetRequest, resp *proto.RedisGetResponse) error
	MysqlInsert(ctx context.Context, req *proto.MysqlInsertRequest, resp *proto.MysqlInsertResponse) error
	MysqlSelect(ctx context.Context, req *proto.MysqlSelectRequest, resp *proto.MysqlSelectResponse) error
	MysqlDelete(ctx context.Context, req *proto.MysqlDeleteRequest, resp *proto.MysqlDeleteResponse) error
}
