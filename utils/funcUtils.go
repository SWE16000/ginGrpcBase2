package utils

import "fmt"

//权限树构造
func MakeTree(data []map[string]interface{},pid int64)([]map[string]interface{}){
	var tree []map[string]interface{}
	for _,value:=range data{
		if value["parent_id"].(int64)==pid{
			temp:=MakeTree(data,value["id"].(int64))
			if len(temp)>0{
				value["children"]=temp
			}else{
				value["children"]=""
			}
			tree=append(tree,value)
		}
	}
    return tree
}
func MakeTree2(data []map[string]interface{},pid int64)([]map[string]interface{}){
	var tree []map[string]interface{}
	for _,value:=range data{
		if int64(value["ParentId"].(float64))==pid{
			temp:=MakeTree2(data,int64(value["Id"].(float64)))
			if len(temp)>0{
				value["children"]=temp
			}else{
				value["children"]=""
			}
			tree=append(tree,value)
		}
	}
	return tree
}
/**
    1     2
     -3    -6 -8  -9
      -4       -7
       -5
 */
//向上查找所有父级权限id
func UpFindPid(data [] map[string]interface{},pid int64)(res []int64)  {
	var arr []int64
	for _,value:=range data{
		 intPid:=int64(value["Id"].(float64))
		if  intPid== pid {
			if int64(value["ParentId"].(float64))!=0{
				temp:=UpFindPid(data,int64(value["ParentId"].(float64)))
				fmt.Println("temp",temp)
				for _,value2:=range temp{
					arr=append(arr,value2)
				}
				arr=append(arr,int64(value["Id"].(float64)))
			}else{
				arr=append(arr,int64(value["Id"].(float64)))
			}

	    }
	}
	return arr
}