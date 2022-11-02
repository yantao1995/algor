package leetcode

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestFunc(t *testing.T) {
	fmt.Println()
}

// str 用例转 二维int数组
func str2Slice2(str string) [][]int {
	slice := strings.Split(str[2:len(str)-2], "],[")
	result := [][]int{}
	for _, v := range slice {
		sli2 := strings.Split(v, ",")
		temp := []int{}
		for _, v2 := range sli2 {
			v2Int, _ := strconv.Atoi(v2)
			temp = append(temp, v2Int)
		}
		result = append(result, temp)
	}
	return result
}

//

// func TestAlgor() {

// 	fmt.Println()
// 	fmt.Println("______________")
// 	//
// 	lfu := Constructor(2)
// 	lfu.Put(1, 1)
// 	lfu.Put(2, 2)
// 	fmt.Println(lfu.Get(1))
// 	lfu.iterator()
// 	fmt.Println("lfu.KeyValueMap", lfu.KeyValueMap)
// 	fmt.Println("lfu.CountLinkMap", lfu.CountLinkMap)
// 	fmt.Println("______________")
// 	lfu.Put(3, 3)
// 	fmt.Println("lfu.KeyCountMap", lfu.KeyCountMap)
// 	fmt.Println("lfu.KeyValueMap", lfu.KeyValueMap)
// 	fmt.Println("lfu.KeyCountMap", lfu.KeyCountMap)
// 	fmt.Println("lfu.CountLinkMap", lfu.CountLinkMap)
// 	fmt.Println("______________")
// 	fmt.Println(lfu.Get(2))
// 	fmt.Println(lfu.Get(3))
// 	lfu.Put(4, 4)
// 	fmt.Println(lfu.Get(1))
// 	fmt.Println(lfu.Get(3))
// 	fmt.Println(lfu.Get(4))
// }

// func (this *LFUCache) iterator() {
// 	fmt.Print("CountSortList   ")
// 	for h := this.CountSortList.Front(); h != nil; h = h.Next() {
// 		fmt.Print("\t", h.Value, "\t")
// 	}
// 	fmt.Print("\nCountLinkMap{1}   ")
// 	for h := this.CountLinkMap{1}.Front(); h != nil; h = h.Next() {
// 		fmt.Print("\t", h.Value, "\t")
// 	}

// }

/////////////////////////////////////////////////////
