package models

import (
	. "github.com/duringnone/microbase"
)

//发送自定义邮件模板
// @param string email 邮箱
// @param string mailTitle 邮箱标题
// @param string mailContent  邮箱内容
// @return error err 错误对象
func (this *Dao) SendDefineEmail(email, mailTitle, mailContent string) error {
	mailPort, _ := ConfigInfo.Int("Email::port")
	emailConf := &MailConf{
		ConfigInfo.String("Email::username"),
		ConfigInfo.String("Email::password"),
		ConfigInfo.String("Email::smtp"),
		mailPort,
	} // 初始化邮箱配置
	attachFilePath := "" //附件路径
	if _, err := this.SendMailNotify(emailConf, email, mailTitle, mailContent, attachFilePath); nil != err {
		return err
	}
	return nil
}

// 获取邮件服务api的错误信息 [所有错误信息加上前缀: Msg_errInfo_prefix]
// @param 错误信息
// @return 拼接上错误信息前缀的错误信息
func (this *Dao) GetCurrSrvErrInfo(str string) string {
	return Msg_errInfo_prefix + str
}
