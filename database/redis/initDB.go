package redis

import "github.com/redis/go-redis/v9"



// 定义一个全局变量
var Db *redis.Client

func InitDB() {
	Db= redis.NewClient(&redis.Options{
		Addr:	  "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:		  0,  // 默认DB 0
	})

}
func init() {
	InitDB()
}