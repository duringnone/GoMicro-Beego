### 服务层

负责 业务逻辑+数据库CURD操作



### 目录简介

```
└─srvmail           // 服务层目录
    ├─conf          // 配置目录
    |                 
    ├─handler    // rpc控制器目录 [定义接口,路由,...]
    | 
    ├─models		// model层 [业务逻辑处理+数据库CURD]
    |                     
    ├─vendor 		// go mod依赖包存放目录 [go mod vendor //将pkg中的依赖包导入当前./vendor目录中]
    |                      
    ├─.gitignore       // git提交忽略文件 [保存不提交的文件/目录]
    │
    ├─go.mod           // go mod包依赖管理文件
    │
    ├─main.go          // 服务层入口文件 [启动rpc服务]
    │
    └─README.md        // 简介


// ************** 详细目录 ************** //

└─srvmail            // 服务目录
    │
    └─conf            // 配置目录
    |   ├─app.conf   	// 服务配置
    |   │
    |   └─...   		//其他配置文件 [可自定义]
    |            
    |                 
    └─handler      // rpc控制器目录 [定义rpc接口]
    |   ├─handler.go       // rpc服务公共加载入口
    |   │
    |   ├─Mail.go  		   // 邮件rpc服务入口
    |   │      
    |   └─MailConfig.go     // 邮件管理rpc服务入口 
    |     
    | 
    └─ models			// model层 [业务逻辑处理+数据库CURD]
    |    ├─Common.go		// 共用方法
    |    │
    |    ├─const_error.go	// 错误常量 [错误码/错误信息]
    |    │
    |    ├─dao.go			// 初始化资源 [如:初始化DB/Redis,rpc日志号,...]
    |    │
    |    └─Mail.go			// Mail业务逻辑
    |    │
    |    └─MailConfig.go	// MailConfig业务逻辑
    |            
    |                       
    └─ vendor 	// go mod依赖包存放目录 [go mod vendor // 将pkg中的依赖包导入当前./vendor目录中]
    |    │
    |    └─...		// go.mod/go.sum文件中的各个依赖包目录
    |
    |
    ├─.gitignore       // git提交忽略文件 [保存不提交的文件/目录]
    │
    ├─go.mod           // go mod包依赖管理文件
    │
    ├─main.go          // 服务层入口文件 [启动rpc服务]
    │
    └─README.md        // 简介     





```



### rpc服务层

```
1. rpc服务层 和 网关层(web层)一般rpc通信(rpc调用时函数级别的调用,比http调用高效)

2. rpc服务层和网关层 共同操作mail.proto中的内容(.proto定义API接口的request/response/method,即请求参数格式/响应返回值格式/rpc接口方法)


```

