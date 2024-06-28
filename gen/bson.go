package gen

import (
	"github.com/sirupsen/logrus"
	"reflect"
	"strings"
	"unicode"
)

// lowerCamelCase 将字符串转换为小写开头的驼峰形式
func lowerCamelCase(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

// getBsonTag 获取对象字段的 BSON 标签
func getBsonTag(obj interface{}, field interface{}) string {
	// 检查 obj 是否为 nil
	if obj == nil {
		logrus.Fatal("obj cannot be nil")
		return ""
	}

	// 检查 field 是否为 nil
	if field == nil {
		logrus.Fatal("field cannot be nil")
		return ""
	}
	if reflect.TypeOf(field).Kind() != reflect.Ptr {
		logrus.Fatal("field must be a pointer")
		return ""
	}
	objVal := reflect.ValueOf(obj)
	if objVal.Kind() != reflect.Ptr {
		logrus.Fatal("obj must be a pointer")
		return ""
	}

	objVal = objVal.Elem()
	objType := objVal.Type()
	tag1 := getTag(objVal, objType, field)
	if tag1 != "" {
		return tag1
	}
	// 如果字段不是当前结构体的直接字段，尝试递归查找嵌套结构体字段的标签
	for i := 0; i < objVal.NumField(); i++ {
		if objVal.Field(i).Type().Kind() == reflect.Struct {
			if tag := getTag(objVal.Field(i), objVal.Field(i).Type(), field); tag != "" {
				return tag
			}
		}
	}
	logrus.Fatal("field not found")
	return ""
}

func getTag(objVal reflect.Value, objType reflect.Type, field any) string {
	// 遍历结构体字段，查找与 fieldValue 相匹配的字段
	for i := 0; i < objVal.NumField(); i++ {
		if objVal.Field(i).Addr().Interface() == field {
			structField := objType.Field(i)
			// 获取BSON标签
			bsonTag := structField.Tag.Get("bson")
			if bsonTag == "" {
				// 如果没有BSON标签，使用字段的小写开头的驼峰形式名称
				return lowerCamelCase(structField.Name)
			}
			if bsonTag == "-" {
				// 处理 "-" 标签的情况
				return "-"
			}
			if strings.Contains(bsonTag, ",inline") {
				// 处理 ",inline" 标签的情况
				return ""
			}
			// 返回第一个逗号之前的部分作为标签
			return strings.Split(bsonTag, ",")[0]
		}
	}
	// 如果字段不是当前结构体的直接字段，尝试递归查找嵌套结构体字段的标签
	for i := 0; i < objVal.NumField(); i++ {
		if objVal.Field(i).Type().Kind() == reflect.Struct {
			if tag := getTag(objVal.Field(i), objVal.Field(i).Type(), field); tag != "" {
				return tag
			}
		}
	}
	return ""
}
