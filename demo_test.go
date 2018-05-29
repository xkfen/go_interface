package main

import (
	"testing"
	"strings"
	"fmt"
	"go_interface/tool"
	"strconv"
)

// 测试字符串数组转字符串
func TestStrArray2Str(t *testing.T){
	array := []string{"katy", "key", "test"}
	tmp := strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
	fmt.Println(tmp)
}

// 测试过滤数组里面的重复元素
func filter(array []string) []string {
	result := make(map[string]int)

	for i, data := range array {
		result[data] = i
	}
	var tmpArr []string
	for k, _ := range result {
		tmpArr = append(tmpArr, k)
	}
	return tmpArr
}
// 测试字符串数组去重
// 测试去除数组中重复元素
func TestFilterDuplicateElem(t *testing.T){
	array := []string {"hello", "katy", "key", "hello world", "key", "katy"}
	r := filter(array)
	fmt.Println(tool.StringifyJson(r))
}

// 从map[string]string中找到最大的value对应的key值
func getMaxKeyValue(hkjlMap map[string]string)(key, value int){
	for k, v := range hkjlMap {
		intKey, _ := strconv.Atoi(k)
		intValue , err := strconv.Atoi(v)
		if err == nil {
			if key < intKey {
				key = intKey
				value = intValue
			}
			if value == intValue {
				if key < intKey {
					key = intKey
					value = intValue
				}
			}
		}
	}
	return
}