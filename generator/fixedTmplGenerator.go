package generator

import (
	"bytes"
	"fmt"
	"github.com/Masterminds/sprig/v3"
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

func generateFixedTmpl(basicPath string) {

	fmt.Printf("\n==============   fixed   ==============\n")

	err := fs.WalkDir(tmpl.FixedTemplateFs, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		parentDir, fileName := filepath.Split(path)
		parentDir = strings.TrimPrefix(parentDir, "fixed/")
		parentDir = strings.TrimSuffix(parentDir, "/")
		fileName = strings.TrimSuffix(fileName, ".tmpl") + ".go"

		packageInfoList := conf.AllPackageInfoList(basicPath)
		// key: tmpl/file 文件夹名称; value: 生成的文件夹路径
		realParentDirMap := lo.SliceToMap(packageInfoList, func(item confBean.PackageInfo) (string, string) {
			return item.DirTmplFileRelative(), item.DirRelativePath()
		})

		parentDir = realParentDirMap[parentDir]

		generateFixedTmplItem(basicPath, path, parentDir, fileName)
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func generateFixedTmplItem(basicPath string, srcFile string, dstPath string, dstFile string) {
	commonData := lo.Assign(commonDataItems)
	commonData["databaseInfos"] = conf.GetIns().Database.DatabaseInfos
	commonData["tables"] = tables

	// region 拼接模板 srcFileContent

	srcFileContent, err := fs.ReadFile(tmpl.FixedTemplateFs, srcFile)
	if err != nil {
		panic(err)
	}

	if conf.GetIns().Path2basic.AppendTmpl != "" {
		appendTmplContent, err := os.ReadFile(filepath.Join(basicPath, conf.GetIns().Path2basic.AppendTmpl, srcFile))
		if err == nil {
			srcFileContent = tmplUtils.Append(srcFileContent, appendTmplContent)
		}
	}

	// endregion

	tpl := template.Must(
		template.New("main").
			Funcs(sprig.TxtFuncMap()).
			Parse(string(srcFileContent)),
	)

	source := bytes.Buffer{}

	err = tpl.Execute(&source, commonData)
	if err != nil {
		fmt.Printf("\n%s 模板替换失败!!!\n\n", srcFile)
		panic(err)
	}

	sourceBytes := source.Bytes()

	// 格式化源代码
	formattedSource, err := format.Source(sourceBytes)
	if err != nil {
		fmt.Printf("\n%s 格式化失败!!!\n\n", srcFile)
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
