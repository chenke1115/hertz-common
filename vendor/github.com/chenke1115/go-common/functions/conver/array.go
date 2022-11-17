/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-17 15:23:25
 * @LastEditTime: 2022-11-17 15:36:38
 * @Description: Do not edit
 */
package conver

import (
	"fmt"
	"reflect"
	"strconv"
)

/**
 * @description: Conver []interface{} To array
 * @param {interface{}} data
 * @return {*}
 */
func ToArray(data interface{}) (strArr []string) {
	interArr := data.([]interface{})
	for _, v := range interArr {
		strArr = append(strArr, v.(string))
	}

	return
}

/**
 * @description: Conver struct To array
 * @param {interface{}} data
 * @return {*}
 */
func StructToArray(data interface{}) (strArr []string) {
	typeInfo := reflect.TypeOf(data)
	valInfo := reflect.ValueOf(data)
	for i := 0; i < typeInfo.NumField(); i++ {
		val := valInfo.Field(i).Interface()
		strArr = append(strArr, val.(string))
	}

	return
}

/**
 * @description: Conver ...interface{} to []string
 * @param {...interface{}} params
 * @return {*}
 */
func ArgsToArray(args ...interface{}) (strArr []string) {
	for _, param := range args {
		switch v := param.(type) {
		case int:
			strV := strconv.FormatInt(int64(v), 10)
			strArr = append(strArr, strV)
		case float64:
			strV := strconv.FormatFloat(v, 'f', -1, 64)
			strArr = append(strArr, strV)
		case string:
			strArr = append(strArr, v)
		case []string:
			strArr = v
		default:
			panic(fmt.Errorf("params type not supported"))
		}
	}

	return
}

/**
 * @description: Conver []string to *[]string
 * @param {[]string} strArr
 * @return {*}
 */
func ToPointArray(strArr []string) (stringArr *[]string) {
	stringArr = &[]string{}
	*stringArr = append(*stringArr, strArr...)
	return
}
