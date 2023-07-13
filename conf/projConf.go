package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/ghoulhyk/dao-gen/conf/confBean"
)

var (
	_conf *ProjectConfig
)

func GetIns() *ProjectConfig {
	return _conf
}

// Init init config.
func Init(confPath string) (*ProjectConfig, error) {
	_conf = &ProjectConfig{}
	_, err := toml.DecodeFile(confPath, &_conf)
	return _conf, err
}

type ProjectConfig struct {
	Path2basic  confBean.Path2basic
	OutsideConf confBean.PackageInfo
	OrmInfo     confBean.OrmInfo
	Database    confBean.DatabaseConf
}
