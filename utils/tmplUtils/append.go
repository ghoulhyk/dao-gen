package tmplUtils

import (
	"github.com/samber/lo"
	"regexp"
	"strings"
)

// Append
// 由于import要写在文件开头
// 将basicTmpl从包名后切开，分别拼接appendTmpls的import和代码段
// 各种样式的import加上模板的样式太难匹配了，所以如果有复杂的import匹配不到，使用 {{/* splitPoint */}} 进行分割
func Append(basicTmpl []byte, appendTmpls ...[]byte) []byte {
	basicTmplStr := string(basicTmpl)
	basicSubTmpl1 := ""
	basicSubTmpl2 := ""
	splitIndex := 0
	packageMatch := regexp.MustCompile("package\\s+({{[ \\S]+}}|\\S+)").FindStringIndex(basicTmplStr)
	if len(packageMatch) == 0 {
		panic("没匹配到 package")
	}
	splitIndex = packageMatch[1]
	basicSubTmpl1 = basicTmplStr[:splitIndex]
	basicSubTmpl2 = basicTmplStr[splitIndex:]

	for _, appendTmpl := range appendTmpls {
		appendTmplStr := string(appendTmpl)
		splitIndex = -1
		splitIndex = strings.Index(appendTmplStr, "{{/* splitPoint */}}")
		if splitIndex == -1 {
			regex := regexp.MustCompile("import\\s*(\\([^()]+?\\)|\"[^\"]*\"|[^\"\\s()<>\\[\\]`]+)*")
			importMatch := regex.FindAllStringIndex(appendTmplStr, -1)

			if len(importMatch) > 0 {
				lastMatch, _ := lo.Last(importMatch)
				splitIndex = lastMatch[1]
			}
		}
		if splitIndex >= 0 {
			basicSubTmpl1 += "\n"
			basicSubTmpl2 += "\n"
			basicSubTmpl1 += appendTmplStr[:splitIndex]
			basicSubTmpl2 += appendTmplStr[splitIndex:]
		} else {
			basicSubTmpl2 += "\n"
			basicSubTmpl2 += appendTmplStr
		}
	}

	return []byte(basicSubTmpl1 + "\n" + basicSubTmpl2)
}
