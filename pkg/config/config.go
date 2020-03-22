package config

import "fmt"

var Config = &serverConfig{}

type serverConfig struct {
	Server server
	MySQL  mySQL
}

type server struct {
	Port        string
	LogFileName string
	LogFilePath string
	ServerId    int64
}

// MySQL mysql配置参数
type mySQL struct {
	Host             string
	Port             int
	User             string
	Password         string
	DBName           string
	Parameters       string
	DBMaxIdleContent int
	DBMaxOpenContent int
}

// MySQL 数据库连接串
func (a mySQL) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		a.User, a.Password, a.Host, a.Port, a.DBName, a.Parameters)
}
