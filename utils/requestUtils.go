package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)


func PostRequest(url string,params map[string]interface{})(res interface{},err error)  {
	//序列化
	postBody, _ :=json.Marshal(params)
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(url, "application/json", responseBody)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	data := string(body)
	var dataMap map[string]interface{}
	//反序列化
	err=json.Unmarshal([]byte(data),&dataMap)
	if err!=nil{
		return nil,err
	}
	return dataMap,nil
}

func Get(url string,params map[string]interface{},header []map[string]interface{})(res interface{},err error)  {
	client := &http.Client{}
	//拼接参数
	str:="?"
	sum:=0
	for index,value:=range params{
		if sum==0{
			str+=index+"="+value.(string)
		}else{
			str+="&"+index+"="+value.(string)
		}
		sum++
	}
	url=url+str
	resp, err := http.NewRequest("GET",url, nil)
	if err != nil {
		return nil,err
	}
	//添加请求头
	if len(header)!=0{
		for _,value:=range header{
			for index2,value2:=range value{
				resp.Header.Add(index2,value2.(string))
			}
		}
	}
	resp.Header.Add("Content-Type", "application/json")
	response, _ := client.Do(resp)
	defer response.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil,err
	}
	data := string(body)
	var dataMap map[string]interface{}
	//反序列化
	err=json.Unmarshal([]byte(data),&dataMap)
	if err!=nil{
		return nil,err
	}
	return dataMap,nil
}
func Post(url string,params map[string]interface{},header []map[string]interface{})(res interface{},err error)  {
	client := &http.Client{}
	//序列化
	postBody, _ :=json.Marshal(params)
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.NewRequest("POST",url, responseBody)
	if err != nil {
		return nil,err
	}
	//添加请求头
	if len(header)!=0{
		for _,value:=range header{
			for index2,value2:=range value{
				resp.Header.Add(index2,value2.(string))
			}
		}
	}
	resp.Header.Add("Content-Type", "application/json")
	response, _ := client.Do(resp)
	defer response.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil,err
	}
	data := string(body)
	fmt.Println(data)
	var dataMap map[string]interface{}
	//反序列化
	err=json.Unmarshal([]byte(data),&dataMap)
	if err!=nil{
		return nil,err
	}
	return dataMap,nil
}