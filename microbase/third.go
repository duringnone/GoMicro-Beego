package microbase

import (
	"errors"
	gomail "gopkg.in/gomail.v2"
	"strings"
)

/**
 * 第三方工具包(邮箱发送)
 */
type MailConf struct {
	Username string // 发件人邮箱
	Pwd      string // 发件人 授权码,而非密码
	Smtp     string // 协议
	Port     int    // (邮件服务器)端口
}

/**
 * 注意：
 *		1） gmail邮件附件容量有上限(多个附件容量总和)，不得超过25MB,原则上建议不要超过20MB，官方链接： https://support.google.com/mail/answer/6584?p=MaxSizeError
 */
//  @param userMail	 接收者邮箱"demon@qq.com"
//	@param mailTitle 邮件标题
//	@param mailContent 邮件HTML文本内容(含html标签)
//	@param attachFilePath 邮件附件文件路径 (务必保证文件存在)
//  @return int,error 错误码,错误信息
func (this *BaseController) SendMailNotify(emailConf *MailConf, userMail, mailTitle, mailContent, attachFilePath string) (int, error) {
	m := gomail.NewMessage()
	m.SetHeader("From", emailConf.Username)
	m.SetHeader("To", userMail)
	m.SetHeader("Subject", mailTitle)
	m.SetBody("text/html", mailContent)
	if fileArr := strings.Split(attachFilePath, ","); "" != attachFilePath && len(fileArr) > 0 {
		for _, file := range fileArr {
			if "" == file {
				continue
			}
			m.Attach(file) // 添加附件
		}
	}

	d := gomail.NewDialer(emailConf.Smtp, emailConf.Port, emailConf.Username, emailConf.Pwd) // 授权码,而非密码

	if err := d.DialAndSend(m); err != nil {
		return -80001, errors.New("send email fail: " + err.Error())
	}
	return 0, nil
}

/**
 * 注意: 此方法功能同上,区别在于此方法在脚本中使用,上述方法在beego中使用
 * 注意：
 *		1） gmail邮件附件容量有上限(多个附件容量总和)，不得超过25MB,原则上建议不要超过20MB，官方链接： https://support.google.com/mail/answer/6584?p=MaxSizeError
 */
func SendMailPush(emailConf *MailConf, userMail, mailTitle, mailContent, attachFilePath string) (int, error) {
	m := gomail.NewMessage()
	m.SetHeader("From", emailConf.Username)
	m.SetHeader("To", userMail)
	m.SetHeader("Subject", mailTitle)
	m.SetBody("text/html", mailContent)
	// 支持添加多个附件（附件路径中不得含有逗号）
	if fileArr := strings.Split(attachFilePath, ","); "" != attachFilePath && len(fileArr) > 0 {
		for _, file := range fileArr {
			if "" == file {
				continue
			}
			m.Attach(file) // 添加附件
		}
	}

	d := gomail.NewDialer(emailConf.Smtp, emailConf.Port, emailConf.Username, emailConf.Pwd) // 授权码,而非密码

	if err := d.DialAndSend(m); err != nil {
		//发送失败错误处理
		return -80002, errors.New("send email fail: " + err.Error())
	}
	return 0, nil
}
