/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-09-15 14:21:38
 * @LastEditTime: 2022-11-17 15:32:04
 * @Description: Do not edit
 */
package conver

import (
	"encoding/json"
	"strconv"
)

/**
 * @description: conver interface to string (获取变量的字符串值)
 * 浮点型 3.0将会转换成字符串3, "3"
 * 非数值或字符类型的变量将会被转换成JSON格式字符串
 * @param {interface{}} value
 * @return {*}
 */
func ToString(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch val := value.(type) {
	case float64:
		key = strconv.FormatFloat(val, 'f', -1, 64)
	case float32:
		key = strconv.FormatFloat(float64(val), 'f', -1, 64)
	case int:
		key = strconv.Itoa(val)
	case uint:
		key = strconv.Itoa(int(val))
	case int8:
		key = strconv.Itoa(int(val))
	case uint8:
		key = strconv.Itoa(int(val))
	case int16:
		key = strconv.Itoa(int(val))
	case uint16:
		key = strconv.Itoa(int(val))
	case int32:
		key = strconv.Itoa(int(val))
	case uint32:
		key = strconv.Itoa(int(val))
	case int64:
		key = strconv.FormatInt(val, 10)
	case uint64:
		key = strconv.FormatUint(val, 10)
	case string:
		key = val
	case []byte:
		key = string(val)
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}
