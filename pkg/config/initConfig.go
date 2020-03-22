package config

import (
	"github.com/spf13/viper"
)

func LoadConf() {
	v := viper.New()

	v.SetConfigFile("./conf/config.yml")
	v.SetConfigType("yml")
	if err1 := v.ReadInConfig(); err1 != nil {
		panic(err1)
		return
	}
	//service
	Config.Server.Port = ":" + v.GetString("service.port")
	Config.Server.ServerId = v.GetInt64("service.server_id")
	Config.Server.LogFileName = v.GetString("service.log.name")
	Config.Server.LogFilePath = v.GetString("service.log.path")
	//db
	Config.MySQL.User = v.GetString("mysql.user")
	Config.MySQL.Host = v.GetString("mysql.host")
	Config.MySQL.Port = v.GetInt("mysql.port")
	Config.MySQL.User = v.GetString("mysql.user")
	Config.MySQL.Password = v.GetString("mysql.password")
	Config.MySQL.DBName = v.GetString("mysql.db_name")
	Config.MySQL.Parameters = v.GetString("mysql.parameters")
	Config.MySQL.DBMaxIdleContent = v.GetInt("mysql.DB_MAX_IDLE_Contents")
	Config.MySQL.DBMaxOpenContent = v.GetInt("mysql.DB_MAX_OPEN_Contents")
}
