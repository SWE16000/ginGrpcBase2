package utils

import "time"

//时间戳转时间
func TimestampToTime(timestamp int64)string   {
	//timeTmeplate:="2006年1月2日 15:04:05"
	timeTmeplate:="2006-01-02 15:04:05"
	//date, _ := time.Parse(timeTmeplate , string(timestamp))
	date := time.Unix(timestamp, 0).Format(timeTmeplate)
	return date
}

//时间转时间戳
func TimesToTimestamp(){
	timeTemplate1 := "2006-01-02 15:04:05" //常规类型
	t1 := "2019-01-08 13:50:30" //外部传入的时间字符串
	stamp, _ := time.ParseInLocation(timeTemplate1, t1, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	stamp.Unix()  //输出：1546926630
	time.ParseInLocation(timeTemplate1,t1,time.Local)
}