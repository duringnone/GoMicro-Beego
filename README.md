# 微服务demo

因历史原因,业务采用Beego为web框架,为兼容历史版本;结合beego+gomicro开发,但寻无果,故自行兼容;

本demo旨在助力尽快落地beego+go-micro,尽量精简,如有异议,欢迎交流,交流群长期有效,会定时更新,群二维码如下:
    https://github.com/duringnone/GoMicro-Beego/blob/master/README-img/chatGroup.png



### 各个工具版本

```
Golang 1.13.*		// Go1.13最佳,建议至少Go1.12(1.12开始支持go mod),GO111MODULE="ON/OFF/AUTO" // 是否开启go mod,go.1.12为ON,go.1.13为auto,go.1.14为off;故go.1.13可自行切换go mod/gopath

Go-Micro.v2.*		// go mod自动拉取,若需使用v1,需自行指定; V2采用mdns为默认服务注册/发现,后续版本可能改为etcd


// ---- 受以下工具影响不大 -----
libprotoc 3.11.4	// protoc工具,此版本默认拉取go-micro.v2.*版本
Beego.v1.12.1		// beego web框架 
Bee 1.11.1			// bee工具
MySQL v.5.7.*		// 数据库
Linux v.3.10.0-957.el7.x86_64	// 服务器CentOS.7.0.*版本


```



### GoMicro架构图

 https://github.com/duringnone/GoMicro-Beego/blob/master/README-img/microService_struct.png



### 真实/demo架构图

https://github.com/duringnone/GoMicro-Beego/blob/master/README-img/demo_struct.png



### 架构简介

```

 └─GoMicro-Beego            // web网关层目录
    |    ├─gateway		// 网关层	[路由管理,参数校验]
    │    |                 
    |    ├─srvmail		// rpc服务层	[真正处理业务逻辑,数据库CURD,Redis,...]
    |    | 
    |    ├─microproto	// 定义RPC通信 request/response参数格式,rpc方法method
    |    │
    |    └─microbase	// 公共方法模块 [可有可无]
	│
    └─...

```



### 笔记

```
一. 真实企业架构中,完整微服务架构分为 4 个部分,如下:
	gateway		// 网关层	[路由管理,参数校验]
	srvmail		// rpc服务层	[真正处理业务逻辑,数据库CURD,Redis,...]
	microproto	// 定义RPC通信 request/response参数格式,rpc方法method
	microbase	// 公共方法模块 [可有可无]
	
	
	分析: 一般 网关层(如:gateway),rpc服务(如:srvmail)各有多个,而microproto/microbase一般共用一个,gateway也可一个; 
	我们真实场景中,只有rpc服务层是多个,gateway,microproto,microbase均只有一个,目的是方便管理,同时又能实现微服务的隔离,实现高效开发; 也有的公司会在web前面做一层网关(如:openorsty限流+服务治理:熔断/降级/...)
		
```



### 二. microproto简介

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



### 三. 私有仓库配置

```

	1. 上述中,提到我使用了4个项目
		gateway,srvmail,microbase,microproto
		
		在企业实际场景中,microbase,microproto一般会放在远程仓库,并且是私有的,下面谈谈如何配置,思路如下:
		1)若microbase和microproto共有的(即外部可访问),直接import即可
		2) 若 microbase 和 microproto是私有的
			A) 配置 环境变量			
            [linux环境]
            set GOPRIVATE="git.test.com,git.demo.com" // 配置私有仓库域名,多个用英文逗号隔开
            set GOSUM="off"	// 关闭包地址校验
            set GOPROXY="https://goproxy.cn,direct"	// 大陆地区使用七牛云的 goproxy.cn 代理会更快(尤其是涉及墙的包,但墙外包goproxy.cn可能有一定延迟,笔者最长遇到10分钟左右延迟)
		
            [windows/MacOS环境]
            unset GOPRIVATE	// 删除原配置
             unset GOSUM	// 删除原配置
            unset GOPROXY	// 删除原配置
            set GOPRIVATE="git.test.com,git.demo.com" // 配置私有仓库域名,多个用英文逗号隔开
             set GOSUM="off"	// 关闭包地址校验
            set GOPROXY="https://goproxy.cn,direct"	// 设置代理
            
            [windows环境]
            go env -w  GOPRIVATE="git.demo.com,git.test.com" // 配置私有仓库域名（多个）
            go env -w GOSUM="off"	// 关闭包地址校验
            go env -w GOPROXY="https://goproxy.cn,direct"	// 设置代理
            
		3) 配置私有仓库时,若仓库只能https登录拉取代码,会有一个问题: 每次git pull都需输入用户名,密码;而线上一般为自动操作,无法手动输入密码; 此时则需更改git全局配置,设置记住密码,只需第一次git pull输入一次密码后就可;操作如下:
            A) vim /root/.gitconfig	// 打开服务器git全局配置,linux一般是/root/.gitconfig
            B) 增加2行,作用是记住第一次手动输入的密码
                [credential]
                helper = store
            C) 若需指定访问仓库的网络协议,可在 /root/.gitconfig中增加以下代码:
			[url "git@git.test.com:"]
    		insteadOf = https://git.test.com
    	
```



### 四. 注意项

```

	1. 服务注册: gateway 和 srvmail 依靠rpc通信,因本demo采用go-micro.v2.*; go-micro.v2.*默认采用mdns进行服务注册/发现,故无需手动注册服务名,若需使用etcd注册服务,则需手动操作;虽实践中代码有etcd部分,但目前本demo暂未就etcd详细展开,详情自行google,会更详细

	2. 服务追踪:服务治理一部分,目前通过ELK(ES+Logstash+Kibana)做日志上报,同每个请求产生唯一日志编号flowNum,并将flowNum贯穿整个request链路,实现服务追踪; 追踪链路实现详情见代码(gateway+srvmail);ELK部分需读者自行安装实现,见谅;

	3. 权限/登录校验: 通过在gateway网关层实现登录(或实现登录服务),Redis存储登录信息,所有rpc服务全局可共享,亦可实现SSO单点登录; 为尽量精简代码篇幅,功能暂未在代码中实现,主要谈下实现思路;  
	
	4. go mod: 微服务 涉及的依赖包较多,建议使用go mod,笔者对比过go mod 和传统GOPATH调用的实现,有 如下区别:
		1) 微服务依赖包过多,GOPATH只能手动更新包,而微服务依赖存在间接依赖(A依赖B,B依赖C,C依赖D,...),此时go mod强大显现
		2) 微服务可自动获取最新版本依赖包,配合go-micro初始化时,自动加载
		3) go mod可自动拉取指定版本的依赖包,故无需开发者手动下载,打包,上线依赖包,提供开发流程效率,也节省了项目的存储空间
		4) 若需了解 项目相关依赖包,可到 GOPATH/pkg下查看; 也可执行 go mod vendor,将依赖包导入到当前项目根目录vendor目录下
		5) go mod更多详情,可查看煎鱼的Go Module入门教学,很详细很受用: https://eddycjy.com/posts/go/go-moduels/2020-02-28-go-modules/
		
```



### 更新内容-20210226

```
更新内容：
	1）交流群码
	2）go mod拉取问题：
		A）需求：
				go mod拉取非master分支代码 （go mod 可以拉取除master分支的最新版本git代码么,比如dev分支最新版本）
		B）我的场景:
				现在有2个服务user,login,user会调用login; user和login都在dev分支迭代开发, 基于go mod, 在dev环境下, 如何实现的user调用login
		C）go.mod文件内容格式：		
        module github.com/duringnone/microproto

        go 1.13

        require (
					github.com/duringnone/commproto v0.0.0-20201230113637-4a1965d03eaa
          github.com/micro/go-micro/v2 v2.9.0 // indirect
        )
	
			分析格式: 
					github.com/duringnone/commproto v0.0.0-20201230113637-4a1965d03eaa		// 4a1965d03eaa 是git仓库的commitId的前12位(可在git仓库查看，或本地git仓库执行"git log")
					
			D) 解决方案：
				a) 方案1: 发布到master分支，并打tag，go.mod中应用tag作为版本号；如：”github.com/micro/go-micro/v2 v2.9.0“ 中的“v2.9.0”对应master的tag,tag=v2.9.0
				b) 方案2: 引入指定branch分支名，或指定commitID；步骤如下：
						1) go clean -modcache // 清除本地go mod缓存
						2） 修改引用包的版本号：【2种任选一种】
		           A） github.com/duringnone/commproto dev // 第一种：指定分支最新版本代码
		           B） github.com/duringnone/commproto 1483a79ff755e6b4857915bae9b3b5e656a21c05 // 第二种：指定代码版本commitID
						3）go run main.go // 执行/编译入口文件
				
				
			E）注意： 
					1）go mod默认拉取master分支最新版本代码，版本号默认格式如：“v0.0.0-20201230113637-4a1965d03eaa”
					2）引用master分支的tag作为版本号
					3）引入dev分支最新版本代码，方案
					4）引入commitID； （因为git中commitID是全局唯一的，跨分支全局）
					5) go mod / go get 底层其实都是基于git 命令的封装



```



### 完整Micro服务结构目录 [更详细的在具体项目中]

```
// ---- 详细目录结构 --------
└─Beego-GoMicro	// Beego+Go-Micro.v2.* demo
    |
    └─gateway            // web网关层目录
    |    ├─conf          // 配置目录
    │    |                 
    |    ├─controllers   // 控制器目录 [定义接口,路由,...]
    |    | 
    |    ├─models		// model层 [校验request参数/处理response值]
    |    | 
    |    ├─routers		// beego路由注册
    |    |                       
    |    ├─rpc			// rpc调用目录 [负责web服务->rpc服务 之间通信]
    |    | 
    |    ├─swagger		// swagger在线文档 目录
    |    |                     
    |    ├─vendor 		// go mod依赖包存放目录 [go mod vendor //将pkg中的依赖包导入当前./vendor目录中]
    |    │
    |    └─main.go      // 网关层/beego入口文件 [启动web服务]
    |
    | 	
    └─srvmail         // 服务层目录
    |    ├─conf       // 配置目录
    │    |                 
    |    ├─handler    // rpc控制器目录 [定义接口,路由,...]
    |    | 
    |    ├─models		// model层 [业务逻辑处理+数据库CURD]
    |    │
    |    └─main.go       // 服务层入口文件 [启动rpc服务]
    | 
    | 
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
    └─microbase           // 共用方法
    |    ├─base.go          // 环境变量,全局变量初始化
    │    |                 
    |    ├─db.go    		// MySQL连接池,初始化,CURD操作封装
    |    | 
    |    ├─service.go		// 字符串/数组/类型转换/...处理方法
    |    |                     
    |    ├─third.go 		// 第三方服务 [邮件发送服务]
    |    │
    |    └─validate.go      // 验证器
    | 	
    └─tb_emails.sql		// 邮件服务demo的建表SQL
    | 	
    └─test.com.conf		// test.com 网关层的nginx vhost的.conf配置文件
    
    
```

