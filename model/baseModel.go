package model

import (
	"fmt"
	initDBGorm "ginGrpcBase2/database/gorm"
	"gorm.io/gorm"
	"strings"
	"time"
)
type BaseModel struct {
	ID        uint           `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

//插入更新
func InsertOrUpdate(table string,data []map[string]interface{},fields string)(err error){
	var fieldsArr []string
	fieldsArr=strings.Split(fields,",")
	//
	var insertVal string
	var updateVal string
	for _,val:=range fieldsArr{
		for _,value:=range data{
			getVal,ok:=value[val]
			fmt.Println("getVal",getVal)
			fmt.Println("getVal.string",getVal.(string))
			if ok{
				insertVal+=getVal.(string)+","
				updateVal+=val+"="+getVal.(string)+","
			}
		}
	}
	updateVal=updateVal[0:len(fieldsArr)-1]
	insertVal=insertVal[0:len(fieldsArr)-1]
	sql := "insert into  " + table + "(" + fields + ") values(" + insertVal + ") on duplicate key update " + updateVal
	fmt.Println("insertUpdate: ", sql)
	err=initDBGorm.Db.Exec(sql).Error
	if err!=nil{
		return err
	}
	return nil
}