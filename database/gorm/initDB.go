package gorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var Db *gorm.DB

func InitDB() {
	//host := viper.GetString("mysql.host")
	//port := viper.GetString("mysql.port")
	//database := viper.GetString("mysql.database")
	//username := viper.GetString("mysql.username")
	//password := viper.GetString("mysql.password")
	//dsn := strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", database}, "")

	dsn := "root:SWE16081@tcp(127.0.0.1:3306)/gintest"
	//dsn := "gintest:zGCWeYrJ6tYnYnxM@tcp(127.0.0.1:3306)/gintest"
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println("err",err)
	// 迁移 schema
	//db.AutoMigrate(&Product{})

}
func init() {
	InitDB()
}