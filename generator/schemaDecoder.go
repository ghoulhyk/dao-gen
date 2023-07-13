package generator

import (
	"daogen/bean"
	"daogen/utils"
	"daogen/utils/annotationUtils"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/samber/lo"
	"go/token"
	"io/fs"
	"path/filepath"
	"strings"
)

func decodeTableList(tableDirPath string) []bean.TableInfo {
	var result []bean.TableInfo
	jumpRoot := false
	filepath.Walk(tableDirPath, func(path string, info fs.FileInfo, err error) error {
		if !jumpRoot {
			jumpRoot = true
			return nil
		}
		if info.IsDir() {
			return nil
		}
		tableInfo := decodeTable(path)
		if tableInfo != nil {
			result = append(result, *tableInfo)
		}
		return nil
	})
	return result
}

func decodeTable(tableDefPath string) *bean.TableInfo {
	file, err := decorator.ParseFile(token.NewFileSet(), tableDefPath, nil, 0)
	if err != nil {
		panic("定义文件解析失败" + tableDefPath)
		return nil
	}

	structDecl, hasTable := lo.Find(file.Decls, func(item dst.Decl) bool {
		startDecl := item.Decorations().Start
		return annotationUtils.ExistTableName(startDecl)
	})
	if !hasTable {
		return nil
	}

	otherDecl := lo.Filter(file.Decls, func(item dst.Decl, index int) bool {
		genDecl, ok := item.(*dst.GenDecl)
		if !ok {
			return false
		}
		if genDecl.Tok == token.IMPORT {
			return false
		}
		startDecl := item.Decorations().Start
		isTable := annotationUtils.ExistTableName(startDecl)
		if isTable {
			return false
		}
		return true
	})

	imports := file.Imports
	startDecl := structDecl.Decorations().Start
	structName := structDecl.(*dst.GenDecl).Specs[0].(*dst.TypeSpec).Name.Name
	tableName, _ := annotationUtils.GetTableName(startDecl)
	databaseName, _ := annotationUtils.GetDatabaseName(startDecl)
	packageName := file.Name.Name
	fieldList := structDecl.(*dst.GenDecl).Specs[0].(*dst.TypeSpec).Type.(*dst.StructType).Fields.List
	columnList := lo.Map(fieldList, func(item *dst.Field, index int) bean.TableColumn {
		return decodeTableColumn(item)
	})
	tableInfo := bean.TableInfo{}
	tableInfo.
		SetImports(imports).
		SetStructName(structName).
		SetTableName(tableName).
		SetDatabaseName(databaseName).
		SetStructInfo(structDecl.(*dst.GenDecl)).
		SetPackageName(packageName).
		SetOtherDecls(otherDecl).
		SetColumnList(columnList)

	return &tableInfo
}

func decodeTableColumn(field *dst.Field) bean.TableColumn {
	fieldName := field.Names[0].Name
	fieldType := utils.GetTypeName(field.Type)
	if strings.HasPrefix(fieldType, "*") {
		fieldType = strings.TrimLeft(fieldType, "*")
	}
	tag := field.Tag.Value
	startAnnotations := field.Decs.Start
	endAnnotations := field.Decs.End

	table, err := bean.NewTableColumn(fieldName, fieldType, tag, startAnnotations, endAnnotations)
	if err != nil {
		panic(err.Error())
	}
	return *table
}
