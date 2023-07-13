package generator

import (
	"github.com/Masterminds/sprig/v3"
	"github.com/ghoulhyk/dao-gen/bean"
	"github.com/ghoulhyk/dao-gen/conf"
	"github.com/ghoulhyk/dao-gen/conf/confBean"
	"github.com/ghoulhyk/dao-gen/tmpl"
	"github.com/samber/lo"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func generateFragmentTmpl(basicPath string, tableInfoList []bean.TableInfo) {
	for _, info := range tableInfoList {
		generateFragmentTmplItem(basicPath, conf.GetDeleterPackageInfo(basicPath), info.ModelBeanName()+".go", info)
		generateFragmentTmplItem(basicPath, conf.GetFieldsPackageInfo(basicPath), info.ModelBeanName()+".go", info)
		generateFragmentTmplItem(basicPath, conf.GetInserterPackageInfo(basicPath), info.ModelBeanName()+".go", info)
		generateFragmentTmplItem(basicPath, conf.GetModelPackageInfo(basicPath), info.ModelBeanName()+".go", info)
		generateFragmentTmplItem(basicPath, conf.GetOrderCondPackageInfo(basicPath), info.ModelBeanName()+".go", info)
		generateFragmentTmplItem(basicPath, conf.GetSelectorPackageInfo(basicPath), info.ModelBeanName()+".go", info)
		generateFragmentTmplItem(basicPath, conf.GetUpdaterPackageInfo(basicPath), info.ModelBeanName()+".go", info)
	}
}

func generateFragmentTmplItem(basicPath string, pkgInfo confBean.PackageInfo, dstFile string, tableInfo bean.TableInfo) {
	srcDir := "fragment/" + pkgInfo.DirTmplFileRelative()
	dstPath := pkgInfo.DirRelativePath()
	mainTmplFilePath := strings.ReplaceAll(filepath.Join(srcDir, "main.tmpl"), "\\", "/")
	subTmplFilePath := strings.ReplaceAll(filepath.Join(srcDir, "subTmpl"), "\\", "/")
	file, err := os.OpenFile(filepath.Join(basicPath, dstPath, dstFile), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

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

	srcFileContent, err := fs.ReadFile(tmpl.FragmentTemplateFs, mainTmplFilePath)
	if err != nil {
		panic(err)
	}

	doAppendSubTempl := func() bool {
		subTmplFileDir, err := tmpl.FragmentTemplateFs.Open(subTmplFilePath)
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
		err = fs.WalkDir(tmpl.FragmentTemplateFs, subTmplFilePath, func(path string, d fs.DirEntry, err error) error {
			if d.IsDir() {
				return nil
			}

			appendTemplContent, err := fs.ReadFile(tmpl.FragmentTemplateFs, path)
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

	err = tpl.Execute(file, data)
	if err != nil {
		panic(err)
	}

}
