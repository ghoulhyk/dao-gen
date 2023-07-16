package generator

import (
	"bytes"
	"github.com/Masterminds/sprig/v3"
	"github.com/ghoulhyk/dao-gen/bean"
	"github.com/ghoulhyk/dao-gen/conf"
	"github.com/ghoulhyk/dao-gen/conf/confBean"
	"github.com/ghoulhyk/dao-gen/tmpl"
	"github.com/samber/lo"
	"go/format"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func generateOrmFragmentTmpl(basicPath string, tableInfoList []bean.TableInfo) {
	for _, info := range tableInfoList {
		generateOrmFragmentTmplItem(basicPath, conf.GetWhereCondPackageInfo(basicPath), info.ModelBeanName()+".go", info)
	}
}

func generateOrmFragmentTmplItem(basicPath string, pkgInfo confBean.PackageInfo, dstFile string, tableInfo bean.TableInfo) {
	ormType := conf.GetIns().OrmInfo.Type
	srcDir := "orm/" + ormType + "/fragment/" + pkgInfo.DirTmplFileRelative()
	dstPath := pkgInfo.DirRelativePath()
	mainTmplFilePath := strings.ReplaceAll(filepath.Join(srcDir, "main.tmpl"), "\\", "/")
	subTmplFilePath := strings.ReplaceAll(filepath.Join(srcDir, "subTmpl"), "\\", "/")

	commonData := lo.Assign(commonDataItems)
	tableData := tableDataItem(tableInfo)
	var fieldDataList []map[string]any
	for _, column := range tableInfo.ColumnList() {
		fieldData := lo.Assign(commonData, tableData, columnDataItem(column))
		fieldDataList = append(fieldDataList, fieldData)
	}
	data := lo.Assign(commonData, tableData)
	data["tableImports"] = tableInfo.ImportsStr()
	data["column"] = fieldDataList

	srcFileContent, err := fs.ReadFile(tmpl.OrmTemplateFs, mainTmplFilePath)
	if err != nil {
		panic(err)
	}

	doAppendSubTempl := func() bool {
		subTmplFileDir, err := tmpl.OrmTemplateFs.Open(subTmplFilePath)
		// err 为空代表，文件夹不存在
		if err != nil {
			return false
		}
		info, err := subTmplFileDir.Stat()
		if err != nil {
			return false
		}
		if !info.IsDir() {
			return false
		}
		return true
	}()
	if doAppendSubTempl {
		err = fs.WalkDir(tmpl.OrmTemplateFs, subTmplFilePath, func(path string, d fs.DirEntry, err error) error {
			if d.IsDir() {
				return nil
			}

			appendTemplContent, err := fs.ReadFile(tmpl.OrmTemplateFs, path)
			if err != nil {
				panic(err)
			}

			srcFileContent = append(srcFileContent, '\n')
			srcFileContent = append(srcFileContent, appendTemplContent...)

			return nil
		})
		if err != nil {
			panic(err)
		}
	}

	tpl := template.Must(
		template.New("main").
			Funcs(sprig.TxtFuncMap()).
			Parse(string(srcFileContent)),
	)

	source := bytes.Buffer{}

	err = tpl.Execute(&source, data)
	if err != nil {
		panic(err)
	}

	sourceBytes := source.Bytes()

	// 格式化源代码
	formattedSource, err := format.Source(sourceBytes)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filepath.Join(basicPath, dstPath, dstFile), formattedSource, 0666)
	if err != nil {
		panic(err)
	}
}
