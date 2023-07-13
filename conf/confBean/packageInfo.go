package confBean

import (
	"encoding/json"
	"github.com/dave/dst"
	"github.com/samber/lo"
	"path/filepath"
	"strconv"
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

func (receiver PackageInfo) ImportStatement() string {
	if receiver.ShouldUseAliasPackageName() {
		return receiver.aliasPackageName + " \"" + receiver.fullPackageName + "\""
	} else {
		return "\"" + receiver.fullPackageName + "\""
	}
}

func (receiver PackageInfo) ImportDeclSpec() dst.Spec {
	decl := &dst.ImportSpec{Path: &dst.BasicLit{Value: strconv.Quote(receiver.fullPackageName)}}
	if receiver.ShouldUseAliasPackageName() {
		decl.Name = dst.NewIdent(receiver.aliasPackageName)
	}
	return decl
}

// PackageNameForRef 其他文件中调用方法时得前缀 {{PackageNameForRef}}.FuncName()
func (receiver PackageInfo) PackageNameForRef(structOrFuncName ...string) string {
	name := lo.Ternary(receiver.aliasPackageName == "", receiver.packageName, receiver.aliasPackageName)
	if len(structOrFuncName) > 0 {
		return name + "." + structOrFuncName[0]
	}
	return name
}

// StarPackageNameForRef 其他文件中调用方法时得前缀 {{PackageNameForRef}}.FuncName()
func (receiver PackageInfo) StarPackageNameForRef(structOrFuncName ...string) string {
	return "*" + receiver.PackageNameForRef(structOrFuncName...)
}

func (receiver PackageInfo) ShouldUseAliasPackageName() bool {
	return receiver.aliasPackageName != "" && receiver.aliasPackageName != receiver.packageName
}

func (receiver PackageInfo) DirAbsPath() string {
	return filepath.Join(receiver.basicPath, receiver.relativeDirPath)
}

func (receiver PackageInfo) DirRelativePath() string {
	return receiver.relativeDirPath
}

func (receiver PackageInfo) DirTmplFileRelative() string {
	return receiver.tmplFileRelativeDir
}

func (receiver PackageInfo) FullPackageName() string {
	return receiver.fullPackageName
}

func (receiver *PackageInfo) UnmarshalJSON(data []byte) (err error) {
	tmp := &struct {
		FullPackageName  string `json:"fullPackageName"`
		AliasPackageName string `json:"aliasPackageName"`
	}{}
	err = json.Unmarshal(data, tmp)
	if err != nil {
		return
	}
	*receiver = NewPackageInfo("", tmp.FullPackageName, tmp.AliasPackageName, "", "")
	return
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
