package serviceImpl

import (
	"context"
	"fmt"

	"odin/proto"
)

type HelloService struct {
}

func NewHelloService() *HelloService {
	return new(HelloService)
}

func (this *HelloService) SayHello(ctx context.Context, req *proto.SayHelloRequest, resp *proto.SayHelloResponse) error {
	reply := fmt.Sprintf("i'm hello service,recv greeting:%s", req.Greeting)
	resp.Reply = reply
	return nil
}
