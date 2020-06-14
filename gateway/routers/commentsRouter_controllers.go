package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["gateway/controllers:MailCtrl"] = append(beego.GlobalControllerRouter["gateway/controllers:MailCtrl"],
		beego.ControllerComments{
			Method:           "AddMailConfig",
			Router:           `/add-mail-config`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["gateway/controllers:MailCtrl"] = append(beego.GlobalControllerRouter["gateway/controllers:MailCtrl"],
		beego.ControllerComments{
			Method:           "GetMailConfigList",
			Router:           `/get-mail-config-list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["gateway/controllers:MailCtrl"] = append(beego.GlobalControllerRouter["gateway/controllers:MailCtrl"],
		beego.ControllerComments{
			Method:           "SendMail",
			Router:           `/send-mail`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["gateway/controllers:MailCtrl"] = append(beego.GlobalControllerRouter["gateway/controllers:MailCtrl"],
		beego.ControllerComments{
			Method:           "UpdateMailConfig",
			Router:           `/update-mail-config`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
