package main

//go:generate goctl rpc protoc ./proto/captcha.proto --go_out=./client/captcha/pb --go-grpc_out=./client/captcha/pb --zrpc_out=./client/captcha --client=true
