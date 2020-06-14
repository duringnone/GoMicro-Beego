// @APIVersion 1.0.0
// @Title Micro Gateway
// @Description Gateway Service Of MicroService
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"gateway/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// gateway路由注册
	ns := beego.NewNamespace("/gateway",
		beego.NSNamespace("/mail",
			beego.NSInclude(
				&controllers.MailCtrl{},
			),
		),
	)
	// 添加命名空间
	beego.AddNamespace(ns)
}
