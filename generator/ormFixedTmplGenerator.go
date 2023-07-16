package generator

import (
	"bytes"
	"fmt"
	"github.com/Masterminds/sprig/v3"
	"github.com/ghoulhyk/dao-gen/conf"
	"github.com/ghoulhyk/dao-gen/conf/confBean"
	"github.com/ghoulhyk/dao-gen/tmpl"
	"github.com/samber/lo"
	"go/format"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

func generateOrmFixedTmpl(basicPath string) {

	ormType := conf.GetIns().OrmInfo.Type
	fixedTmplDir := fmt.Sprintf("orm/%s/fixed", ormType)

	err := fs.WalkDir(tmpl.OrmTemplateFs, fixedTmplDir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		parentDir, fileName := filepath.Split(path)
		parentDir = strings.TrimPrefix(parentDir, fixedTmplDir)
		parentDir = strings.Trim(parentDir, "/")
		fileName = strings.TrimSuffix(fileName, ".tmpl") + ".go"

		packageInfoList := conf.AllPackageInfoList(basicPath)
		// key: tmpl/file 文件夹名称; value: 生成的文件夹路径
		realParentDirMap := lo.SliceToMap(packageInfoList, func(item confBean.PackageInfo) (string, string) {
			return item.DirTmplFileRelative(), item.DirRelativePath()
		})

		parentDir = realParentDirMap[parentDir]

		generateOrmFixedTmplItem(basicPath, path, parentDir, fileName)
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func generateOrmFixedTmplItem(basicPath string, srcFile string, dstPath string, dstFile string) {
	srcFileName := path.Base(srcFile)
	tpl := template.Must(
		template.New(srcFileName).
			Funcs(sprig.TxtFuncMap()).
			ParseFS(tmpl.OrmTemplateFs, srcFile),
	)

	commonData := lo.Assign(commonDataItems)
	commonData["databaseInfos"] = conf.GetIns().Database.DatabaseInfos

	source := bytes.Buffer{}

	err := tpl.Execute(&source, commonData)
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
