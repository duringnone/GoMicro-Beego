package models

import (
	"gateway/rpc"
	. "github.com/duringnone/microbase"
	"github.com/duringnone/microproto/mail"
)

/**
 * 系统模板邮件配置
 */

// 添加邮件模板
func (this *Dao) AddMailConfig() (int, string, *mail.AddMailConfigResponse) {
	params := make(map[string]string)
	rsp := new(mail.AddMailConfigResponse)
	var err error
	params["emailName"] = this.GetString("emailName")
	params["emailTitle"] = this.GetString("emailTitle")
	params["emailContent"] = this.GetString("emailContent")

	// rpc调用
	req := &mail.AddMailConfigRequest{
		SerialNo:     this.GwSerialNo,
		EmailName:    params["emailName"],
		EmailTitle:   params["emailTitle"],
		EmailContent: params["emailContent"],
	}
	rsp, err = rpc.GetRpcHandler().AddMailConfig(req)
	if nil != err {
		return Code_RPC_Response_Error, err.Error(), rsp
	}
	return 0, Msg_Success, rsp
}

// 获取模板列表
func (this *Dao) GetMailConfigList() (int, string, *mail.GetMailConfigListResponse) {
	params := make(map[string]string)
	rsp := new(mail.GetMailConfigListResponse)
	params["page"] = this.GetString("page")
	params["pageSize"] = this.GetString("pageSize")

	// rpc调用
	req := &mail.GetMailConfigListRequest{
		SerialNo: this.GwSerialNo,
		Page:     ToInt64(params["page"]),
		PageSize: ToInt64(params["pageSize"]),
	}
	rsp, err := rpc.GetRpcHandler().GetMailConfigList(req)
	if nil != err {
		return Code_RPC_Response_Error, err.Error(), rsp
	}
	return 0, Msg_Success, rsp
}

// 更新邮件模板
func (this *Dao) UpdateMailConfig() (int, string, *mail.UpdateMailConfigResponse) {
	params := make(map[string]string)
	rsp := new(mail.UpdateMailConfigResponse)
	var err error
	params["emailTitle"] = this.GetString("emailTitle")
	params["emailContent"] = this.GetString("emailContent")
	params["eId"] = this.GetString("eId")
	if errCode, errInfo := IsEmptyMulti(params); 0 != errCode {
		return errCode, errInfo, rsp
	}

	// rpc调用
	req := &mail.UpdateMailConfigRequest{
		SerialNo:     this.GwSerialNo,
		EId:          ToInt64(params["eId"]),
		EmailTitle:   params["emailTitle"],
		EmailContent: params["emailContent"],
	}
	rsp, err = rpc.GetRpcHandler().UpdateMailConfig(req)
	if nil != err {
		return Code_RPC_Response_Error, err.Error(), rsp
	}
	return 0, Msg_Success, rsp
}
