# go-rpc-gateway
基于Go语言的RPC+HTTP+Gateway+Swagger

![](http://pic.w-blog.cn/800FD94A-504B-4AA1-9992-C98962E84287.png)

grpc-gateway是protoc的一个插件 。它读取Grpc服务定义，并生成反向代理服务器，将RESTful JSON API请求转换为Grpc的方式调用。主要是根据 google.api.http定义中思想完成的，以下为grpc-gateway结构图： ￼ 

![](http://pic.w-blog.cn/8DAABBE9-BA58-446A-8BF2-82CB0DB754A1.png)

一、环境准备

go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

go get -u github.com/golang/protobuf/protoc-gen-go

二、文件介绍

