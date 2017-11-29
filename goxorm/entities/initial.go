package entities

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

//var mydb *sql.DB
var myengine *xorm.Engine

func init() {
	//https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
	engine, err := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	//db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
	//mydb = db
	myengine = engine
	engine.Sync2(new(UserInfo))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
