#export GOPROXY=https://goproxy.io
#export GOPROXY=https://goproxy.cn
#export GOPROXY:=https://proxy.golang.com.cn
#export GOPROXY:=https://mirrors.aliyun.com/goproxy
export GOPROXY:=https://repo.huaweicloud.com/repository/goproxy
#export GONOPROXY:=code.x.com/*,github.com/cnzhangquan/*,github.com/timectrl/*
#export GONOSUMDB:=*
export GO111MODULE:=on

ifdef GOPATH
export PATH:=$(GOPATH)/bin:$(PATH)
else
export PATH:=$(HOME)/go/bin:$(PATH)
endif

C_STATIC_LINK_FLAGS:=-ldflags="-extldflags=-static"
OS:=$(shell uname -s)
ifeq ($(OS),Darwin)
	C_STATIC_LINK_FLAGS:=
endif


PROGRAM:=$(patsubst %.go,%,$(wildcard *.go))
PROGRAM_SRC:=$(wildcard *.go)
SUB_MODULE_SRC:=$(shell find . -mindepth 2 -type f -name '*.go' |grep -v '^./vendor' |grep -v '^/docs')

ALL_SRC:=$(PROGRAM_SRC) $(SUB_MODULE_SRC) $(wildcard staticHandler/*) $(wildcard SQL/InitSQL/*.sql) $(wildcard SQL/QuerySQL/*.sql) Makefile




all: build

build: $(patsubst %,bin.%,$(PROGRAM))


# when using go mod, this config is NOT need. this just for doc
# and mod meta: <meta name="go-import" content="root-path vcs repo-url">
config:
	git config --global url."ssh://git@github.com:22/cnzhangquan/".insteadOf "https://github.com/cnzhangquan/"
	git config --global url."ssh://git@github.com:22/cnzhangquan/".insteadOf  "https//github.com/cnzhangquan/"
	git config --global url."ssh://git@github.com:22/timectrl/".insteadOf    "https://github.com/timectrl/"
	git config --global url."ssh://git@github.com:22/timectrl/".insteadOf     "http://github.com/timectrl/"
	git config --global url."ssh://git@code.x.com:10022/".insteadOf   "https://code.x.com/"
	git config --global url."ssh://git@code.x.com:10022/".insteadOf    "http://code.x.com/"

bin.%: %.go go.mod $(ALL_SRC) sqlbuild apidocs
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
	@#find  ./ -type f -name "*.go"|grep -v -E '^./vendor/'|grep -v '/GOPATH/'|grep -v '/SQL/'|xargs -I {} go fmt -mod=readonly {}






sql: sqlbuild

sqlbuild:
	@make -C SQL/InitSQL  build
	@make -C SQL/QuerySQL build

sqlclean:
	@make -C SQL/InitSQL  clean
	@make -C SQL/QuerySQL clean

apidocs: docs/.timestamp
docs/.timestamp: $(ALL_SRC)
	@echo PATH=$(PATH)
	which swag &>/dev/null || go install github.com/swaggo/swag/cmd/swag@latest
	swag init $(patsubst %,-g %,$(wildcard *.go))
	@touch docs/.timestamp

go.mod: project.mod apidocs
	cp -f $< $@
	go mod tidy

vendor: go.mod
	@echo process vendor
	@go mod vendor

clean:
	rm -fr bin.* docs

distclean: clean sqlclean
	rm -fr go.mod go.sum

realclean: distclean
	rm -fr vendor
	@#go clean -modcache
	@#[ -d GOPATH ] && chmod -R 777 GOPATH || true
	@#rm -fr GOPATH



print-%:
	@echo '$*=$($*)'


.PHONY: all build fmt vendor sqlcleen clean distclean realclean

.SECONDARY:

