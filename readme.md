使用
```
go install github.com/ghoulhyk/dao-gen@latest
```
安装命令。

调用之后，能将如数据表定义转换成简单的dao

- 数据表定义：
```go
package _db1

import (
	"time"
)

// User
// @database db_1_test
// @tableName user
// @comment 用户信息表
type User struct {
	Id         *uint32    `json:"id" colAttr:"'id' pk autoincr"`
	Name       *string    `json:"name" colAttr:"'name'"` // 姓名
	Gender     *uint8     `json:"gender" colAttr:"'gender'"`
	CreateTime *time.Time `json:"createTime" colAttr:"'create_time'"`
}
```

- 生成后的代码的调用示例
```
dao.NewClient().Inserter().User().SetName("用户001").SetGender(1).Insert()

dao.NewClient().Selector().User().
    Where(func(cond *whereCond.User) {
        cond.And().Name().Equ("用户001")
    }).Single()

dao.NewClient().Updater().User().
    Where(func(cond *whereCond.User) {
        cond.And().Name().Equ("用户001")
    }).
    SetGender(1).
    Update()

dao.NewClient().Deleter().User().
    Where(func(cond *whereCond.User) {
        cond.And().Name().Equ("用户001")
    }).
    Delete()
```

更多示例详见 https://github.com/ghoulhyk/dao-gen-demo