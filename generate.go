package main

//go:generate protoc --go_out=./client/captcha --go_opt=paths=source_relative --go-grpc_out=./client/captcha --go-grpc_opt=paths=source_relative -I=./proto ./proto/captcha.proto
