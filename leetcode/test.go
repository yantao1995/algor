package leetcode

import (
	"fmt"
	"reflect"
)

//测试
func TestAlgor() {
	a := '2' - 48
	b := '4' - 48
	c := a * b
	fmt.Println(c%10+48, reflect.TypeOf(c))
	fmt.Println(multiply("12", "13"))
}
