package validator

import (
	"backend/logger"
	"backend/utils"
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	v "github.com/go-playground/validator/v10"
)

func getJSONFieldName(t reflect.Type, structField string) string {
	field, ok := t.FieldByName(structField) // 获取结构体字段
	if !ok {
		return strings.ToLower(structField)
	}

	jsonTag := field.Tag.Get("json") // 获取字段 json 标签
	if jsonTag == "" || jsonTag == "-" {
		return strings.ToLower(structField)
	}

	return strings.Split(jsonTag, ",")[0] // 获取 json 标签中的字段名 "name,omitempty"
}

func GetErrorMap(err error, obj interface{}) map[string]string {
	if err == nil {
		return nil
	}
	validationErrors, ok := err.(v.ValidationErrors)
	if !ok {
		return map[string]string{
			"request": "参数错误",
		}
	}

	t := reflect.TypeOf(obj)     // 获取结构体类型
	if t.Kind() == reflect.Ptr { // 判断底层类型是否是指针类型， 如果是指针类型，则获取结构体类型
		t = t.Elem() // 根据地址找到结构体类型，相当于是reflect.TypeOf
	}

	errMap := make(map[string]string) // 创建一个空的错误映射

	for _, fieldErr := range validationErrors {
		field := getJSONFieldName(t, fieldErr.Field()) // 获取字段名

		switch fieldErr.Tag() { // 触发的验证标签
		case "required":
			errMap[field] = fmt.Sprintf("%s 为必填项", field)
		case "email":
			errMap[field] = fmt.Sprintf("%s 邮箱格式不正确", field)
		case "min":
			errMap[field] = fmt.Sprintf("%s 长度不能小于 %s", field, fieldErr.Param())
		default:
			errMap[field] = fmt.Sprintf("%s 不合法", field)
		}
	}
	return errMap
}

func BindJSON(c *gin.Context, obj interface{}) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		errMap := GetErrorMap(err, obj)
		utils.ValidationError(c, "参数错误", errMap)
		logger.Errorf("参数校验失败: %v", err)
		return false
	}
	return true
}
