package validator

import (
	"fmt"
	"github.com/go-playground/locales/zh_Hans_CN"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"my_gvb/utils/errmsg"
	"reflect"
)

// Validate 验证数据
func Validate(data any) (string, int) {
	validate := validator.New()                 // 实例化验证器
	uni := unTrans.New(zh_Hans_CN.New())        // 实例化翻译器
	trans, _ := uni.GetTranslator("zh_Hans_CN") // 获取翻译器中的内容

	err := zhTrans.RegisterDefaultTranslations(validate, trans) // 将翻译器注册到验证器中
	if err != nil {
		fmt.Println("err:", err)
	}
	validate.RegisterTagNameFunc(func(field reflect.StructField) string { // 注册自定义字段名称
		label := field.Tag.Get("label")
		return label
	})

	err = validate.Struct(data) // 验证数据
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans), errmsg.ERROR
		}
	}
	return "", errmsg.SUCCSE
}
