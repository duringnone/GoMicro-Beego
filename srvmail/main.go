package main

import (
	"fmt"
	. "github.com/duringnone/microbase"
	"github.com/duringnone/microproto/mail"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"srvmail/handler"
	//"github.com/micro/go-micro/v2/registry/etcd"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 使用etcd注册
	//micReg := etcd.NewRegistry(registryOptions)
	// 初始化rpc服务配置
	srv := micro.NewService(
		micro.Name(ConfigInfo.String("MicroServiceApi::SRV_NAMESPACE")),   // 服务注册名,如: test.com.srv.mail
		micro.Address(":"+ConfigInfo.String("MicroServiceApi::SRV_PORT")), // 服务端口,如: 8821 ***注: micro.Address(":8821")
		//micro.Registry(micReg), 	// 指定etcd注册方式(***注: go-Micro-v2.*版本: 默认使用mdns,故mdns同一内网可自动寻址,无需手动注册)
		micro.Version("latest"), // 使用版本号
	)
	// 定义行为事件 [行为触发时机如: rpc服务初始化前,rpc服务返回结果之后,....]
	srv.Init(
		micro.BeforeStart(func() error {
			logger.Error(fmt.Sprintf("[srv:%s]启动前的日志:", ConfigInfo.String("MicroServiceApi::SRV_NAMESPACE")))
			return nil
		}),
		micro.AfterStart(func() error {
			logger.Error(fmt.Sprintf("[srv:%s]启动后的日志:", ConfigInfo.String("MicroServiceApi::SRV_NAMESPACE")))
			GetEtcdList()
			return nil
		}),
	)

	err := mail.RegisterMailHandler(srv.Server(), handler.Handler()) // 注册服务 [事件,如: 功能方法,sendMail,AddMailConfig,...]
	if nil != err {
		fmt.Println("RegisterMailHandler Fail: " + err.Error())
	}

	// 启动运行rpc服务
	if err := srv.Run(); nil != err {
		panic(fmt.Sprintf("SrvMail Api Run Fail: %v", err))
	}

}

// *************************************************** //
// **********	etcd相关(v2,默认mdns,可忽略)	********** //
// *************************************************** //
// 读取etcd的服务列表
func GetEtcdList() {
	fmt.Println("服务列表:")
	list, e := registry.ListServices()
	for k, v := range list {
		fmt.Println(k, *v)
	}
	fmt.Println(e)
}

// 获取etcd连接信息
func registryOptions(ops *registry.Options) {
	ops.Addrs = []string{"127.0.0.1:2379"}
}
