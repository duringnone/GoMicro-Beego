# Gateway-网关层

微服务网关层(web层), 对外提供http服务的API接口



### 1. 目录简介

````
└─gateway            // web网关层目录
    ├─conf          // 配置目录
    |                 
    ├─controllers    // 控制器目录 [定义接口,路由,...]
    | 
    ├─models		// model层 [校验request参数/处理response值]
    | 
    ├─routers		// beego路由注册
    |                       
    ├─rpc			// rpc调用目录 [负责web服务->rpc服务 之间通信]
    | 
    ├─swagger		// swagger在线文档 目录
    |                     
    ├─vendor 		// go mod依赖包存放目录 [go mod vendor //将pkg中的依赖包导入当前./vendor目录中]
    |                      
    └─.gitignore       // git提交忽略文件 [保存不提交的文件/目录]
    │
    ├─go.mod           // go mod包依赖管理文件
    │
    ├─main.go          // 网关层/beego入口文件 [启动web服务]
    │
    └─README.md        // 简介


// ************** 详细目录 ************** //

└─gateway            // web网关层目录
    └─main.go          // 网关层/beego入口文件 [启动web服务]
    │
    └─go.mod           // go mod包依赖管理文件
    │
    └─.gitignore       // git提交忽略文件 [保存不提交的文件/目录]
    │
    └─README.md        // 简介
    │
    └─conf            // 配置目录
    |   ├─app.conf   	// beego配置文件
    |   │
    |   └─errMsg.conf   // 可增加自定义配置文件
    |   │
    |   └─...   		//其他配置文件 
    |            
    |                 
    └─controllers      // 控制器目录 [定义接口,路由,...]
    |   ├─Mail.go        // 邮件管理控制器 [邮件业务api接口] 
    |   │
    |   ├─Pprof.go     	// pprof分析控制器 [pprof性能分析]
    |   │      
    |   └─....           // 其他业务控制器
    |     
    | 
    └─ models			// model层 [校验request参数/处理response值]
    |    ├─const_error.go	// 错误常量 [错误码/错误信息]
    |    │
    |    ├─dao.go			// 初始化资源 [如:初始化全局日志号,...]
    |    │
    |    └─Mail.go			// 对应Mail控制器
    |    │
    |    └─MailConfig.go	// 对应MailConfig控制器
    |                    
    | 
     └─ routers			// beego路由注册
    |    ├─router.go 		// [批量]注册路由/注册路由命名空间
    |    │
    |    └─commentsRouter_controllers.go	// 存放路由信息 
    |            
    |                       
    └─ rpc			// rpc调用目录 [负责web服务->rpc服务 之间通信]
    |    ├─Common.go	// rpc共用部分
    |    │
    |    ├─Mail.go		// 邮件发送rpc通信
    |    │
    |    └─MailConfig.go // 邮件模板rpc通信
    |            
    | 
    └─ swagger	// swagger在线文档 目录
    |    │
    |    └─...
    |                     
    └─ vendor 	// go mod依赖包存放目录 [go mod vendor // 将pkg中的依赖包导入当前./vendor目录中]
    |    │
    |    └─...		// go.mod/go.sum文件中的各个依赖包目录
    |            
    |                      
    └─goods        // 其他...       




````



### 简介

```
网关层(即web层,后续统一称为网关层)gateway简介

1. 网关层采用beego的目的: 
	复用beego的路由管理,复用beego中的swagger在线文档管理,实时生成API文档,提高开发效率; 
	**注意: 若需使用swagger,则需安装bee工具,并在GOPATH路径下执行: bee run -gendoc=true -downdoc=true; 非GOPATH路径下会报错,亲测
	
2.原生beego 和 此微服务demo的gateway 区别:
	1) 原生beego为web框架,完成 控制器-->model业务逻辑处理-->数据库操作; 而 gateway 中只负责控制器部分+request参数校验(可做可不做),业务逻辑和数据库CURD均在rpc服务中进行
	2) gateway中只负责 管理路由+参数校验, 邮件管理逻辑,在srvmail中完成,gateway和srvmail通过rpc通信
	3) 本demo采用 Go-Micro-v2.*,默认服务注册采用mdns (解决内网动态IP寻址问题),故无需手动注册服务(区别于etcd服务注册/发现)




```

