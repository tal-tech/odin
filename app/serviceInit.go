package app

import (
	"odin/app/service"
	"odin/app/serviceImpl"
)

func NewService() *service.Odin {
	s := service.NewServiceBridge()
	HelloService := serviceImpl.NewHelloService()
	s.SayHelloImpl = HelloService
	UserService := serviceImpl.NewUserService()
	s.UserInfoImpl = UserService
	s.AddUserImpl = UserService
	s.UpdateUserImpl = UserService
	DemoService := serviceImpl.NewDemoService()
	s.RedisSetImpl = DemoService
	s.RedisGetImpl = DemoService
	s.MysqlInsertImpl = DemoService
	s.MysqlSelectImpl = DemoService
	s.MysqlDeleteImpl = DemoService
	return service.NewOdin(s)
}