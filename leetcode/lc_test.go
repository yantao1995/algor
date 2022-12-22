package leetcode

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestFunc(t *testing.T) {
	fmt.Println(maxScore([]int{1231, 521, 2539, 1213, 2699, 1667, 1361, 1231, 521, 2539, 1213, 2699, 1667, 1361}))
}

// str 用例  转 一维 string 数组
func str2SliceString(str string) []string {
	result := []string{}
	json.Unmarshal([]byte(str), &result)
	return result
}

// str 用例  转 一维int数组
func str2SliceInt(str string) []int {
	result := []int{}
	json.Unmarshal([]byte(str), &result)
	return result
}

// str 用例  转 二维int数组
func str2SliceInt2(str string) [][]int {
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
