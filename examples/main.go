package main

import (
	"context"
	"fmt"
	"os"

	p "odin/proto"

	"github.com/tal-tech/xtools/addrutil"
	"github.com/tal-tech/xtools/confutil"
	"github.com/tal-tech/xtools/rpcxutil"

	"github.com/smallnest/rpcx/client"
)

func main() {
	/*
		basePath := confutil.GetConf("Registry", "basePath")
		registryConfig := confutil.GetConfArrayMap("Registry")
	*/
	option := client.DefaultOption

	var d client.ServiceDiscovery

	if confutil.GetConf("Registry", "status") == "on" {

		//d = rpcxutil.GetClientDiscovery(rpcxutil.GetServiceBasePath(), "Odin")
		wrapClient := rpcxutil.NewWrapClient(rpcxutil.GetServiceBasePath(), "Odin", client.Failtry, client.RandomSelect, option)

		//helloservice
		req := &p.SayHelloRequest{
			Greeting: "hello, i'm odin client",
		}
		resp := &p.SayHelloResponse{}
		err := wrapClient.WrapCall(context.Background(), "SayHello", req, resp)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to call SayHello: %v\n", err)
		}
		fmt.Fprintf(os.Stdout, "SayHello: %s\n", resp.Reply)

		//userservice
		reqAdd := &p.AddUserRequest{
			Name: "学而思",
			Age:  10,
			City: "beijing",
		}
		respAdd := &p.AddUserResponse{}
		err = wrapClient.WrapCall(context.Background(), "AddUser", reqAdd, respAdd)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to call AddUser: %v\n", err)
		}
		fmt.Fprintf(os.Stdout, "AddUser: %+v\n", respAdd)

		reqInfo := &p.UserInfoRequest{
			Id: respAdd.Id,
		}
		respInfo := &p.UserInfoResponse{}
		err = wrapClient.WrapCall(context.Background(), "UserInfo", reqInfo, respInfo)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to call UserInfo: %v\n", err)
		}
		fmt.Fprintf(os.Stdout, "UserInfo: %s\n", respInfo)

		reqUpdate := &p.UpdateUserRequest{
			Id:   respAdd.Id,
			Name: "网校",
			Age:  20,
			City: "beijing",
		}
		respUpdate := &p.UpdateUserResponse{}
		err = wrapClient.WrapCall(context.Background(), "UpdateUser", reqUpdate, respUpdate)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to call UpdateUser: %v\n", err)
		}
		fmt.Fprintf(os.Stdout, "UpdateUser: %s\n", respUpdate)

		err = wrapClient.WrapCall(context.Background(), "UserInfo", reqInfo, respInfo)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to call UserInfo: %v\n", err)
		}
		fmt.Fprintf(os.Stdout, "UserInfo: %s\n", respInfo)
	} else {

		addr, _ := addrutil.Extract("")

		serviceAddr := confutil.GetConf("Server", "network") + "@" + addr + ":" + confutil.GetConf("Server", "port")

		d = client.NewPeer2PeerDiscovery(serviceAddr, "")

		xclient := client.NewXClient("Odin", client.Failtry, client.RandomSelect, d, option)
		defer xclient.Close()

		//helloservice
		req := &p.SayHelloRequest{
			Greeting: "hello, i'm odin client",
		}
		resp := &p.SayHelloResponse{}
		err := xclient.Call(context.Background(), "SayHello", req, resp)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to call SayHello: %v\n", err)
		}
		fmt.Fprintf(os.Stdout, "SayHello: %s\n", resp.Reply)

		//userservice
		reqAdd := &p.AddUserRequest{
			Name: "学而思",
			Age:  10,
			City: "beijing",
		}
		respAdd := &p.AddUserResponse{}
		err = xclient.Call(context.Background(), "AddUser", reqAdd, respAdd)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to call AddUser: %v\n", err)
		}
		fmt.Fprintf(os.Stdout, "AddUser: %+v\n", respAdd)

		reqInfo := &p.UserInfoRequest{
			Id: respAdd.Id,
		}
		respInfo := &p.UserInfoResponse{}
		err = xclient.Call(context.Background(), "UserInfo", reqInfo, respInfo)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to call UserInfo: %v\n", err)
		}
		fmt.Fprintf(os.Stdout, "UserInfo: %s\n", respInfo)

		reqUpdate := &p.UpdateUserRequest{
			Id:   respAdd.Id,
			Name: "网校",
			Age:  20,
			City: "beijing",
		}
		respUpdate := &p.UpdateUserResponse{}
		err = xclient.Call(context.Background(), "UpdateUser", reqUpdate, respUpdate)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to call UpdateUser: %v\n", err)
		}
		fmt.Fprintf(os.Stdout, "UpdateUser: %s\n", respUpdate)

		err = xclient.Call(context.Background(), "UserInfo", reqInfo, respInfo)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to call UserInfo: %v\n", err)
		}
		fmt.Fprintf(os.Stdout, "UserInfo: %s\n", respInfo)
	}
}
