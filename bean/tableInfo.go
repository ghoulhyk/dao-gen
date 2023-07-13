package bean

import (
	"fmt"
	"github.com/dave/dst"
	"github.com/ghoulhyk/dao-gen/conf"
	"github.com/ghoulhyk/dao-gen/conf/confBean"
	"strings"
)

type TableInfo struct {
	imports      []*dst.ImportSpec // 原有的import
	structName   string            // 类名
	tableName    string            // 表名
	databaseName string            // 库名
	structInfo   *dst.GenDecl      // 原 Struct 结构体
	packageName  string            // 原包名，用于 joinSelector 生成时初始化判断

	otherDecls []dst.Decl // 原有的其它除表主结构（StructInfo）外，其他的定义

	columnFieldInfoList TableColumnList // 表内字段数组，由GetFieldInfoList方法生成
}

// region setter 初始化使用

func (receiver *TableInfo) SetImports(imports []*dst.ImportSpec) *TableInfo {
	receiver.imports = imports
	return receiver
}
func (receiver *TableInfo) SetStructName(structName string) *TableInfo {
	receiver.structName = structName
	return receiver
}
func (receiver *TableInfo) SetTableName(tableName string) *TableInfo {
	receiver.tableName = tableName
	return receiver
}
func (receiver *TableInfo) SetDatabaseName(databaseName string) *TableInfo {
	receiver.databaseName = databaseName
	return receiver
}
func (receiver *TableInfo) SetStructInfo(structInfo *dst.GenDecl) *TableInfo {
	receiver.structInfo = structInfo
	return receiver
}
func (receiver *TableInfo) SetPackageName(packageName string) *TableInfo {
	receiver.packageName = packageName
	return receiver
}
func (receiver *TableInfo) SetOtherDecls(otherDecls []dst.Decl) *TableInfo {
	receiver.otherDecls = otherDecls
	return receiver
}
func (receiver *TableInfo) SetColumnList(columnList TableColumnList) *TableInfo {
	receiver.columnFieldInfoList = columnList
	return receiver
}

// endregion

// region getter

func (receiver *TableInfo) Imports() []*dst.ImportSpec {
	return receiver.imports
}

// ImportsStr 导包的具体文字
func (receiver *TableInfo) ImportsStr() []string {
	result := []string{}
	for _, spec := range receiver.Imports() {
		importStr := ""
		if spec.Name.Name != "" {
			importStr += spec.Name.Name
			importStr += " "
		}
		importStr += spec.Path.Value
		result = append(result, importStr)
	}
	return result
}
func (receiver *TableInfo) TableName() string {
	return receiver.tableName
}
func (receiver *TableInfo) DatabaseName() string {
	return receiver.databaseName
}
func (receiver *TableInfo) StructInfo() *dst.GenDecl {
	return receiver.structInfo
}
func (receiver *TableInfo) PackageName() string {
	return receiver.packageName
}
func (receiver *TableInfo) OtherDecls() []dst.Decl {
	return receiver.otherDecls
}
func (receiver *TableInfo) ColumnList() TableColumnList {
	return receiver.columnFieldInfoList
}

// endregion

// region 模板中的相关参数: {{.xxx}}

// StructName
// 原类名
func (receiver *TableInfo) StructName() string {
	return receiver.structName
}

// StructNameUntitled
// 首字母小写的原类名
func (receiver *TableInfo) StructNameUntitled() string {
	structName := receiver.StructName()
	return strings.ToLower(string(structName[0])) + structName[1:]
}

// ModelBeanName
// 生成的实体类名称
func (receiver *TableInfo) ModelBeanName() string {
	return receiver.StructName()
}

// InserterDataModelName
// 生成的插入时所使用的实体类名称
func (receiver *TableInfo) InserterDataModelName() string {
	return receiver.StructNameUntitled() + "DataModel"
}

// region Params

func (receiver *TableInfo) SelectorParamsStructName() string {
	return receiver.StructNameUntitled() + "SelectorParams"
}

func (receiver *TableInfo) InserterParamsStructName() string {
	return receiver.StructNameUntitled() + "InserterParams"
}

func (receiver *TableInfo) BulkInserterParamsStructName() string {
	return receiver.StructNameUntitled() + "BulkInserterParams"
}

func (receiver *TableInfo) DeleterParamsStructName() string {
	return receiver.StructNameUntitled() + "DeleterParams"
}

func (receiver *TableInfo) UpdaterParamsStructName() string {
	return receiver.StructNameUntitled() + "UpdaterParams"
}

// endregion

// region Wrapper

func (receiver *TableInfo) SelectorWrapperStructName() string {
	return receiver.StructName() + "Wrapper"
}

func (receiver *TableInfo) InserterWrapperStructName() string {
	return receiver.StructName() + "Wrapper"
}

func (receiver *TableInfo) BulkInserterWrapperStructName() string {
	return receiver.StructName() + "BulkWrapper"
}

func (receiver *TableInfo) DeleterWrapperStructName() string {
	return receiver.StructName() + "Wrapper"
}

func (receiver *TableInfo) UpdaterWrapperStructName() string {
	return receiver.StructName() + "Wrapper"
}

// endregion

// region WhereCond

// WhereCondStructName
// whereCond 外层的类名
func (receiver *TableInfo) WhereCondStructName() string {
	return receiver.StructName()
}

// WhereCondInnerStructName
// whereCond 内层的类名
func (receiver *TableInfo) WhereCondInnerStructName() string {
	return receiver.StructName() + "_Inner"
}

// WhereCondStructConstructorName
// whereCond 外层类 构造函数名
func (receiver *TableInfo) WhereCondStructConstructorName() string {
	return fmt.Sprintf("New_%s", receiver.StructName())
}

// WhereCondInnerStructConstructorName
// whereCond 内层类 构造函数名
func (receiver *TableInfo) WhereCondInnerStructConstructorName() string {
	return fmt.Sprintf("New_%s_Inner", receiver.StructName())
}

// OrderCondStructName
// orderCond 类名
func (receiver *TableInfo) OrderCondStructName() string {
	return receiver.StructName()
}

// DatabaseDefFieldName
// 本表的库名在 databaseDef 包中对应的字段名
func (receiver *TableInfo) DatabaseDefFieldName() string {
	if receiver.DatabaseName() == "" {
		return ""
	}
	databaseInfoConf := receiver.GetDatabaseInfo()
	if databaseInfoConf == nil {
		return ""
	}
	return databaseInfoConf.FieldName()
}

// endregion

// endregion

// GetDatabaseInfo
// 获取库信息
func (receiver *TableInfo) GetDatabaseInfo() *confBean.DatabaseItemConf {
	return conf.GetIns().Database.FindDatabaseConfByName(receiver.DatabaseName())
}
