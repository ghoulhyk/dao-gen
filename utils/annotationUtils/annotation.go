package annotationUtils

import (
	"github.com/dave/dst"
	"github.com/samber/lo"
	"strings"
)

func Exist(annotations dst.Decorations, findAnno string) bool {
	_, exist := existAnno(annotations, findAnno)
	return exist
}

func FindStr(annotations dst.Decorations, findAnno string) (val string, exist bool) {
	return existAnno(annotations, findAnno)
}

func existAnno(annotations dst.Decorations, findAnno string) (val string, exist bool) {
	if !strings.HasPrefix(findAnno, "@") {
		findAnno = "@" + findAnno
	}
	for _, s := range annotations.All() {
		s = strings.TrimSpace(s)
		s = strings.TrimLeft(s, "/")
		s = strings.TrimLeft(s, "*")   // 针对 /* xxx */
		s = strings.TrimRight(s, "*/") // 针对 /* xxx */
		s = strings.TrimSpace(s)
		hasVal := true
		firstSpaceIndex := strings.Index(s, " ")
		if firstSpaceIndex == -1 {
			hasVal = false
			firstSpaceIndex = len(s)
		}

		annoNameFragment := lo.Substring(s, 0, uint(firstSpaceIndex))
		if annoNameFragment == findAnno {
			if hasVal {
				val = lo.Substring(s, firstSpaceIndex, uint(len(s)))
				val = strings.TrimSpace(val)
				return val, true
			}
			return "", true
		}
	}
	return "", false
}
