package models

import (
	"fmt"
	. "github.com/duringnone/microbase"
	"github.com/hyperjiang/php"
	"time"
)

type Dao struct {
	*BaseController
	GwSerialNo string //	请求日志(每次请求全局唯一)
}

func init() {
	fmt.Println("#### 当前环境: ", RunMode)
}

func NewDao(base *BaseController) (*Dao, error) {
	dao := Dao{
		base,
		"LogSerialNo_" + php.Date("YmdHis", time.Now().Unix()) + "_" + fmt.Sprintf("%v", time.Now().UnixNano()), // 日志
	}
	return &dao, nil
}
