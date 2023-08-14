package generator

import (
	"bytes"
	"fmt"
	"github.com/Masterminds/sprig/v3"
	"github.com/ghoulhyk/dao-gen/bean"
	"github.com/ghoulhyk/dao-gen/conf"
	"github.com/ghoulhyk/dao-gen/conf/confBean"
	"github.com/ghoulhyk/dao-gen/tmpl"
	"github.com/ghoulhyk/dao-gen/utils/tmplUtils"
	"github.com/samber/lo"
	"go/format"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func generateFragmentTmpl(basicPath string, tableInfoList []bean.TableInfo) {
	for _, info := range tableInfoList {
		fmt.Printf("\n==============   %s   ==============\n", info.TableName())
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

	// region 拼接模板 srcFileContent

	srcFileContent, err := fs.ReadFile(tmpl.FragmentTemplateFs, mainTmplFilePath)
	if err != nil {
		panic(err)
	}

	hasSysAppendTmpl := func() bool {
		subTmplFileDir, err := tmpl.FragmentTemplateFs.Open(subTmplFilePath)
		// err 不为空，代表 文件夹不存在
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
	if hasSysAppendTmpl {
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

	hasUserCustomAppendTmpl := func() bool {
		if conf.GetIns().Path2basic.AppendTmpl == "" {
			return false
		}
		subTmplFileDir, err := os.Open(filepath.Join(basicPath, conf.GetIns().Path2basic.AppendTmpl, srcDir))
		// err 不为空，代表 文件夹不存在
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
	if hasUserCustomAppendTmpl {
		dirEntrys, err := os.ReadDir(filepath.Join(basicPath, conf.GetIns().Path2basic.AppendTmpl, srcDir))
		if err == nil && len(dirEntrys) > 0 {
			var appendTemplContentList [][]byte
			for _, entry := range dirEntrys {
				if entry.IsDir() {
					continue
				}
				appendTemplContent, err := os.ReadFile(filepath.Join(basicPath, conf.GetIns().Path2basic.AppendTmpl, srcDir, entry.Name()))
				if err == nil {
					appendTemplContentList = append(appendTemplContentList, appendTemplContent)
				}
			}
			if len(appendTemplContentList) > 0 {
				srcFileContent = tmplUtils.Append(srcFileContent, appendTemplContentList...)
			}
		}
	}

	// endregion

	tpl := template.Must(
		template.New("main").
			Funcs(sprig.TxtFuncMap()).
			Parse(string(srcFileContent)),
	)

	source := bytes.Buffer{}

	err = tpl.Execute(&source, data)
	if err != nil {
		fmt.Printf("\n%s 模板替换失败!!!\n\n", srcDir)
		panic(err)
	}

	sourceBytes := source.Bytes()

	// 格式化源代码
	formattedSource, err := format.Source(sourceBytes)
	if err != nil {
		fmt.Printf("\n%s 格式化失败!!!\n\n", srcDir)
		fmt.Printf("\n%s\n\n", string(sourceBytes))
		panic(err)
	}

	err = os.WriteFile(filepath.Join(basicPath, dstPath, dstFile), formattedSource, 0666)
	if err != nil {
		fmt.Printf("\n%s/%s 写入失败!!!\n\n", dstPath, dstFile)
		panic(err)
	}
	fmt.Printf("%s/%s 写入成功\n", dstPath, dstFile)
}
