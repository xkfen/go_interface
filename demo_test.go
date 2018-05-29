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
func getMaxKeyValue(hkjlMap map[string]string) (key, value int) {
	for k, v := range hkjlMap {
		// 将string类型的k转为int,这个绝对不会有err,因为key都是数字1-24
		intKey, _ := strconv.Atoi(k)
		// 将string类型的value转int，如果没有错误，就说明该value是字符串数字
		intValue, err := strconv.Atoi(v)
		if err == nil {
			// 如果value相等，取最大的key
			if value == intValue && key < intKey{
				value = intValue
				key = intKey
			}else if key < intKey && value <= intValue{
				value = intValue
				key = intKey
			}
		}
	}
	return
}

// 查找数组中最小的元素
func getMinValueInArray(array []int) int {
	min := array[0]
	for i := 1; i < len(array)-1; i++ {
		if min > array[i] {
			min = array[i]
		}
	}
	return min
}