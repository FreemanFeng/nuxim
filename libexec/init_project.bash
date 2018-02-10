#!/bin/bash

PROJECT=$1

if [ -z "$PROJECT" ]; then
    PROJECT=nuxim
fi

SRTPATH=$(cd "$(dirname "$0")"; pwd)

################################################
# 需安装go并设置环境变量 https://golang.org/dl/
################################################
export GOROOT=/usr/lib/go-1.9
export GOPATH=$SRTPATH/../

mkdir -p $SRTPATH/../data/$PROJECT
mkdir -p $SRTPATH/../conf/$PROJECT
mkdir -p $SRTPATH/../pids/$PROJECT
mkdir -p $SRTPATH/../output/$PROJECT
mkdir -p $SRTPATH/../logs/$PROJECT

rm -rf  $SRTPATH/../src/github.com/pquerna/ffjson
go get -u github.com/pquerna/ffjson
