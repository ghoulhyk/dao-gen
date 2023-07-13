package generator

import (
	"daogen/bean"
	"github.com/samber/lo"
)

// 本项用于 bulkInserterGroup.tmpl、deleterGroup.tmpl、inserterGroup.tmpl、selectorGroup.tmpl、updaterGroup.tmpl
var tables []map[string]any

var commonDataItems map[string]any

func initialize(basicPath string, tableList []bean.TableInfo) {
	tables = lo.Map(tableList, func(item bean.TableInfo, index int) map[string]any {
		return map[string]any{
			"name":                          item.StructName(),
			"inserterWrapperStructName":     item.InserterWrapperStructName(),
			"bulkInserterWrapperStructName": item.BulkInserterWrapperStructName(),
			"selectorWrapperStructName":     item.SelectorWrapperStructName(),
			"updaterWrapperStructName":      item.UpdaterWrapperStructName(),
			"deleterWrapperStructName":      item.DeleterWrapperStructName(),
		}
	})
	commonDataItems = commonDataItem(basicPath)
}
