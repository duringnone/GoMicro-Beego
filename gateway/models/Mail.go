package models

import (
	"gateway/rpc"

	"github.com/duringnone/microproto/mail"
	"strings"
)

/**
 * 邮件系统
 */

// rpc版本
// 发送邮件
func (this *Dao) SendMail() (int, string, *mail.SendMailResponse) {
	rsp := new(mail.SendMailResponse)
	params := make(map[string]string)
	params["mail"] = strings.Trim(this.GetString("mail"), ",")
	params["title"] = this.GetString("title")
	params["content"] = this.GetString("content")

	// rpc调用
	req := &mail.SendMailRequest{
		SerialNo: this.GwSerialNo,
		Title:    params["title"],
		Content:  params["content"],
		Mail:     params["mail"],
	}
	rsp, err := rpc.GetRpcHandler().SendMail(req)
	if nil != err {
		return Code_RPC_Response_Error, err.Error(), rsp
	}
	return 0, Msg_Success, rsp
}
