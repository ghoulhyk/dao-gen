package bean

import (
	"errors"
	"fmt"
	"github.com/dave/dst"
	"github.com/ghoulhyk/dao-gen/conf"
	"github.com/ghoulhyk/dao-gen/utils"
	"github.com/samber/lo"
	"reflect"
	"regexp"
	"strings"
)

type TableColumn struct {

	// region 外部传入

	fieldName        string // 字段名
	fieldType        string // 字段类型
	tag              string // json:"id" colAttr:"'id' pk autoincr"
	startAnnotations dst.Decorations
	endAnnotations   dst.Decorations

	// endregion

	endAnnotationStr string
	columnTag        string // colAttr:"'id' pk autoincr" 中的 'id' pk autoincr
	columnName       string // 数据库中的列名 ; colAttr:"'id' pk autoincr" 中的 id
}

type JoinTableColumn struct {
	TableColumn

	// region 外部传入

	isTableModel bool       // joinSelector 使用，标识本字段是不是一个表
	tableInfo    *TableInfo // joinSelector 使用，本字段关联的表
	isMainTable  bool       // joinSelector 使用，标识本字段是不是join的主表

	// endregion

}

func NewTableColumn(fieldName string, fieldType string, tag string, startAnnotations dst.Decorations, endAnnotations dst.Decorations) (*TableColumn, error) {
	endAnnotationStr := strings.Join(endAnnotations.All(), "\t")
	columnTag := reflect.StructTag(tag).Get("colAttr")
	if columnTag == "" {
		return nil, errors.New(fmt.Sprintf("[%s]字段无 colAttr tag [%s]", fieldName, tag))
	}
	columnName := regexp.MustCompile("'([\\s\\S]+)'").FindStringSubmatch(columnTag)[1]
	return &TableColumn{
		fieldName:        fieldName,
		fieldType:        fieldType,
		tag:              tag,
		startAnnotations: startAnnotations,
		endAnnotations:   endAnnotations,

		endAnnotationStr: endAnnotationStr,
		columnTag:        columnTag,
		columnName:       columnName,
	}, nil
}

func (receiver TableColumn) EndAnnotationStr() string {
	return receiver.endAnnotationStr
}

func (receiver TableColumn) StartAnnotations() dst.Decorations {
	return receiver.startAnnotations
}

func (receiver TableColumn) EndAnnotations() dst.Decorations {
	return receiver.endAnnotations
}

func (receiver TableColumn) Tag() string {
	return receiver.tag
}

func (receiver TableColumn) FieldName() string {
	return receiver.fieldName
}

func (receiver TableColumn) UntitledFieldName() string {
	result := utils.UnCapitalize(receiver.fieldName)
	if lo.Contains(utils.GoKeywordList(), result) {
		result += "x"
	}
	return result
}

func (receiver TableColumn) TitledFieldName() string {
	return utils.Capitalize(receiver.fieldName)
}

// FieldType
// withDoPkg : 如果是自定义的类，是否带上实体类包的包名
func (receiver TableColumn) FieldType(isPointer bool, withModelPkg bool) string {
	fileType := receiver.fieldType
	//if utils.IsEnum(fileType) && withModelPkg {
	//	fileType = conf.GetModelPackageInfo("").PackageNameForRef() + "." + fileType
	//}
	if isPointer {
		return "*" + fileType
	}
	return fileType
}

func (receiver TableColumn) ColumnName() string {
	return receiver.columnName
}

func (receiver TableColumn) IsPk() bool {
	columnTagAttrList := strings.Split(receiver.columnTag, " ")
	return lo.Contains(columnTagAttrList, "pk")
}

func (receiver TableColumn) IsAutoincr() bool {
	columnTagAttrList := strings.Split(receiver.columnTag, " ")
	return lo.Contains(columnTagAttrList, "autoincr")
}

func (receiver TableColumn) IsNullable() bool {
	columnTagAttrList := strings.Split(receiver.columnTag, " ")
	return lo.Contains(columnTagAttrList, "nullable")
}

// GetXormTag
// 实体类中字段的xorm标签
func (receiver TableColumn) GetXormTag() string {
	//if receiver.IsTableModel {
	//	return `xorm:"extends"`
	//}
	tag := "`xorm:\"'" + receiver.ColumnName() + "'"
	if receiver.IsPk() {
		tag += " pk"
	}
	if receiver.IsAutoincr() {
		tag += " autoincr"
	}
	tag += "\"`"
	return tag
}

// ModelFieldTag
// 实体类中字段的标签
func (receiver TableColumn) ModelFieldTag() string {
	colAttrRegexp := utils.GetColAttrRegexp()
	tagValue := receiver.Tag()
	// 删除 colAttr:"xxx" 部分
	tagValue = colAttrRegexp.ReplaceAllString(tagValue, " ")
	tagValue = strings.Trim(tagValue, "`")
	tagValue = strings.TrimSpace(tagValue)

	if conf.GetIns().OrmInfo.IsXorm() {
		xormTag := receiver.GetXormTag()
		xormTag = strings.Trim(xormTag, "`")
		return fmt.Sprintf("`%s %s`", tagValue, xormTag)
	}

	return ""
}

//func (receiver TableColumn) JoinCondFieldName() string {
//	condFieldName := receiver.FieldName
//	if condFieldName == "" {
//		condFieldName = receiver.GetFieldType()
//		condFieldName, _ = lo.Last(strings.Split(condFieldName, "."))
//	}
//	return condFieldName
//}
//
//func (receiver TableColumn) TableNameOrAlias() string {
//	condFieldName := receiver.FieldName
//	if condFieldName != "" {
//		return strings.ToLower(string(condFieldName[0])) + condFieldName[1:]
//	}
//	return receiver.TableInfo.TableName()
//}
//
//func (receiver TableColumn) UntitledJoinCondFieldName() string {
//	condFieldName := receiver.JoinCondFieldName()
//	return strings.ToLower(string(condFieldName[0])) + condFieldName[1:]
//}
//
//func (receiver TableColumn) JoinCondOp() string {
//	if annotationUtils.Exist(receiver.StartAnnotations, "@leftJoin") {
//		return "LEFT"
//	}
//	if annotationUtils.Exist(receiver.StartAnnotations, "@rightJoin") {
//		return "RIGHT"
//	}
//	if annotationUtils.Exist(receiver.StartAnnotations, "@innerJoin") {
//		return "INNER"
//	}
//	return ""
//}
//
//func (receiver TableColumn) TitledJoinCondOp() string {
//	op := receiver.JoinCondOp()
//	if op == "" {
//		return ""
//	}
//	return strings.ToUpper(string(op[0])) + strings.ToLower(op[1:])
//}
//
//func (receiver TableColumn) JoinCondSql() string {
//	switch receiver.JoinCondOp() {
//	case "LEFT":
//		val, _ := annotationUtils.FindStr(receiver.StartAnnotations, "@leftJoin")
//		return val
//	case "RIGHT":
//		val, _ := annotationUtils.FindStr(receiver.StartAnnotations, "@rightJoin")
//		return val
//	case "INNER":
//		val, _ := annotationUtils.FindStr(receiver.StartAnnotations, "@innerJoin")
//		return val
//	}
//	return ""
//}

// IsInnerCustomType
// 是否是在本包内定义的类
func (receiver TableColumn) IsInnerCustomType() bool {
	return utils.IsInnerCustomType(receiver.FieldType(false, false))
}
