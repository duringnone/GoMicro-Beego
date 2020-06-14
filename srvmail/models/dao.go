package models

import (
	. "github.com/duringnone/microbase"
)

type Dao struct {
	DB *DBMysql
	*BaseController
	SrvSerialNo string //	请求日志(每次请求全局唯一)
}

var (
	dbUser, dbPwd, dbHost, dbPort, dbName, dbName2 string
	redisHost, redisPort, redisPwd                 string
	redisProdHost, redisProdPort, redisProdPwd     string
	mongoUser, mongoPwd, mongoDbName               string
	mongoHost                                      []string
)

func init() {
	dbUser = ConfigInfo.String("DB_" + RunMode + "::user")
	dbPwd = ConfigInfo.String("DB_" + RunMode + "::pwd")
	dbHost = ConfigInfo.String("DB_" + RunMode + "::host")
	dbPort = ConfigInfo.String("DB_" + RunMode + "::port")
	dbName = ConfigInfo.String("DB_" + RunMode + "::db")

	redisHost = ConfigInfo.String("Redis_" + RunMode + "::host")
	redisPort = ConfigInfo.String("Redis_" + RunMode + "::port")
	redisPwd = ConfigInfo.String("Redis_" + RunMode + "::pwd")
}

// @param serialNo 流水号
func NewDao(serialNo string, base *BaseController) (*Dao, error) {
	if "" == serialNo {
		panic("NewDao()中,流水号serialNo 不得为空")
	}
	// 动态控制调试模式,切换到调试模式下的数据库
	db, err := NewDB(dbUser, dbPwd, dbHost, dbPort, dbName, base)
	if err != nil {
		return nil, err
	}
	dao := Dao{
		db,
		base,
		serialNo, // 邮件服务流水号
	}
	return &dao, nil
}
