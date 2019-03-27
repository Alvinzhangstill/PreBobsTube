package util

import (
	"strings"
)

// StrToStrid 将字符串小写 用 ‘-’ 连接
// Big Tits --> big-tits
func StrToStrid(str string) string {
	// 字母小写
	lowerStr := strings.ToLower(str)
	// 空格分割
	strSplit := strings.Split(lowerStr, " ")
	// '-' 连接
	result := strings.Join(strSplit, "-")
	return result
}

// 将字符串转为SQL查询的list格式
// 例如  abc####def  -->  ('abc','def')
// select * from xx where a in (abc,def)
func StrToSQLFilterList(str, splitStr string) string {
	strSplitList := strings.Split(str, splitStr)
	resultStr := "("
	for i, v := range strSplitList {
		vIDs := StrToStrid(v) // 将 Big Tits --> big-tits
		if i+1 != len(strSplitList) {
			resultStr += "'" + vIDs + "'," // 'a',
		} else {
			resultStr += "'" + vIDs + "'" // 'b')
		}
	}
	resultStr += ")"
	return resultStr
}
