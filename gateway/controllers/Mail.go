package controllers

import (
	"gateway/models"
	. "github.com/duringnone/microbase"
)

// 网关层-邮件服务
type MailCtrl struct {
	dao *models.Dao
	BaseController
}

// 预处理,加载数据库资源
func (this *MailCtrl) Prepare() {
	var err error
	if this.dao, err = models.NewDao(&this.BaseController); err != nil {
		this.StdOutMsg(-98, err.Error(), "", this.dao.GwSerialNo)
		return
	}
}

// SendMail
// @Title 自定义邮件发送
// @Description 自定义邮件发送
// @Param	mail			query 	string	true		"邮箱"
// @Param	title			query 	string	false		"邮件标题"
// @Param	content			query 	string	false		"邮件内容"
// @Success 200 msg=Success
// @router /send-mail [get]
func (this *MailCtrl) SendMail() {
	code, msg, data := this.dao.SendMail()
	this.StdOutMsg(code, msg, data, this.dao.GwSerialNo)
	return
}

// AddMailConfig
// @Title 添加邮件模板
// @Description 添加邮件模板
// @Param	emailName			query 	string	true		"邮件配置名"
// @Param	emailTitle			query 	string	true		"邮件标题"
// @Param	emailContent		query 	string	true		"邮件内容"
// @Success 200 msg=Success
// @router /add-mail-config [get]
func (this *MailCtrl) AddMailConfig() {
	status, msg, res := this.dao.AddMailConfig()
	this.StdOutMsg(status, msg, res.Data, this.dao.GwSerialNo)
	return
}

// UpdateMailConfig
// @Title 更新邮件模板
// @Description 更新邮件模板
// @Param	eId					query 	string	true		"模板邮件表ID"
// @Param	emailTitle			query 	string	true		"模板邮件标题"
// @Param	emailContent		query 	string	true		"模板邮件内容"
// @Success 200 msg=Success
// @router /update-mail-config [get]
func (this *MailCtrl) UpdateMailConfig() {
	status, msg, res := this.dao.UpdateMailConfig()
	this.StdOutMsg(status, msg, res.Data, this.dao.GwSerialNo)
	return
}

// GetMailConfigList
// @Title 获取件模板列表
// @Description 获取件模板列表
// @Param	page				query 	string	false		"当前页数"
// @Param	pageSize			query 	string	false		"每页显示数"
// @Success 200 msg=Success
// @router /get-mail-config-list [get]
func (this *MailCtrl) GetMailConfigList() {
	status, msg, res := this.dao.GetMailConfigList()
	this.StdOutMsg(status, msg, res.Data, this.dao.GwSerialNo)
	return
}
