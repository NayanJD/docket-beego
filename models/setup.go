package models

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
)

var SqlConnString string
var Orm orm.Ormer

func init() {
	orm.RegisterModel(&User{})

	sqlConn, ok := beego.AppConfig.String("sqlconn")

	if ok != nil {
		panic("sqlconn not set in config")
	}

	SqlConnString = sqlConn

	orm.RegisterDataBase("default", "postgres", SqlConnString)

	Orm = orm.NewOrm()
}
