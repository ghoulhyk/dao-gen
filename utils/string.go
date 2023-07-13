package utils

import (
	"regexp"
	"strings"
)

// Camel2Snake 将驼峰命名法字符串转换为蛇形命名法字符串
func Camel2Snake(str string) string {
	// 匹配所有大写字母和数字
	reg, _ := regexp.Compile("[A-Z0-9]+")
	// 在所有匹配项前添加下划线，并将其全部转换为小写
	snakeStr := strings.ToLower(reg.ReplaceAllStringFunc(str, func(s string) string {
		return "_" + s
	}))
	// 如果字符串以下划线开头，则去掉开头的下划线
	if strings.HasPrefix(snakeStr, "_") {
		snakeStr = snakeStr[1:]
	}
	return snakeStr
}

// Snake2Camel 下划线写法转为驼峰写法
func Snake2Camel(str string) string {
	str = strings.Replace(str, "_", " ", -1)
	str = Capitalize(str)
	return strings.Replace(str, " ", "", -1)
}

// Capitalize 首字母转为大写
func Capitalize(str string) string {
	return strings.ToUpper(string(str[0])) + str[1:]
}

// UnCapitalize 首字母转为小写
func UnCapitalize(str string) string {
	return strings.ToLower(string(str[0])) + str[1:]
}
