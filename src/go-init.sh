#!/bin/bash

go mod init raspberrypi.local/cronTrigger
go get github.com/eclipse/paho.mqtt.golang
go get github.com/gorilla/websocket
go get golang.org/x/net/proxy
go get github.com/robfig/cron/v3
