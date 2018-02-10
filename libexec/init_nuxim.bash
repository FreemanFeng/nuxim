#!/bin/bash

SRTPATH=$(cd "$(dirname "$0")"; pwd)

################################################
# 需安装go并设置环境变量 https://golang.org/dl/
# export GOROOT=/vobs/tools/go
# export GOPATH=$SRTPATH/../
################################################
export GOROOT=/usr/lib/go-1.9
export GOPATH=$SRTPATH/../

rm -rf  $SRTPATH/../src/github.com/pquerna/ffjson
go get -u github.com/pquerna/ffjson

# Build Caddy
rm -rf  $SRTPATH/../src/github.com/mholt/caddy/caddy
go get -u github.com/mholt/caddy/caddy

rm -rf  $SRTPATH/../src/github.com/caddyserver/builds
go get -u github.com/caddyserver/builds

cd $SRTPATH/../src/github.com/mholt/caddy/caddy
go run build.go
if [ -f ./caddy ]; then
    rsync -a caddy $SRTPATH/../../bin/
    rsync -a caddy $SRTPATH/../bin/
fi
cd $SRTPATH
