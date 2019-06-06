#!/bin/bash

set -ex

TAG=${1:-devel}

GOOS=windows GOARCH=amd64 go build -o binaries/hvinfo.exe hvinfo.go
go build -o binaries/hvinfo hvinfo.go

cp binaries/hvinfo binaries/hvinfo-${TAG}-linux-amd64
cp binaries/hvinfo.exe binaries/hvinfo-${TAG}-windows-amd64.exe
