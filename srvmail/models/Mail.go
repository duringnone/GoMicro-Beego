package models

import (
	. "github.com/duringnone/microbase"
	"github.com/duringnone/microproto/mail"
	"strings"
)

/**
 * 邮件系统
 */

// 单条发送 [对外]
func (this *Dao) SendMail(req *mail.SendMailRequest) (int64, string) {
	params := make(map[string]string)
	params["mail"] = strings.Trim(req.Mail, ",")
	if errCode, errInfo := IsEmptyMulti(params); 0 != errCode {
		return int64(errCode), errInfo
	}
	params["title"] = req.Title
	params["content"] = req.Content
	if errCode, errInfo := IsEmptyMulti(params); 0 != errCode {
		return int64(errCode), errInfo
	}
	// 验证邮箱格式
	if false == CheckEmail(params["mail"]) {
		return Code_ParamsFormatErr, "邮箱格式错误"
	}
	// 发送邮件
	if err := this.SendDefineEmail(params["mail"], params["title"], params["content"]); nil != err {
		return Code_EmailSendFail, err.Error()
	}
	return 0, Msg_Success
}
