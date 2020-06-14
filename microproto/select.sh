#!/bin/bash


## 操作前提: 本地需安装 protoc 工具

## 问题: protoc操作过滤 go-micro的json返回值json忽略零值
## 解决方案
# 执行编译,.proto文件生成micro/go格式文件
`protoc --proto_path=. --go_out=. --micro_out=. mail/mail.proto`
# 零值可正常返回,若需忽略零值,则注释下面命令
` ls mail/*.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'`
