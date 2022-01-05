package leetcode

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	ms := Constructor()
	ms.Insert("apple", 3)
	fmt.Println(ms.Sum("ap"))
	ms.Insert("app", 2)
	fmt.Println(ms.Sum("ap"))
}
