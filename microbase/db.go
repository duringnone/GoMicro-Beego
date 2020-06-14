package microbase

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"sync"
)

type DBMysql struct {
	Pool *sql.DB
	Base *BaseController
}

//声明DB连接池锁
var dbmutex sync.RWMutex

//多节点数据库连接池
var dbpool = make(map[string]*sql.DB)

//初始化
func NewDB(u, p, h, P, db string, base *BaseController) (*DBMysql, error) {
	if u == "" || p == "" || h == "" || P == "" || db == "" {
		return nil, errors.New("Please Check Init Param！")
	}
	node := Md5(u + p + h + P + db)
	dbmutex.RLock()
	if _, ok := dbpool[node]; ok {
		dbmutex.RUnlock()
		return &DBMysql{
			Pool: dbpool[node],
			Base: base,
		}, nil
	}
	dbmutex.RUnlock()
	dbresource, err := sql.Open("mysql", u+":"+p+"@tcp("+h+":"+P+")/"+db+"?charset=utf8&allowOldPasswords=1")
	if err != nil {
		return nil, err
	}
	dbresource.SetMaxOpenConns(200)
	dbresource.SetMaxIdleConns(20)
	dbmutex.Lock()
	dbpool[node] = dbresource
	dbmutex.Unlock()
	return &DBMysql{
		Pool: dbresource,
		Base: base,
	}, nil
}

//数据库查询封装列表
func (this *DBMysql) DBSelect(sql string) ([]map[string]string, error) {
	resArr, err := this.ExecQuery(sql, 3)
	if err != nil {
		err = errors.New("ExecQuery:" + sql + ";Error:" + err.Error())
		return resArr, err
	}
	return resArr, nil
}

//数据库查询封装单条
func (this *DBMysql) DBSelectRow(sql string) (map[string]string, error) {
	resArr, err := this.ExecQuery(sql, 3)
	if err != nil {
		err = errors.New("ExecQuery:" + sql + ";Error:" + err.Error())
		return nil, err
	}
	if len(resArr) == 0 {
		return nil, nil
	} else {
		return resArr[0], nil
	}
}

//数据库更新，插入封装
func (this *DBMysql) DBUpdate(sql string) (int, error) {
	var err error
	var res int
	res, err = this.ExecUpdate(sql, 3)
	if err != nil {
		if res != -1062 {
			err = errors.New("ExecQuery:" + sql + ";Error:" + err.Error())
		}
		return res, err
	}
	return res, nil
}

//数据库事务
func (this *DBMysql) DBTrans(sqls []string) (int, error) {
	var res int
	var err error
	res, err = this.ExecTrans(sqls, 5)
	if err != nil {
		if res != -1062 {
			err = errors.New("DBTrans:" + strings.Join(sqls, ",") + ";Error:" + err.Error())
		}
		return res, err
	}

	return res, nil
}

//查询
func (this *DBMysql) ExecQuery(sql string, timeout int) ([]map[string]string, error) {
	var res []map[string]string
	rows, err := this.Pool.Query(sql)
	if err != nil {
		return res, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return res, err
	}
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	for rows.Next() {
		record := make(map[string]string)
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		if err != nil {
			return res, err
		}
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			} else {
				record[columns[i]] = ""
			}
		}
		res = append(res, record)
	}

	return res, nil
}

//更新
func (this *DBMysql) ExecUpdate(sql string, timeout int) (int, error) {
	var res int

	result, err := this.Pool.Exec(sql)
	if err != nil {
		return 0, err
	}
	var affectedRows int64

	if strings.HasPrefix(strings.ToUpper(sql), "INSERT") {

		affectedRows, _ = result.LastInsertId()
		//没有返回最新插入id  则返回AffectedRows
		if affectedRows == int64(0) {
			affectedRows, _ = result.RowsAffected()
		}
	} else {
		affectedRows, _ = result.RowsAffected()
	}

	res = int(affectedRows)

	return res, nil
}

//事务
func (this *DBMysql) ExecTrans(sqls []string, timeout int) (int, error) {
	var res int
	tx, err := this.Pool.Begin()
	if err != nil {
		return 0, err
	}
	var affectedRows, sucnum int64
	for _, sql := range sqls {
		result, err := tx.Exec(sql)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
		if strings.HasPrefix(sql, "insert") {
			affectedRows, _ = result.LastInsertId()
			//没有返回最新插入id  则返回AffectedRows
			if affectedRows == int64(0) {
				affectedRows, _ = result.RowsAffected()
			}
		} else {
			affectedRows, _ = result.RowsAffected()
		}
		sucnum += sucnum
	}
	tx.Commit()
	res = int(sucnum)

	return res, nil
}
