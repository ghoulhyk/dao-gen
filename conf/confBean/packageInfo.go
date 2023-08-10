package confBean

import (
	"github.com/samber/lo"
	"strings"
)

type PackageInfo struct {

	// import j "encoding/json"

	basicPath           string
	fullPackageName     string // 完整包名 （上面的 encoding/json ）
	aliasPackageName    string // 包名别名 （上面的 j ）
	relativeDirPath     string // 实际文件系统的完整路径
	tmplFileRelativeDir string // 模板文件的相对目录

	packageName string // 包名 （上面的 json ）
}

// NewPackageInfo
// basicPath			: 生成文件夹的基础路径
// fullPackageName		: 完整包名
// aliasPackageName		: 包名别名
// relativeDirPath		: 生成文件夹相对于 basicPath 的路径
// tmplFileRelativeDir	: 模板文件的相对目录
func NewPackageInfo(basicPath string, fullPackageName string, aliasPackageName string, relativeDirPath string, tmplFileRelativeDir string) PackageInfo {
	packageName, _ := lo.Last(strings.Split(fullPackageName, "/"))
	return PackageInfo{
		basicPath:           basicPath,
		fullPackageName:     fullPackageName,
		aliasPackageName:    aliasPackageName,
		relativeDirPath:     relativeDirPath,
		tmplFileRelativeDir: tmplFileRelativeDir,

		packageName: packageName,
	}
}

// ImportStatement
// 引用包的语句
// 本方法在模板中有使用
func (receiver PackageInfo) ImportStatement() string {
	if receiver.shouldUseAliasPackageName() {
		return receiver.aliasPackageName + " \"" + receiver.fullPackageName + "\""
	} else {
		return "\"" + receiver.fullPackageName + "\""
	}
}

// RefName
// 其他文件中调用方法时得前缀
// 本方法在模板中有使用
func (receiver PackageInfo) RefName() string {
	name := lo.Ternary(receiver.aliasPackageName == "", receiver.packageName, receiver.aliasPackageName)
	return name
}

// 引用时是否使用别名
func (receiver PackageInfo) shouldUseAliasPackageName() bool {
	return receiver.aliasPackageName != "" && receiver.aliasPackageName != receiver.packageName
}

// DirRelativePath
// 生成的代码相对路径
func (receiver PackageInfo) DirRelativePath() string {
	return receiver.relativeDirPath
}

// DirTmplFileRelative
// 模板文件夹名
func (receiver PackageInfo) DirTmplFileRelative() string {
	return receiver.tmplFileRelativeDir
}

func (receiver PackageInfo) FullPackageName() string {
	return receiver.fullPackageName
}

func (receiver *PackageInfo) UnmarshalTOML(data interface{}) (err error) {
	getData := func(dataMap map[string]interface{}, key string) string {
		if dataMap[key] == nil {
			return ""
		}
		return dataMap[key].(string)
	}
	dataMap := data.(map[string]interface{})
	fullPackageName := getData(dataMap, "fullPackageName")
	aliasPackageName := getData(dataMap, "PackageName")
	*receiver = NewPackageInfo("", fullPackageName, aliasPackageName, "", "")
	return nil
}
