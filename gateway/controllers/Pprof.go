package controllers

import (
	"gateway/models"
	. "github.com/duringnone/microbase"
	"net/http/pprof"
)

// pprof性能分析工具
type PprofCtrl struct {
	dao *models.Dao
	BaseController
}

// 预处理,加载数据库资源
func (this *PprofCtrl) Prepare() {
	var err error
	if this.dao, err = models.NewDao(&this.BaseController); err != nil {
		this.StdOutMsg(-98, err.Error(), "", this.dao.GwSerialNo)
		return
	}
}

func (c *PprofCtrl) Get() {
	switch c.Ctx.Input.Param(":app") {
	default:
		pprof.Index(c.Ctx.ResponseWriter, c.Ctx.Request)
	case "":
		pprof.Index(c.Ctx.ResponseWriter, c.Ctx.Request)
	case "cmdline":
		pprof.Cmdline(c.Ctx.ResponseWriter, c.Ctx.Request)
	case "profile":
		pprof.Profile(c.Ctx.ResponseWriter, c.Ctx.Request)
	case "symbol":
		pprof.Symbol(c.Ctx.ResponseWriter, c.Ctx.Request)
	}
	c.Ctx.ResponseWriter.WriteHeader(200)
}
