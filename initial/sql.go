package initial

import (
	"fmt"
	_ "github.com/MobileCPX/PreBobsTube/models/web"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	//_ "github.com/go-sql-driver/mysql"
)

func InitSql() {
	user := beego.AppConfig.String("psqluser")
	passwd := beego.AppConfig.String("psqlpass")
	host := beego.AppConfig.String("psqlurls")
	port, err := beego.AppConfig.Int("psqlport")
	dbname := beego.AppConfig.String("psqldb")

	//fmt.Printf(
	//	"user=%s password=%s dbname=%s host=%s port=%d",
	//	user, passwd, dbname, host, port)

	if nil != err {
		port = 5432
	}
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
	orm.Debug = true
	orm.DefaultRowsLimit = -1
	_ = orm.RegisterDriver("postgres", orm.DRPostgres) // 注册驱动
	_ = orm.RegisterDataBase("default", "postgres",
		fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
			user, passwd, dbname, host, port))
	_ = orm.RunSyncdb("default", false, true)

	////跟换成MySQL：
	//orm.RegisterDriver("mysql", orm.DRMySQL)
	//orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/prebobstube?charset=utf8")
	//orm.SetMaxIdleConns("default",1000)
	//orm.SetMaxOpenConns("default",2000)
	////要指定默认的数据库
	//_ = orm.RunSyncdb("default", false, true)

}
