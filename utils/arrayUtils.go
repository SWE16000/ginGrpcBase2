package utils

//数组分割
func ArraySlice(data[]map[string]interface{},start int,end int)[]map[string]interface{}{
	var res []map[string]interface{}
	for index,value:=range data{
		if index>=start&&index<end{
			res=append(res,value)
		}
	}
	return res
}