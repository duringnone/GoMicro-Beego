package microbase

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/hyperjiang/php"
	"net/url"
	"strconv"
	"strings"
	"text/template"
	"time"
)

// ***************************************************** //
// *****************	字符串处理 ******************* //
// ***************************************************** //
// Md5加密
// @param string 加密前的字符串
// @return string 加密后的字符串
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// 生成唯一流水号
// @return string 唯一流水号
// @param string 流水前缀名
// @return 唯一流水号
func CreateUniqueFlowNum(prefix ...string) string {
	pre := "System-Flow"
	if len(prefix) > 0 {
		pre = prefix[0]
	}
	return pre + "-" + php.Date("YmdHis", time.Now().Unix()) + "-" + fmt.Sprintf("%v", time.Now().UnixNano())
}

// 过滤特殊字符
// @param key 字符串
// @param sType 类型 [可选]
// @return string 过滤后字符串
func Escape(key string, sType ...string) string {
	if key == "" {
		return ""
	}
	specialStrList := []string{";", "(", ")", "\r", "\n", "*"} // 需要过滤的特殊字符
	if len(sType) > 0 {
		return ToString(ToInt(key))
	} else {
		for _, specStr := range specialStrList {
			key = strings.Replace(key, specStr, "", -1)
		}
		return template.HTMLEscapeString(template.JSEscapeString(key))
	}
}

// ***************************************************** //
// *****************	数组/map处理 ******************* //
// ***************************************************** //
// 合并数组
// @params map数组切片 args 多个map数组，逗号间隔，格式：[]map[string]string,[]map[string]string,...
// @return []map[string]string res
func ArrayMerge(args ...[]map[string]string) []map[string]string {
	res := []map[string]string{}
	if len(args) <= 0 {
		return res
	}
	res = args[0]
	for i, c := 1, len(args); i < c; i++ {
		for _, val := range args[i] {
			res = append(res, val)
		}
	}
	return res
}

// 数组/map去重 [支持5种格式格式:[]string,map,[]int] ; 若为map邮箱重复,只返回第一组匹配的值
// @param m 被去重的参数; 三种格式支持: []int,[]string,map[string]string,map[int]string,map[string]int
// @return 去重后的参数
func ArrayUnique(m interface{}) interface{} {
	var ret interface{}
	switch m.(type) {
	// map
	case map[string]string: // map[string]string
		res := make(map[string]string)
		tmp := make(map[string]string)
		for k, v := range m.(map[string]string) {
			if _, ok := tmp[v]; ok {
				continue
			}
			tmp[v] = k
		}
		for k, v := range tmp {
			res[v] = k
		}
		ret = res
	case map[string]int64: // map[string]int64
		res := make(map[string]int64)
		tmp := make(map[int64]string)
		for k, v := range m.(map[string]int64) {
			if _, ok := tmp[v]; ok {
				continue
			}
			tmp[v] = k
		}
		for k, v := range tmp {
			res[v] = k
		}
		ret = res
	case map[string]int: // map[string]int
		res := make(map[string]int)
		tmp := make(map[int]string)
		for k, v := range m.(map[string]int) {
			if _, ok := tmp[v]; ok {
				continue
			}
			tmp[v] = k
		}
		for k, v := range tmp {
			res[v] = k
		}
		ret = res
	case map[int]string: // map[int]string
		res := make(map[int]string)
		tmp := make(map[string]int)
		for k, v := range m.(map[int]string) {
			if _, ok := tmp[v]; ok {
				continue
			}
			tmp[v] = k
		}
		for k, v := range tmp {
			res[v] = k
		}
		ret = res
	// 切片
	case []string: // []string切片
		res := []string{}
		tmp := make(map[string]int)
		for k, v := range m.([]string) {
			if _, ok := tmp[v]; ok {
				continue
			}
			tmp[v] = k
		}
		for k, _ := range tmp {
			res = append(res, k)
		}
		ret = res
	case []int: // []int切片
		res := []int{}
		tmp := make(map[int]int)
		for k, v := range m.([]int) {
			if _, ok := tmp[v]; ok {
				continue
			}
			tmp[v] = k
		}
		for k, _ := range tmp {
			res = append(res, k)
		}
		ret = res
	case []int64: // []int64切片
		res := []int64{}
		tmp := make(map[int64]int)
		for k, v := range m.([]int64) {
			if _, ok := tmp[v]; ok {
				continue
			}
			tmp[v] = k
		}
		for k, _ := range tmp {
			res = append(res, k)
		}
		ret = res
	}
	return ret
}

// 过滤活动列表中占用传输带宽的字段值
// @param []map 数据列表
// @param string field 需要被过滤的字段名(支持多个字段,"name" 或 "name,age,sex,...")
// @return []map 过滤后的结果列表
func FilterListByField(m []map[string]string, field string) []map[string]string {
	if len(m) == 0 {
		return m
	}
	fieldArr := strings.Split(field, ",")
	for _, single := range m {
		// 过滤多个字段
		for _, field := range fieldArr {
			if "" == field {
				continue
			}
			delete(single, field) // 去除字段
		}
	}
	return m
}

// 对列表中的字段进行urlDecode
// @param []map 数据列表
// @param string field 需要被urlDecode的字段名(支持多个字段,"name" 或 "name,age,sex,...")
// @return []map urlDecode后的结果列表
func UrlDecodeListByField(m []map[string]string, field string) []map[string]string {
	if len(m) == 0 {
		return m
	}
	fieldArr := strings.Split(field, ",")
	for _, single := range m {
		// 多个字段urlDecode
		for _, field := range fieldArr {
			// 排除无效字段
			if _, ok := single[field]; !ok || "" == field {
				continue
			}
			single[field], _ = url.QueryUnescape(single[field]) // 单字段urlDecode
		}
	}
	return m
}

// 值是否在数组/map中
// @param val 查询的值
// @param arr 被查询的数组/map
// @return bool true-值在数组中,false-值不再数组中
func InArray(val string, arr interface{}) bool {
	isIn := false
	switch arr.(type) {
	case []string: //
		// 去重
		tmp := arr.([]string)
		temp := ArrayUnique(tmp)
		tmp = temp.([]string)
		// 判断是否包含
		for _, v := range tmp {
			if ToString(val) == v {
				isIn = true
				continue
			}
		}
	case []int: //
		// 去重
		tmp := arr.([]int)
		temp := ArrayUnique(tmp)
		tmp = temp.([]int)
		// 判断是否包含
		for _, v := range tmp {
			if ToInt(val) == v {
				isIn = true
				continue
			}
		}
	case []int64: //
		// 去重
		tmp := arr.([]int64)
		temp := ArrayUnique(tmp)
		tmp = temp.([]int64)
		// 判断是否包含
		for _, v := range tmp {
			if ToInt64(val) == v {
				isIn = true
				continue
			}
		}
	case map[string]string: //
		// 去重
		tmp := arr.(map[string]string)
		temp := ArrayUnique(tmp)
		tmp = temp.(map[string]string)
		// 判断是否包含
		for _, v := range tmp {
			if val == v {
				isIn = true
				continue
			}
		}
	case map[string]int64: //
		// 去重
		tmp := arr.(map[string]int64)
		temp := ArrayUnique(tmp)
		tmp = temp.(map[string]int64)
		// 判断是否包含
		for _, v := range tmp {
			if ToInt64(val) == v {
				isIn = true
				continue
			}
		}
	case map[string]int: //
		// 去重
		tmp := arr.(map[string]int)
		temp := ArrayUnique(tmp)
		tmp = temp.(map[string]int)
		// 判断是否包含
		for _, v := range tmp {
			if ToInt(val) == v {
				isIn = true
				continue
			}
		}
	case map[int]string: //
		// 去重
		tmp := arr.(map[int]string)
		temp := ArrayUnique(tmp)
		tmp = temp.(map[int]string)
		// 判断是否包含
		for _, v := range tmp {
			if val == v {
				isIn = true
				continue
			}
		}
	default:
	}
	return isIn
}

// ***************************************************** //
// *****************	数据类型转换	******************* //
// ***************************************************** //
//将参数转换成int
// @param interface{} 任意类型 [string/int系列/float系列]
// @return int int类型
func ToInt(key interface{}) int {
	var res int
	switch key.(type) {
	case string:
		res, _ = strconv.Atoi(key.(string))
	case int:
		res = key.(int)
	case int32:
		res = int(key.(int32))
	case int64:
		res = int(key.(int64))
	case uint32:
		res = int(key.(uint32))
	case uint64:
		res = int(key.(uint64))
	case float32:
		res = int(key.(float32))
	case float64:
		res = int(key.(float64))
	}
	return res
}

//将参数转换成int64
// @param interface{} 任意类型 [string/int系列/float系列]
// @return int64 64位int整型
func ToInt64(key interface{}) int64 {
	var res int64
	switch key.(type) {
	case string:
		res, _ = strconv.ParseInt(key.(string), 10, 64) // 参数2:n-n进制,如:10-10进制(被操作数/结果??),参数3:结果类型,64-int64,32-int32,8-int8...
	case int:
		res = key.(int64)
	case int32:
		res = int64(key.(int32))
	case int64:
		res = key.(int64)
	case uint32:
		res = int64(key.(uint32))
	case uint64:
		res = int64(key.(uint64))
	case float32:
		res = int64(key.(float32))
	case float64:
		res = int64(key.(float64))
	}
	return res
}

//将参数转换成string
// @param interface{} 任意类型 [string/int系列/float系列]
// @return string 字符串
func ToString(key interface{}) string {
	var res string
	switch key.(type) {
	case string:
		res = key.(string)
	case int:
		res = strconv.Itoa(key.(int))
	case int32:
		res = strconv.FormatInt(int64(key.(int32)), 10)
	case int64:
		res = strconv.FormatInt(key.(int64), 10)
	case uint32:
		res = strconv.FormatUint(uint64(key.(uint32)), 10)
	case uint64:
		res = strconv.FormatUint(key.(uint64), 10)
	case float32:
		res = strconv.FormatFloat(float64(key.(float32)), 'G', 20, 64)
	case float64:
		res = strconv.FormatFloat(key.(float64), 'G', 20, 64)
	}
	return res
}

//将参数转换成float32
// @param interface{} 任意类型 [string/int系列/float系列]
// @return float32 32位浮点数
func ToFloat32(key interface{}) float32 {
	var res float32
	switch key.(type) {
	case string:
		f, _ := strconv.ParseFloat(key.(string), 32)
		res = float32(f)
	case int:
		res = float32(key.(int))
	case int32:
		res = float32(key.(int32))
	case int64:
		res = float32(key.(int64))
	case uint32:
		res = float32(key.(uint32))
	case uint64:
		res = float32(key.(uint64))
	case float32:
		res = float32(key.(float32))
	case float64:
		res = float32(key.(float64))
	}
	return res
}
