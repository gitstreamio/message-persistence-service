package main

import (
	"message-persistence-service/server"
)


//go:generate go get github.com/golang/mock/gomock
//go:generate go get github.com/golang/mock/mockgen
//go:generate $GOPATH/bin/mockgen -package common -source common/interfaces.go -destination common/mock_interfaces.go
func main() {
	server.Run()
}
