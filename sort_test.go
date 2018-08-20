package main

import (
	"testing"
	"sort"
	"fmt"
)

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
// 为了读序列进行排序，需要定义一个实现了这三个方法的类型,然后对这个类型的实例应用sort.Sort函数

type StringSlice []string
func (p StringSlice) Len() int {return len(p)}
func (p StringSlice) Less(i, j int) bool {return p[i] < p[j]}
func (p StringSlice) Swap(i, j int) {p[i], p[j] = p[j], p[i]}

func TestStringSliceSort(t *testing.T)  {
	names := []string{"test3", "test1", "test2"}
	fmt.Println("排序之前")
	for _, data := range names {
		fmt.Println(data)
	}
	sort.Sort(StringSlice(names))
	fmt.Println("排序之后")
	for _, data := range names {
		fmt.Println(data)
	}
}

