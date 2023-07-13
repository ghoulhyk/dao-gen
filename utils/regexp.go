package utils

import "regexp"

// GetColAttrRegexp
// tag中数据库字段信息的正则
// 例: `json:"id" colAttr:"'id' pk autoincr"`   中的   colAttr:"'id' pk autoincr"
func GetColAttrRegexp() *regexp.Regexp {
	return regexp.MustCompile(" *colAttr:\"([\\s\\S]+)\" *")
}
