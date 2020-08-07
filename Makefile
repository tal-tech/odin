
.PHONY: build

SERVICE := odin
MAIN := cmd/odin
#REGISTRY := etcd
REGISTRY := zookeeper
AUTHOR := $(shell git log --pretty=format:"%an"|head -n 1)
VERSION := $(shell git rev-list HEAD | head -1)
BUILD_INFO := $(shell git log --pretty=format:"%s" | head -1)
BUILD_DATE := $(shell date +%Y-%m-%d\ %H:%M:%S)
CUR_PWD := $(shell pwd)
export GO111MODULE=on
export GOPROXY=http://goproxy.xesv5.com

LD_FLAGS='-X "$(SERVICE)/version.TAG=$(TAG)" -X "$(SERVICE)/version.VERSION=$(VERSION)" -X "$(SERVICE)/version.AUTHOR=$(AUTHOR)" -X "$(SERVICE)/version.BUILD_INFO=$(BUILD_INFO)" -X "$(SERVICE)/version.BUILD_DATE=$(BUILD_DATE)"'

default: build

build:
	go build -tags "$(REGISTRY)" -ldflags $(LD_FLAGS) -gcflags "-N" -i -o ./bin/$(SERVICE) ./$(MAIN)
dev: build
	cp $(CUR_PWD)/conf/conf_dev.ini $(CUR_PWD)/conf/conf.ini && ./bin/$(SERVICE) -v=true
clean:
	rm bin/*
