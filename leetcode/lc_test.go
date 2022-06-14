package leetcode

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
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
