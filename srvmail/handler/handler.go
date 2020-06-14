package handler

import (
	. "github.com/duringnone/microbase"
	"github.com/duringnone/microproto/mail"
	"srvmail/models"
)

/**
 * rpc层共用入口层
 */
// handler私有对象
type handler struct {
	dao *models.Dao
}

// handler层入口函数
func Handler() mail.MailHandler {
	return &handler{}
}

// 初始化加载资源
// @param serialNo 流水号
func (h *handler) InitDao(serialNo string) {
	var err error
	// 初始化时,选择new(BaseController),而非nil,因为goMicro依赖
	if h.dao, err = models.NewDao(serialNo, new(BaseController)); err != nil {
		h.dao.StdLogger(-98, err.Error(), "", serialNo, ConfigInfo.String("MicroServiceApi::SRV_NAME")) // 上报错误
		panic("初始化全局资源Dao失败: " + err.Error())
		return
	}
}
