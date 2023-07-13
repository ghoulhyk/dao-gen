package generator

import (
	"fmt"
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
	tpl := template.Must(
		template.ParseFS(tmpl.OrmTemplateFs, srcFile),
	)

	commonData := lo.Assign(commonDataItems)
	commonData["databaseInfos"] = conf.GetIns().Database.DatabaseInfos

	file, err := os.OpenFile(filepath.Join(basicPath, dstPath, dstFile), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = tpl.Execute(file, commonData)
	if err != nil {
		panic(err)
	}
}
