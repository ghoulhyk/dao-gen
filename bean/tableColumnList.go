package bean

import (
	"github.com/samber/lo"
)

type TableColumnList []TableColumn

func (receiver TableColumnList) WithoutAutoincrPk() TableColumnList {
	result := lo.Filter(receiver, func(v TableColumn, i int) bool {
		return !(v.IsPk() && v.IsAutoincr())
	})
	return result
}

func (receiver TableColumnList) AutoIncrPkFieldInfo() (TableColumn, bool) {
	return lo.Find[TableColumn](receiver, func(v TableColumn) bool {
		return v.IsPk() && v.IsAutoincr()
	})
}

func (receiver TableColumnList) HasAutoIncrPk() bool {
	_, exist := receiver.AutoIncrPkFieldInfo()
	return exist
}

func (receiver TableColumnList) HasInnerCustomType() bool {
	return lo.ContainsBy(receiver, func(v TableColumn) bool {
		return v.IsInnerCustomType()
	})
}
