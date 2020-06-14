package rpc

import (
	"context"
	"errors"
	"fmt"
	. "github.com/duringnone/microbase"
	"github.com/duringnone/microproto/mail"
	"github.com/micro/go-micro/v2/metadata"
)

/**
 * RPC调用
 */

// 添加邮件模板
func (h *mailHandler) AddMailConfig(request *mail.AddMailConfigRequest) (*mail.AddMailConfigResponse, error) {
	// mail包调用
	s := initRpcService()                                                                    // 初始化服务连接
	client := mail.NewMailService(ConfigInfo.String("MicroServerSrv::SRV_MAIL"), s.Client()) // 邮件系统
	rsp, err := client.AddMailConfig(context.Background(), request)
	if err != nil {
		return rsp, errors.New(fmt.Sprintf("errInfo: %s ,\n MailResponse: %v ", err.Error(), rsp))
	}
	if 0 != rsp.Code {
		return rsp, errors.New("MailResponse: " + rsp.Msg)
	}
	return rsp, nil
}

// 更新邮件模板内容
func (h *mailHandler) UpdateMailConfig(request *mail.UpdateMailConfigRequest) (*mail.UpdateMailConfigResponse, error) {
	s := initRpcService()                                                                    // 初始化服务连接
	client := mail.NewMailService(ConfigInfo.String("MicroServerSrv::SRV_MAIL"), s.Client()) // 邮件系统
	rsp, err := client.UpdateMailConfig(context.Background(), request)
	if err != nil {
		return rsp, errors.New(fmt.Sprintf("errInfo: %s ,\n MailResponse: %v ", err.Error(), rsp))
	}
	if 0 != rsp.Code {
		return rsp, errors.New("MailResponse: " + rsp.Msg)
	}
	return rsp, nil
}

// 获取邮件模板列表
func (h *mailHandler) GetMailConfigList(request *mail.GetMailConfigListRequest) (*mail.GetMailConfigListResponse, error) {
	s := initRpcService()                                                                    // 初始化服务连接
	client := mail.NewMailService(ConfigInfo.String("MicroServerSrv::SRV_MAIL"), s.Client()) // 邮件系统
	rsp, err := client.GetMailConfigList(context.Background(), request)
	if err != nil {
		return rsp, errors.New(fmt.Sprintf("errInfo: %s ,\n MailResponse: %v ", err.Error(), rsp))
	}
	if 0 != rsp.Code {
		return rsp, errors.New("MailResponse: " + rsp.Msg)
	}
	return rsp, nil
}

// ******************************************************************** //
// *********************	拓展: 可忽略 	*************************** //
// ******************************************************************** //
// *****注:
//      AddMailConfig2() 和 AddMailConfig() 区别: 后者是前者的封装,后者提供selector和更多样功能的实现,建议用后者AddMailConfig()
// 添加邮件模板 [proto中底层实现AddMailConfig的过程,直接rpc调用]
func (h *mailHandler) AddMailConfig2(request *mail.AddMailConfigRequest) (*mail.AddMailConfigResponse, error) {
	// 初始化服务连接
	s := initRpcService()
	c := s.Client()
	//初始化请求(服务名,方法,请求参数,响应参数)
	req := c.NewRequest(ConfigInfo.String("MicroServerSrv::SRV_MAIL"), "Mail.AddMailConfig", request)
	rsp := &mail.AddMailConfigResponse{}

	// 创建上下文和元数据  [用于数据传输,web->rpc,个人理解: 类似Context上下文]
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "test.com.user",
		"X-From-Id": "AddMailConfig",
	})

	// 发送rpc请求
	if err := c.Call(ctx, req, rsp); err != nil {
		fmt.Println("call err: ", err, rsp)
		return rsp, errors.New(fmt.Sprintf("errInfo: %s ,\n MailResponse: %v ", err.Error(), rsp))
	}
	return rsp, nil
}
