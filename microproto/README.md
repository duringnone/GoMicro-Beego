# microproto
common proto for GoMicro-Beego





### 目录简介

```

    └─microproto      // rpc通信协议层
    |    │
    |    ├─mail         // 邮件rpc服务定义 [request/response/method]
    |    │ 	├─mail.pb.go         // rpc参数配置 [pb版本]
    │    |	|                 
    |    |	├─mail.pb.micro.go    // rpc参数配置 [go-micro版本]
    |    |	| 
    |    |	└─mail.proto		// 定义rpc通信的request/response格式,和rpc方法名
    |    │
    |    └─select.sh    // 编译./proto生成.pb,.micro.pb文件
    |
    | 
	└─... 

```



### 笔记

```
1. 前提: 必须先安装protoc工具, 用于生成rpc通信的内容 (requet/response格式,method)

2. 流程: 
	1) 定义mail.proto协议文件(定义rpc的request/response参数格式,method)
	2) 执行 ./select.sh 生成mail.pb.go , mail.pb.micro.go
	
3. select.sh 文件
	该文件中有2行命令:
		# 执行编译,.proto文件生成micro/go格式文件
		protoc --proto_path=. --go_out=. --micro_out=. mail/mail.proto

		# 零值可正常返回,若需忽略零值,则注释下面命令
		ls mail/*.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}' // 	protoc编译mail.proto时,返回值设置属性 omitempty; omitempty:默认会过滤返回值中的零值; 若需返回零值,则用此命令手动过滤掉mail.pb.go中的omitempty属性



```

