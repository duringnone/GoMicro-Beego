package rpc

/**
 * rpc层公共组件
 * 定义复用结构体
 */

// rpc层返回结构
type RpcResponse struct {
	Code int64    `json:错误码`
	Msg  string   `json:错误信息`
	Data []string `json:返回数据`
}

// rpc对象 [外部调用结构: GetRpcHander()-->RpcHandler-->mailHandler极其所属方法]
// 注: [可继续在此结构体中添加属性/方法,如:获取当前服务名GetCurrServiceName()]
type RpcHandler struct {
	*mailHandler // 邮件rpc对象
}

// 返回rpc层对象
func GetRpcHandler() RpcHandler {
	return RpcHandler{}
}
