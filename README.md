# Odin

## Introduction

 be referred to Odin is a Rpcx-based rpc framework.There are many rpc frameworks in the industry, such as Dubbo, Dubbox, motan, gRPC, thrift, etc. Among all the above frameworks, the rpcx framework has obvious advantages in both function and performance. The features of rpcx framework and benchmark result can be referred to be referred to [github.com/smallnest/rpcx](https://github.com/smallnest/rpcx) 

## Quick Start

### Install
```
//Recommended  $GOPATH/src  as your workspace
$ cd $GOPATH/src/

//Clone the framework to local
$ git clone git@github.com:tal-tech/odin.git
```

### Build
```
//You can customize the Makefile
make
```

### Run
```
//run in frontend
bin/odin
```

## Config
```
//conf/conf.ini
//listen port
[Server]
network=tcp
port=11900

//Register Addr
[Registry]
//use or not use registry
status=off 
//Registry Address
addrs=127.0.0.1:2181
basePath=/odin_demo
```

## Example
```
go run -tags 'zookeeper' examples/main.go  -c $GOPATH/src/odin/conf/conf.ini
//Output
SayHello: i'm hello service,recv greeting:hello, i'm odin client
AddUser: &{Id:3}
UserInfo: &{学而思 %!s(int=10) beijing}
UpdateUser: &{}
UserInfo: &{网校 %!s(int=20) beijing}
```
