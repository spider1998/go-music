package app

import (
	_ "github.com/Go-SQL-Driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"xorm.io/core"
)

// LoadDB 创建DB
func LoadDB(dsn string) (x *xorm.Engine, err error) {
	x, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		log.Fatal("数据库连接失败:", err)
		return
	}
	x.SetMapper(core.GonicMapper{})
	x.ShowSQL(true)
	x.ShowExecTime(true)
	return
}
