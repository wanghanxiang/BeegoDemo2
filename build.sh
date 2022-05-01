#!/bin/bash
main="main.go"
appname="beegodemo"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${appname} ${main}