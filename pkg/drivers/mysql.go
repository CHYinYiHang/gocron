package drivers

import (
	_ "database/sql"
	"github.com/CHYinYiHang/gocron/pkg/config"
	"github.com/CHYinYiHang/gocron/pkg/logging"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//show processlist 连接数查询
//select COUNT('Host') FROM information_schema.processlist

var MysqlDbPool *sqlx.DB
var MysqlDbError error

func LoadMysql() {

	MysqlDbPool, MysqlDbError = sqlx.Open("mysql", config.Config.MySQL.DSN())
	if MysqlDbError != nil {
		logging.Error("============数据库打开失败============")
		logging.Error(config.Config.MySQL.DSN())
		panic(MysqlDbError)
		return
	}

	MysqlDbError = MysqlDbPool.Ping()
	if MysqlDbError != nil {
		logging.Error("============数据库连接失败============")
		logging.Error(config.Config.MySQL.DSN())
		panic(MysqlDbError)
		return
	}

	MysqlDbPool.SetMaxIdleConns(config.Config.MySQL.DBMaxIdleContent)
	MysqlDbPool.SetMaxOpenConns(config.Config.MySQL.DBMaxOpenContent)
}
