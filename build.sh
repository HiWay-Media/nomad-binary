#!/bin/sh
#
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -o nomad-binary ./...
mkdir -p /tmp/bin
pwd
ls -al
tar -czvf /tmp/bin/nomad-binary.tar.gz --transform='s|.*/||'  ./nomad-binary 
#tar --ignore-zeros -xf
ls -l /tmp/bin