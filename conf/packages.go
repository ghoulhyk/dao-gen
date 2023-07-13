package conf

import (
	"daogen/conf/confBean"
	"daogen/utils"
	"github.com/samber/lo"
	"os"
	"path/filepath"
	"strings"
)

var packageInfoHolder = map[string]confBean.PackageInfo{}

func AllPackageInfoList(basicPath string) []confBean.PackageInfo {
	var result []confBean.PackageInfo
	result = append(result, GetModelPackageInfo(basicPath))
	result = append(result, GetUtilPackageInfo(basicPath))
	result = append(result, GetErrorsPackageInfo(basicPath))
	result = append(result, GetBasePackageInfo(basicPath))
	result = append(result, GetBaseModelPackageInfo(basicPath))
	result = append(result, GetBaseCondPackageInfo(basicPath))
	result = append(result, GetWhereCondPackageInfo(basicPath))
	result = append(result, GetOrderCondPackageInfo(basicPath))
	result = append(result, GetFieldsPackageInfo(basicPath))
	result = append(result, GetInserterPackageInfo(basicPath))
	result = append(result, GetDeleterPackageInfo(basicPath))
	result = append(result, GetUpdaterPackageInfo(basicPath))
	result = append(result, GetSelectorPackageInfo(basicPath))
	result = append(result, GetJoinSelectorPackageInfo(basicPath))
	result = append(result, GetJoinCondPackageInfo(basicPath))
	result = append(result, GetDatabaseDefPackageInfo(basicPath))
	return result
}

func GetModelPackageInfo(basicPath string) confBean.PackageInfo {
	relativePath := "model/do"
	importName := ""
	key := "model"
	tmplFileRelativeDir := "model"
	return createPackageInfo(basicPath, relativePath, importName, tmplFileRelativeDir, key)
}

func GetUtilPackageInfo(basicPath string) confBean.PackageInfo {
	relativePath := "util"
	importName := ""
	key := "util"
	tmplFileRelativeDir := "util"
	return createPackageInfo(basicPath, relativePath, importName, tmplFileRelativeDir, key)
}

func GetErrorsPackageInfo(basicPath string) confBean.PackageInfo {
	basicDbDirName := filepath.Base(basicPath)
	relativePath := basicDbDirName + "errors"
	importName := ""
	key := "errors"
	tmplFileRelativeDir := "errors"
	return createPackageInfo(basicPath, relativePath, importName, tmplFileRelativeDir, key)
}

func GetBasePackageInfo(basicPath string) confBean.PackageInfo {
	relativePath := "base"
	importName := ""
	key := "base"
	tmplFileRelativeDir := "base"
	return createPackageInfo(basicPath, relativePath, importName, tmplFileRelativeDir, key)
}

func GetBaseModelPackageInfo(basicPath string) confBean.PackageInfo {
	relativePath := "baseModel"
	importName := ""
	key := "baseModel"
	tmplFileRelativeDir := "baseModel"
	return createPackageInfo(basicPath, relativePath, importName, tmplFileRelativeDir, key)
}

func GetBaseCondPackageInfo(basicPath string) confBean.PackageInfo {
	relativePath := "baseCond"
	importName := ""
	key := "baseCond"
	tmplFileRelativeDir := "baseCond"
	return createPackageInfo(basicPath, relativePath, importName, tmplFileRelativeDir, key)
}

func GetWhereCondPackageInfo(basicPath string) confBean.PackageInfo {
	relativePath := "whereCond"
	importName := ""
	key := "whereCond"
	tmplFileRelativeDir := "whereCond"
	return createPackageInfo(basicPath, relativePath, importName, tmplFileRelativeDir, key)
}

func GetOrderCondPackageInfo(basicPath string) confBean.PackageInfo {
	relativePath := "orderCond"
	importName := ""
	key := "orderCond"
	tmplFileRelativeDir := "orderCond"
	return createPackageInfo(basicPath, relativePath, importName, tmplFileRelativeDir, key)
}

func GetFieldsPackageInfo(basicPath string) confBean.PackageInfo {
	relativePath := "fields"
	importName := ""
	key := "fields"
	tmplFileRelativeDir := "fields"
	return createPackageInfo(basicPath, relativePath, importName, tmplFileRelativeDir, key)
}

func GetInserterPackageInfo(basicPath string) confBean.PackageInfo {
	relativePath := "inserter"
	importName := ""
	key := "inserter"
	tmplFileRelativeDir := "inserter"
	return createPackageInfo(basicPath, relativePath, importName, tmplFileRelativeDir, key)
}

func GetDeleterPackageInfo(basicPath string) confBean.PackageInfo {
	relativePath := "deleter"
	importName := ""
	key := "deleter"
	tmplFileRelativeDir := "deleter"
	return createPackageInfo(basicPath, relativePath, importName, tmplFileRelativeDir, key)
}

func GetUpdaterPackageInfo(basicPath string) confBean.PackageInfo {
	relativePath := "updater"
	importName := ""
	key := "updater"
	tmplFileRelativeDir := "updater"
	return createPackageInfo(basicPath, relativePath, importName, tmplFileRelativeDir, key)
}

func GetSelectorPackageInfo(basicPath string) confBean.PackageInfo {
	relativePath := "selector"
	importName := ""
	key := "selector"
	tmplFileRelativeDir := "selector"
	return createPackageInfo(basicPath, relativePath, importName, tmplFileRelativeDir, key)
}

func GetJoinSelectorPackageInfo(basicPath string) confBean.PackageInfo {
	relativePath := "joinSelector"
	importName := ""
	key := "joinSelector"
	tmplFileRelativeDir := "joinSelector"
	return createPackageInfo(basicPath, relativePath, importName, tmplFileRelativeDir, key)
}

func GetJoinCondPackageInfo(basicPath string) confBean.PackageInfo {
	relativePath := "joinCond"
	importName := ""
	key := "joinCond"
	tmplFileRelativeDir := "joinCond"
	return createPackageInfo(basicPath, relativePath, importName, tmplFileRelativeDir, key)
}

func GetDatabaseDefPackageInfo(basicPath string) confBean.PackageInfo {
	relativePath := "databaseDef"
	importName := ""
	key := "databaseDef"
	tmplFileRelativeDir := "databaseDef"
	return createPackageInfo(basicPath, relativePath, importName, tmplFileRelativeDir, key)
}

// GetOutsideConfPackageInfo 外部配置
func GetOutsideConfPackageInfo(basicPath string) confBean.PackageInfo {
	return confBean.NewPackageInfo("", GetIns().OutsideConf.FullPackageName(), GetIns().OutsideConf.PackageNameForRef(), "", "")
}

// relativePath		相对于 basicPath 的相对路径
// pkgPath			导入其他文件的包路径
// importName		导入其他文件的包名
// key				缓存key
func createPackageInfo(basicPath string, relativePath string, aliasPackageName string, tmplFileRelativeDir string, key string) confBean.PackageInfo {
	packageInfo, hasKey := packageInfoHolder[key]
	if !hasKey {
		packageName, _ := lo.Last(strings.Split(relativePath, "/"))
		packageAbsDirPath := filepath.Join(basicPath, relativePath)
		err := utils.MkDirsIfNotExist(packageAbsDirPath)
		if err != nil {
			panic("创建 " + packageName + " package 文件夹失败;" + err.Error())
		}

		clear, err := func() (func(), error) {
			filePath := filepath.Join(basicPath, relativePath, "placeholder_for_db_generator_.go")
			content := strings.NewReplacer(
				"{{PACKAGE_NAME}}", packageName,
				"{{AUTO_GENERATED_ANNOTATION}}", utils.AutoGeneratedFileAnnotation(),
			).Replace(`
{{AUTO_GENERATED_ANNOTATION}}

package {{PACKAGE_NAME}}

// client 需要获取 ModuleName，需要在 BasicPackage 文件夹中有代码文件，本文件做占位
`)
			content = strings.TrimPrefix(content, "\n")
			err = os.WriteFile(filePath, []byte(content), 0755)
			if err != nil {
				panic("写入 " + packageName + " package 占位文件失败;" + err.Error())
			}
			return func() {
				os.Remove(filePath)
			}, nil
		}()
		defer clear()

		fullPackageName, err := utils.FullPkgName(packageAbsDirPath)
		if err != nil {
			panic("获取 " + packageName + " package 全包名失败;" + err.Error())
		}
		packageInfo = confBean.NewPackageInfo(basicPath, fullPackageName, aliasPackageName, relativePath, tmplFileRelativeDir)
		packageInfoHolder[key] = packageInfo
	}
	return packageInfo
}
