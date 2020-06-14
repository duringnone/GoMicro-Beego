package rpc

import (
	"context"
	"errors"
	"fmt"
	. "github.com/duringnone/microbase"
	"github.com/duringnone/microproto/mail"
	"github.com/micro/go-micro/v2"
)

/**
 * 邮件服务rpcApi
 */
func initRpcService() micro.Service {
	service := micro.NewService()
	service.Init()
	return service
}

// 邮件rpc实例
type mailHandler struct{}

// 发送邮件
func (h *mailHandler) SendMail(request *mail.SendMailRequest) (*mail.SendMailResponse, error) {
	s := initRpcService()                                                                    // 初始化服务连接
	client := mail.NewMailService(ConfigInfo.String("MicroServerSrv::SRV_MAIL"), s.Client()) // 邮件系统
	rsp, err := client.SendMail(context.Background(), request)
	if err != nil {
		fmt.Println("call err: ", err, rsp)
		return rsp, errors.New(fmt.Sprintf("errInfo: %s ,\n MailResponse: %v ", err.Error(), rsp))
	}
	if 0 != rsp.Code {
		return rsp, errors.New("MailResponse: " + rsp.Msg)
	}
	return rsp, nil
}
