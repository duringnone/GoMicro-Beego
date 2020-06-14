package models

/******		错误码	******/
// 业务错误码
const Code_EmailSendFail = -10007   // 邮件发送失败
const Code_ParamsFormatErr = -10008 // 参数格式错误
const Code_JsonDecodeErr = -10009   // json_decode出错
const Code_JsonEncodeErr = -10010   // json_encode出错

// DB数据库错误码
const Code_DB_DataExisted = -20001 //当前查重数据已存在
const Code_DB_SelectErr = -20002   //数据库查询异常
const Code_DB_InsertErr = -20003   //数据库插入异常
const Code_DB_UpdateErr = -20004   // 数据库修改记录异常
const Code_DB_TransFail = -20005   // 事务失败

/******		错误信息	******/
const (
	Msg_Success = "Success" // 成功返回信息

	Msg_errInfo_prefix = "MailApi Error: " // MailApi错误前缀

)
