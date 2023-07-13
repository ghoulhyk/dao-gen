package utils

import (
	"github.com/samber/lo"
	"strings"
)

var innerCustomTypeListHolder []string

func AddInnerCustomType(list []string) {
	innerCustomTypeListHolder = append(innerCustomTypeListHolder, list...)
}

func InnerCustomTypeList() []string {
	return innerCustomTypeListHolder
}

func IsInnerCustomType(fileType string) bool {
	fileType = strings.Trim(fileType, "*")
	return lo.Contains(innerCustomTypeListHolder, fileType)
}
