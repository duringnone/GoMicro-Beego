package handler

import (
	"context"
	. "github.com/duringnone/microbase"
	"github.com/duringnone/microproto/mail"
	"srvmail/models"
)

// 添加邮件模板
func (h *handler) AddMailConfig(ctx context.Context, req *mail.AddMailConfigRequest, rsp *mail.AddMailConfigResponse) error {
	if "" == req.SerialNo {
		rsp.Code = int64(models.Code_ParamsFormatErr)
		rsp.Msg = h.dao.GetCurrSrvErrInfo("日志流水号SerialNo不得为空")
		return nil
	}
	h.InitDao(req.SerialNo)
	rsp.Code, rsp.Msg = h.dao.AddMailConfig(req) // 添加邮件模板
	if 0 != rsp.Code {
		rsp.Msg = h.dao.GetCurrSrvErrInfo(rsp.Msg)
	}
	h.dao.StdLogger(rsp.Code, rsp.Msg, rsp.Data, h.dao.SrvSerialNo, ConfigInfo.String("MicroServiceApi::SRV_NAME")) // 上报
	return nil
}

// 修改邮箱模板内容
func (h *handler) UpdateMailConfig(ctx context.Context, req *mail.UpdateMailConfigRequest, rsp *mail.UpdateMailConfigResponse) error {
	if "" == req.SerialNo {
		rsp.Code = int64(models.Code_ParamsFormatErr)
		rsp.Msg = h.dao.GetCurrSrvErrInfo("日志流水号SerialNo不得为空")
		return nil
	}
	h.InitDao(req.SerialNo)
	rsp.Code, rsp.Msg = h.dao.UpdateMailConfig(req) // 修改邮箱模板内容
	if 0 != rsp.Code {
		rsp.Msg = h.dao.GetCurrSrvErrInfo(rsp.Msg)
	}
	h.dao.StdLogger(rsp.Code, rsp.Msg, rsp.Data, h.dao.SrvSerialNo, ConfigInfo.String("MicroServiceApi::SRV_NAME")) // 上报
	return nil
}

// 获取邮箱模板列表/详情
func (h *handler) GetMailConfigList(ctx context.Context, req *mail.GetMailConfigListRequest, rsp *mail.GetMailConfigListResponse) error {
	if "" == req.SerialNo {
		rsp.Code = int64(models.Code_ParamsFormatErr)
		rsp.Msg = h.dao.GetCurrSrvErrInfo("日志流水号SerialNo不得为空")
		return nil
	}
	h.InitDao(req.SerialNo)
	rsp.Code, rsp.Msg, rsp.Data = h.dao.GetMailConfigList(req) // 修改邮箱模板发布状态
	if 0 != rsp.Code {
		rsp.Msg = h.dao.GetCurrSrvErrInfo(rsp.Msg)
	}
	h.dao.StdLogger(rsp.Code, rsp.Msg, rsp.Data, h.dao.SrvSerialNo, ConfigInfo.String("MicroServiceApi::SRV_NAME")) // 上报
	return nil
}
