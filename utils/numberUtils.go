package utils

import (
	"math/rand"
	"strings"
)

func RoundNumber(lens int)string  {
	var Char=[]string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z","0","1","2","3","4","5","6","7","8","9"}
	str := strings.Builder{}
	length := len(Char)
	for i:=0;i<lens;i++{
		l := Char[rand.Intn(length)]
		str.WriteString(l)
	}
	return str.String()
}

func SelMake(data map[string]interface{})map[string]interface{}  {
	_,ok:=data["page"]
	if !ok{
		data["page"]=1
	}
	_,ok2:=data["pageSize"]
	if !ok2{
		data["pageSize"]=15
	}
	return data
}