package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

type DBConfig struct {
	DbType   string
	UserName string
	Password string
	Host     string
	Port     int
	DbName   string
}

func InitDB(config *DBConfig) {
	connectStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.UserName,
		config.Password,
		config.Host,
		config.Port,
		config.DbName,
	)
	fmt.Println(connectStr)
	db, err := gorm.Open(config.DbType, connectStr)
	if err != nil {
		panic("数据库连接失败:")
	}

	db.SingularTable(true)
	//db.LogMode(true)
	//db.SetLogger(&log.OrmLogger{})
	DB = db
}
