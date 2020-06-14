package microbase

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/json-iterator/go"
	"net"
	"strings"
)

type BaseController struct {
	beego.Controller
	JsonName string
	Logs     []map[string]interface{}
}

var Ip string
var RunMode string
var ConfigInfo config.Configer
var Json = jsoniter.ConfigCompatibleWithStandardLibrary

func init() {
	ConfigInfo = beego.AppConfig
	// 根据ip判断是否为DEV环境,CONST_DEV_IP 在gateway/conf/app.conf, srvmail/conf/app.conf,中配置
	Ip = GetLocalIP()
	if InArray(Ip, ConfigInfo.Strings("CONST_DEV_IP")) {
		RunMode = "DEV"
	} else {
		RunMode = "PROD"
	}
}

// 获取当前 IPV4 地址
func GetLocalIP() string {
	addrs, _ := net.InterfaceAddrs()

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// 标准输出格式 [rpc响应]
// @param outputCode 错误码,0-成功,非0-错误
// @param sMsg 错误信息,Success-成功,否则为失败信息
// @param sData 数据
// @param vSerialNo 日志流水号(每次请求唯一)
// @param toJson 数据
// @return 输出json格式,如: {code:0,msg:success,data:[]}
func (this *BaseController) StdOutMsg(outputCode interface{}, sMsg interface{}, sData interface{}, vSerialNo string, toJson ...bool) {
	var logStr string
	logStr = "[vSerialNo]:" + vSerialNo + "<br/><br/>"
	logStr += "[ServiceName]: GatewayService <br/><br/>" // 网关层服务
	logStr += "[RequestURL]:" + this.Ctx.Request.RequestURI + "<br/><br/>"
	if len(this.Logs) > 0 {
		for _, v := range this.Logs {
			for k, val := range v {
				s, _ := Json.Marshal(val)
				logStr += k + strings.Replace(string(s), "\x00", "", -1) + "<br/><br/>"
			}
		}
		// 如需错误上报,可自行实现,不上报不影响业务功能
	}
	rettype := Escape(this.GetString("r1"))
	callback := Escape(this.GetString("callback"))
	download := Escape(this.GetString("download")) //添加支持文件下载功能
	if rettype != "" && len(toJson) == 0 {         // 是否支持跨域(将返回值赋值给js变量r1,Get请求跨域加载r1)
		this.JsonName = "var " + rettype + " = "
	}
	if download != "" { // 支持get参数download,get直接下载文件
		file, _ := sMsg.(string)
		this.Ctx.Output.Download(file)
		return
	}
	res := make(map[string]interface{})
	res["errCode"] = outputCode
	res["errMsg"] = sMsg
	res["data"] = sData
	res["logSerialNo"] = vSerialNo
	this.Ctx.Output.Header("serial", vSerialNo)
	if callback == "" {
		this.Data["json"] = &res
		this.ServeJSON()
	} else {
		this.Data["jsonp"] = &res
		this.ServeJSONP()
	}
	return
}

// 微服务上报日志方法
// @param outputCode 错误码,0-成功,非0-错误
// @param sMsg 错误信息,Success-成功,否则为失败信息
// @param sData 数据
// @param srvName 服务名,如:srvMail-邮件服务
// @param vSerialNo 微服务日志流水号(每次请求唯一)
func (this *BaseController) StdLogger(outputCode interface{}, sMsg interface{}, sData interface{}, vSerialNo, srvName string) {
	var logStr string
	logStr = "[vSerialNo]:" + vSerialNo + "<br/><br/>"  // 流水号
	logStr += "[ServiceName]:" + srvName + "<br/><br/>" // 服务名
	if len(this.Logs) > 0 {
		for _, v := range this.Logs {
			for k, val := range v {
				s, _ := Json.Marshal(val)
				logStr += k + strings.Replace(string(s), "\x00", "", -1) + "<br/><br/>"
			}
		}
		// 如需错误上报,可自行实现,不上报不影响业务功能
	}
	return
}
