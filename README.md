# Odin

## Introduction

Odin是基于go语言的rpc框架，框架除了致力于提供高性能的服务间调用能力外，也提供完善的服务治理功能、支持多种服务注册发现机制。为了业务方使用框架更加便捷，框架还具有配套管理工具自动生成代码，提高开发效率。

## Features

### 高性能
首先，框架底层基于rpcx框架，其是一个纯Go语言的rpc框架，与主流rpc框架进行性能对比，优势明显。其性能仅弱于Go原生rpc调用。

### 服务治理
odin框架可提供统一的服务注册管理，仅通过配置地址方式即可方便使用以及切换服务注册中心。在支持原生容错和负载均衡机制基
础上，开发插件系统，包括限流、断路器、打点统计、耗时报警等。

### 开发便捷
Odin结合配套的辅助工具rigger，可以直接生成框架模板，业务使用方只需定义对外提供接口，可自动生成服务代码，开发只需编写
业务逻辑。提供给其他服务的client代码，同样可命令生成，方便调用。

### 自定义支持
Odin框架目前已支持日志Trace跨服务传递，记录一次完整请求的所有记录，根据同一TraceID，查看全部链路。其他包括动态插件都
可自定义开发，只需最终在main注入即可。

## Quick Start

### Install
```
//进入开发目录$GOPATH/src
$ cd $GOPATH/src/

//Clone项目到开发目录
$ git clone git@github.com:tal-tech/odin.git
```

### Build
```
//Makefile可依需求自定义
make
```

### Run
```
//启动
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
//注册中心启用开关
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

## 框架共建
我们的目标是将Odin打造成一个高性能、高可靠、易用的微服务框架，欢迎大家共同参与，做更多的创新以及贡献，包括但不限于
以下内容：
* 支持更多的调用方式；
* 扩展中间件功能；
* 提供更灵活、更便捷的服务治理功能；

## 联系我们
issue: [https://github.com/tal-tech/odin/issues](https://github.com/tal-tech/odin/issues)  
