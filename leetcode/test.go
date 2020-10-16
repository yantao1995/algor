package leetcode

import (
	"fmt"
)

//测试
func TestAlgor() {
	words := []string{"word", "world", "row"}
	order := "worldabcefghijkmnpqstuvxyz"

	fmt.Println(isAlienSorted(words, order))
}
