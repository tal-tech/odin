package service

import (
	"context"
	"odin/proto"
	rpcxplugin "github.com/tal-tech/odinPlugin"
	"github.com/tal-tech/odinPlugin/wrap"
	"odin/app/serviceInterface"
)

//rpcx服务注册类型
type Odin struct {
	*rpcxplugin.RpcxPlugin
	service serviceInterface.Service
}

//传入实现Service接口的类型
func NewOdin(service serviceInterface.Service) *Odin {
	this := new(Odin)
	this.service = service
	return this
}

//服务方法调用入口，先通过Wrapcall执行plugin方法，后调用service实现业务方法
func (this *Odin) SayHello(ctx context.Context, req *proto.SayHelloRequest, resp *proto.SayHelloResponse) error {
	fn := func(w *wrap.Wrap) error {
		ctx := w.GetCtx()
		err := this.service.SayHello(ctx, req, resp)
		return err
	}
	return this.Wrapcall(ctx, "SayHello", fn)
}

//服务方法调用入口，先通过Wrapcall执行plugin方法，后调用service实现业务方法
func (this *Odin) UserInfo(ctx context.Context, req *proto.UserInfoRequest, resp *proto.UserInfoResponse) error {
	fn := func(w *wrap.Wrap) error {
		ctx := w.GetCtx()
		err := this.service.UserInfo(ctx, req, resp)
		return err
	}
	return this.Wrapcall(ctx, "UserInfo", fn)
}

//服务方法调用入口，先通过Wrapcall执行plugin方法，后调用service实现业务方法
func (this *Odin) AddUser(ctx context.Context, req *proto.AddUserRequest, resp *proto.AddUserResponse) error {
	fn := func(w *wrap.Wrap) error {
		ctx := w.GetCtx()
		err := this.service.AddUser(ctx, req, resp)
		return err
	}
	return this.Wrapcall(ctx, "AddUser", fn)
}

//服务方法调用入口，先通过Wrapcall执行plugin方法，后调用service实现业务方法
func (this *Odin) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest, resp *proto.UpdateUserResponse) error {
	fn := func(w *wrap.Wrap) error {
		ctx := w.GetCtx()
		err := this.service.UpdateUser(ctx, req, resp)
		return err
	}
	return this.Wrapcall(ctx, "UpdateUser", fn)
}

//服务方法调用入口，先通过Wrapcall执行plugin方法，后调用service实现业务方法
func (this *Odin) RedisSet(ctx context.Context, req *proto.RedisSetRequest, resp *proto.RedisSetResponse) error {
	fn := func(w *wrap.Wrap) error {
		ctx := w.GetCtx()
		err := this.service.RedisSet(ctx, req, resp)
		return err
	}
	return this.Wrapcall(ctx, "RedisSet", fn)
}

//服务方法调用入口，先通过Wrapcall执行plugin方法，后调用service实现业务方法
func (this *Odin) RedisGet(ctx context.Context, req *proto.RedisGetRequest, resp *proto.RedisGetResponse) error {
	fn := func(w *wrap.Wrap) error {
		ctx := w.GetCtx()
		err := this.service.RedisGet(ctx, req, resp)
		return err
	}
	return this.Wrapcall(ctx, "RedisGet", fn)
}

//服务方法调用入口，先通过Wrapcall执行plugin方法，后调用service实现业务方法
func (this *Odin) MysqlInsert(ctx context.Context, req *proto.MysqlInsertRequest, resp *proto.MysqlInsertResponse) error {
	fn := func(w *wrap.Wrap) error {
		ctx := w.GetCtx()
		err := this.service.MysqlInsert(ctx, req, resp)
		return err
	}
	return this.Wrapcall(ctx, "MysqlInsert", fn)
}

//服务方法调用入口，先通过Wrapcall执行plugin方法，后调用service实现业务方法
func (this *Odin) MysqlSelect(ctx context.Context, req *proto.MysqlSelectRequest, resp *proto.MysqlSelectResponse) error {
	fn := func(w *wrap.Wrap) error {
		ctx := w.GetCtx()
		err := this.service.MysqlSelect(ctx, req, resp)
		return err
	}
	return this.Wrapcall(ctx, "MysqlSelect", fn)
}

//服务方法调用入口，先通过Wrapcall执行plugin方法，后调用service实现业务方法
func (this *Odin) MysqlDelete(ctx context.Context, req *proto.MysqlDeleteRequest, resp *proto.MysqlDeleteResponse) error {
	fn := func(w *wrap.Wrap) error {
		ctx := w.GetCtx()
		err := this.service.MysqlDelete(ctx, req, resp)
		return err
	}
	return this.Wrapcall(ctx, "MysqlDelete", fn)
}

