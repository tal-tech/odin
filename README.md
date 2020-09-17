# Odin

<p align="center">
 <a href="https://tal-tech.github.io/odin-doc/" target="_blank">
     <img src="https://xesftp.oss-cn-beijing.aliyuncs.com/oa/res/odin.jpg?raw=true"  alt="Odin Logo" align=center />
 </a> 
</p>

业内存在多款流行的微服务框架，包括Montan、rpcx、gRpc、Dubbo等

Odin框架是好未来在使用rpcx框架过程中不断完善、改进的，目前已在好未来多个部门使用，承载大量核心服务，经历暑期大班直播高峰考验

2019年 开始搭建框架主架构，并在此基础上不断完善、补充新功能

2020年8月 开源 欢迎大家一起参与共建



## Odin框架文档    

[Document](https://tal-tech.github.io/odin-doc/) 

[中文文档](https://www.yuque.com/tal-tech/odin) 

## Introduction

Odin是基于go语言的rpc框架，框架除了致力于提供高性能的服务间调用能力外，也提供完善的服务治理功能、支持多种服务注册发现机制。为了业务方使用框架更加便捷，框架还具有配套管理工具自动生成代码，提高开发效率。

## Features

### 高性能
首先，框架底层基于rpcx框架，其是一个纯Go语言的rpc框架，与主流rpc框架进行性能对比，优势明显。其性能仅弱于Go原生rpc调用。

### 服务治理
odin框架可提供统一的服务注册管理，仅通过配置地址方式即可方便使用以及切换服务注册中心。在支持原生容错和负载均衡机制基
础上，开发插件系统，包括限流、断路器、打点统计、耗时报警等。

### 开发便捷
Odin结合配套的辅助工具[rigger](http://github.com/tal-tech/rigger)，可以直接生成框架模板，业务使用方只需定义对外提供接口，可自动生成服务代码，开发只需编写
业务逻辑。提供给其他服务的client代码，同样可命令生成，方便调用。

### 自定义支持
Odin框架目前已支持日志Trace跨服务传递，记录一次完整请求的所有记录，根据同一TraceID，查看全部链路。其他包括动态插件都
可自定义开发，只需最终在main注入即可。

## Quick Start

### Install
```
rigger new micro rpcproject
正克隆到 '/winshare/go/src/rpcproject'...
rpcprojec项目已创建完成, 使用:
cd /winshare/go/src/rpcproject && rigger build 
开始你的微服务之旅！
```

### Build
```
//Makefile可依需求自定义
rigger build
```

### Run
```
//启动
rigger start
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
rigger example zookeeper
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
