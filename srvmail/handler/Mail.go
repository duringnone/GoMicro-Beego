package handler

import (
	"context"
	. "github.com/duringnone/microbase"
	"github.com/duringnone/microproto/mail"
	"srvmail/models"
)

// 发送邮件
func (h *handler) SendMail(ctx context.Context, req *mail.SendMailRequest, rsp *mail.SendMailResponse) error {
	if "" == req.SerialNo {
		rsp.Code = int64(models.Code_ParamsFormatErr)
		rsp.Msg = h.dao.GetCurrSrvErrInfo("日志流水号SerialNo不得为空")
		return nil
	}
	h.InitDao(req.SerialNo)                 // 初始化资源(DB/Redis...)
	rsp.Code, rsp.Msg = h.dao.SendMail(req) // 发送邮件
	if 0 != rsp.Code {
		rsp.Msg = h.dao.GetCurrSrvErrInfo(rsp.Msg)
	}
	h.dao.StdLogger(rsp.Code, rsp.Msg, rsp.Data, h.dao.SrvSerialNo, ConfigInfo.String("MicroServiceApi::SRV_NAME")) // 上报
	return nil
}
