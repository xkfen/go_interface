package _interface

import (
	"testing"
	"fmt"
)

type Adder interface {
	Add(x, y int) int
}
type AdderFunc func(x, y int) int

func (a AdderFunc) Add(x, y int) int {
	return a(x, y)
}

func Do(adder Adder) int {
	return adder.Add(1, 2)
}
func TestAddFunc(t *testing.T) {
	a := AdderFunc(func(x, y int) int {
		return x + y
	})
	fmt.Println(Do(a))
}
