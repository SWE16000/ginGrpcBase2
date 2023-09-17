package utils

import (
	"errors"
	"fmt"
	"github.com/go-playground/locales/zh_Hans_CN"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"regexp"
)

func Validate(data interface{},r interface{}) (string) {
	////验证对象 并将英文报错转换成中文报错（message）
	validate := validator.New()
	uni := unTrans.New(zh_Hans_CN.New())
	trans,_ := uni.GetTranslator("zh_Hans_CN")
	err := zhTrans.RegisterDefaultTranslations(validate,trans)
	if err != nil {
		fmt.Println("err:",err)
		return ""
	}
	//将验证法字段名 映射为中文名
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})
	err = validate.Struct(data)
	if err != nil {
		//错误可能有多个 遍历 返回一个
		for _,v := range err.(validator.ValidationErrors) {
			s := reflect.TypeOf(r)
			errTag := v.Tag() + "_msg"
			// 获取对应binding得错误消息
			filed, _ := s.FieldByName(v.Field())
			errTagText := filed.Tag.Get(errTag)

			//return v.Translate(trans)
			return errTagText
		}
	}
	return ""
}
func CheckPhone(phone string)error  {
	phoneRegular:="^1[3456789][0-9]{9}$"
	//调用正则规则
	reg:=regexp.MustCompile(phoneRegular)
	//是否匹配
	if !reg.MatchString(phone){
		return errors.New("用户手机号格式错误")
	}
	return nil
}
