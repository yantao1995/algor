package leetcode

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	ser := Constructor()
	deser := Constructor()

	str := "1,2,3,null,null,4,5,6,7,null,null,null,null,null,null"
	root := deser.deserialize(str)
	data := ser.serialize(root)
	fmt.Println("__1___", data)
	ans := deser.deserialize(data)
	data2 := ser.serialize(ans)
	fmt.Println("__2___", data2)
	fmt.Println()
}
