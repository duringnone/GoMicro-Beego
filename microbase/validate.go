package microbase

import (
	"regexp"
	"strings"
)

/**
 * 验证器
 */

// 邮箱格式验证
// @param string 邮箱
// @return string true-正确邮箱格式,false-错误邮箱格式
func CheckEmail(email string) (b bool) {
	email = strings.ToLower(email) //将字符中所有大写字符全都转为小写
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	if !reg.MatchString(email) {
		return false
	}
	return true
}

// 验证单个string入参是否非空
// @param string 验证参数
// @return string true-非空,false-空
func IsEmpty(s string) (b bool) {
	if "" == s {
		return false
	}
	return true
}

// 验证string类型入参是否为空(支持多个string参数)
// @param map 验证map是否为空
// @return string 某个参数为空的错误信息: *** is Required
func IsEmptyMulti(m map[string]string) (int, string) {
	for k, v := range m {
		if "" == v {
			return -10001, k + " is Required"
		}
	}
	return 0, ""
}
