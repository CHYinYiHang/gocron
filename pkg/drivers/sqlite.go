package drivers

//
//import (
//	_ "database/sql"
//	"fmt"
//	"github.com/jmoiron/sqlx"
//	_ "github.com/mattn/go-sqlite3"
//	"hopitalServer/pkg/config"
//)
//
//var SqlLiteDbPool *sqlx.DB
//var SqlLiteDbError error
//
//func init() {
//
//	SqlLiteDbPool, SqlLiteDbError = sqlx.Open("sqlite3", config.Config.Sqlite.DSN())
//	if SqlLiteDbError != nil {
//		fmt.Println("数据库打开失败！")
//		fmt.Println(SqlLiteDbError)
//		panic(SqlLiteDbError)
//	}
//
//	SqlLiteDbError = SqlLiteDbPool.Ping()
//	if SqlLiteDbError != nil {
//		fmt.Println("数据库连接失败！")
//		fmt.Println(SqlLiteDbError)
//		panic(SqlLiteDbError)
//	}
//
//	SqlLiteDbPool.SetMaxIdleConns(config.Config.MySQL.DBMaxIdleContent)
//	SqlLiteDbPool.SetMaxOpenConns(config.Config.MySQL.DBMaxOpenContent)
//	fmt.Println(SqlLiteDbPool, "linkDB is ok")
//}
