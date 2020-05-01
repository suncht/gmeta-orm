package main

import (
	"fmt"
	"gmeta-orm/db"
	"gmeta-orm/def"
)

func main() {
	dbConfig := &db.DBConfig{
		DbType:   "mysql",
		DbName:   "autocoder",
		UserName: "root",
		Password: "123456",
		Host:     "127.0.0.1",
		Port:     3306,
	}
	db.InitDB(dbConfig)

	var tableDefService = &db.DefaultTableDefService{}

	loader := def.NewLoader()
	loader.SetTableDefService(tableDefService)
	def, _ := loader.Get("bz_users")
	fmt.Println(def)
	def, _ = loader.Get("bz_users")
	fmt.Println(def)
}
