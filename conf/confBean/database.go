package confBean

import (
	"daogen/utils"
	"github.com/samber/lo"
)

type DatabaseConf struct {
	DatabaseInfos []DatabaseItemConf
}

type DatabaseItemConf struct {
	Name            string   // 生成的 databaseDef 的字段名
	DefDatabaseList []string // 实体类中的 @database 库名
	OutsideConfKey  string   // 互斥; 外部配置文件的key
	ActualName      string   // 互斥; 实际的库名
}

func (receiver DatabaseItemConf) ShouldUseOutsideConf() bool {
	return receiver.OutsideConfKey != ""
}

func (receiver DatabaseItemConf) FieldName() string {
	return utils.Capitalize(receiver.Name)
}

func (receiver DatabaseConf) FindDatabaseConfByName(schemaName string) *DatabaseItemConf {
	for _, schemaInfo := range receiver.DatabaseInfos {
		if lo.Contains(schemaInfo.DefDatabaseList, schemaName) {
			return &schemaInfo
		}
	}
	return nil
}
