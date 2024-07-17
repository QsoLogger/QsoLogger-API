#export GOPROXY=https://goproxy.io
#export GOPROXY=https://goproxy.cn
#export GOPROXY=https://mirrors.aliyun.com/goproxy
export GOPROXY=https://repo.huaweicloud.com/repository/goproxy
export GO111MODULE=on

C_STATIC_LINK_FLAGS:=-ldflags="-extldflags=-static"
OS:=$(shell uname -s)
ifeq ($(OS),Darwin)
	C_STATIC_LINK_FLAGS:=
endif

PROGRAM:=$(patsubst %.go,%,$(wildcard *.go))
PROGRAM_SRC:=$(wildcard *.go)
SUB_MODULE_SRC:=$(shell find . -mindepth 2 -type f -name '*.go' |grep -v '^./vendor')

ALL_SRC:=$(PROGRAM_SRC) $(SUB_MODULE_SRC) $(wildcard staticHandler/*) $(wildcard SQL/InitSQL/*.sql) $(wildcard SQL/QuerySQL/*.sql) Makefile




all: build

build: $(patsubst %,bin.%,$(PROGRAM))

bin.%: %.go go.mod $(ALL_SRC) sqlbuild
	@make -C SQL/InitSQL build
	@make -C SQL/QuerySQL build
	@if [ -d vendor ]; \
	then \
		echo build with local vendor: $@ ; \
		CGO_ENABLED=1 go build -v $(C_STATIC_LINK_FLAGS) -tags sqlite_omit_load_extension -modcacherw -trimpath -mod=vendor -o $@ $< ; \
	else \
		echo build with upstream: $@; \
		CGO_ENABLED=1 go build -v $(C_STATIC_LINK_FLAGS) -tags sqlite_omit_load_extension -modcacherw -trimpath -mod=mod    -o $@ $< ; \
	fi



fmt: go.mod
	@find  ./ -type f -name "*.go"|grep -v -E '^./vendor/'|grep -v '/GOPATH/'|grep -v '/SQL/'|xargs -I {} go fmt {}





go.mod: project.mod
	cp -f $< $@
	go mod tidy


sql: sqlbuild

sqlbuild:
	@make -C SQL/InitSQL  build
	@make -C SQL/QuerySQL build

sqlclean:
	@make -C SQL/InitSQL  clean
	@make -C SQL/QuerySQL clean


vendor: go.mod
	@echo process vendor
	@go mod vendor

clean:
	rm -fr bin.*

distclean: clean sqlclean
	rm -fr go.mod go.sum

realclean: distclean
	rm -fr vendor
	@#go clean -modcache





.PHONY: all build fmt vendor sqlcleen clean distclean realclean

.SECONDARY:

