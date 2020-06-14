package models

import (
	"fmt"
	. "github.com/duringnone/microbase"
	"github.com/duringnone/microproto/mail"
	"github.com/hyperjiang/php"
	"net/url"
	"time"
)

// 添加邮件模板
func (this *Dao) AddMailConfig(req *mail.AddMailConfigRequest) (int64, string) {
	params := make(map[string]string)
	params["EmailName"] = req.EmailName
	params["EmailTitle"] = req.EmailTitle
	params["EmailContent"] = req.EmailContent
	if errCode, errInfo := IsEmptyMulti(params); 0 != errCode {
		return int64(errCode), errInfo
	}
	// 查重
	selRes, err := this.DB.DBSelectRow(fmt.Sprintf("SELECT * FROM tb_emails WHERE email_name='%s' LIMIT 1", params["EmailName"]))
	if nil != err {
		return Code_DB_SelectErr, "GetEmailConfigInfo SQL Fail: " + err.Error()
	}
	if len(selRes) > 0 {
		return Code_DB_DataExisted, "current EmailConfig was exists"
	}
	// 添加
	addSql := fmt.Sprintf("INSERT INTO tb_emails SET email_name='%s',email_content='%s',email_title='%s'", params["EmailName"], url.QueryEscape(params["EmailContent"]), url.QueryEscape(params["EmailTitle"]))
	if _, err := this.DB.DBUpdate(addSql); nil != err {
		return Code_DB_TransFail, "AddEmailConfig SQL Fail: " + err.Error()
	}
	return 0, Msg_Success
}

// 更新邮件模板
func (this *Dao) UpdateMailConfig(req *mail.UpdateMailConfigRequest) (int64, string) {
	params := make(map[string]string)
	params["EId"] = ToString(req.EId)
	params["EmailContent"] = req.EmailContent
	params["EmailTitle"] = req.EmailTitle

	if errCode, errInfo := IsEmptyMulti(params); 0 != errCode {
		return int64(errCode), errInfo
	}
	if req.EId <= 0 {
		return Code_ParamsFormatErr, "invalid EId"
	}
	// 查重
	selRes, err := this.DB.DBSelectRow(fmt.Sprintf("SELECT * FROM tb_emails WHERE e_id=%d LIMIT 1", req.EId))
	if nil != err {
		return Code_DB_SelectErr, "GetEmailConfigInfo SQL Fail: " + err.Error()
	}
	if len(selRes) == 0 {
		return Code_DB_DataExisted, "current EmailConfig is not exists"
	}
	// 修改
	updateSql := fmt.Sprintf("UPDATE tb_emails SET e_modify_dt='%s',email_content='%s',email_title='%s' WHERE e_id=%d", php.LocalDate("Y-m-d H:i:s", time.Now().Unix()), url.QueryEscape(params["EmailContent"]), url.QueryEscape(params["EmailTitle"]), req.EId)
	if _, err := this.DB.DBUpdate(updateSql); nil != err {
		return Code_DB_TransFail, "UpdateEmailConfig SQL Fail: " + err.Error()
	}
	return 0, Msg_Success
}

// 获取邮件模板列表/详情
func (this *Dao) GetMailConfigList(req *mail.GetMailConfigListRequest) (int64, string, *mail.MailConfigList) {
	listRet := new(mail.MailConfigList)

	// 默认分页 [每页20条]
	page := req.Page
	pageSize := req.PageSize
	if req.Page <= 0 || req.PageSize <= 0 {
		page = 1
		pageSize = 20
	}
	// 查询 [不提供全表扫描]
	fields := "*"
	sql := fmt.Sprintf("SELECT %s  FROM tb_emails  ORDER BY e_id DESC LIMIT %d,%d ", fields, (page-1)*pageSize, pageSize)
	countSql := "SELECT count(*) counts FROM tb_emails "
	counts, err := this.DB.DBSelectRow(countSql)
	if nil != err {
		return Code_DB_SelectErr, "GetMailList Count SQL Fail: " + err.Error(), listRet
	}
	list, err := this.DB.DBSelect(sql)
	if nil != err {
		return Code_DB_SelectErr, "GetMailList List SQL Fail: " + err.Error(), listRet
	}
	list = UrlDecodeListByField(list, "email_title,email_content") // 批量urlDecode
	// 切片=> json => struct
	jsonObj, err := Json.Marshal(list)
	if nil != err {
		return Code_JsonEncodeErr, "JsonEncode Fail: " + err.Error(), listRet
	}
	var resList []*mail.MailConfigSingle
	err = Json.Unmarshal(jsonObj, &resList)
	if nil != err {
		return Code_JsonEncodeErr, "JsonEncode Fail: " + err.Error(), listRet
	}
	listRet.TotalCount = ToInt64(counts["counts"])
	listRet.PageSize = pageSize
	listRet.Page = page
	listRet.List = resList
	return 0, Msg_Success, listRet
}
