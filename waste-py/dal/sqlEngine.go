package dal

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"waste-py/backend/backend_model"
	"waste-py/backend/backend_model/basic_model"
	"waste-py/web/model"
)

var DB *xorm.Engine

func SqlEngine(conn string, driver string, show bool) {
	db, err := xorm.NewEngine(driver, conn)
	if err != nil {
		panic(fmt.Sprintf("连接数据库异常:%v\n", err))
	}
	db.ShowSQL(show)

	err = db.Sync2(
		new(model.Consumers),
		new(backend_model.Good),
		new(basic_model.Hr2020),
		new(model.Address),
		new(model.Order),
		new(model.Cities),
		new(model.Area),
		new(model.Position),
	)
	if err != nil {
		panic(fmt.Sprintf("创建数据表出现异常:%v\n", err))
	}
	DB = db
	return
}
