#!/bin/sh
#
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -o nomad-binary ./...
mkdir -p /tmp/bin
pwd
ls -al
tar -czvf /tmp/bin/nomad-binary.tar --transform='s|.*/||'  /tmp/bin/nomad-binary 
ls -l /tmp/bin