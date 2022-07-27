package leetcode

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	sk := Constructor()
	sk.Add(9)
	sk.Add(4)
	sk.Add(5)
	sk.Add(6)
	sk.Add(9)
	sk.Print()
	fmt.Println(sk.Erase(2))
	sk.Print()
	fmt.Println(sk.Erase(1))
	sk.Print()
	sk.Add(2)
	sk.Print()
	fmt.Println(sk.Search(7))
	fmt.Println(sk.Search(4))
	sk.Add(5)
	sk.Print()
	fmt.Println(sk.Erase(6))
	sk.Print()
	fmt.Println(sk.Search(5))
	//fmt.Println(sk.Search(2))
	//fmt.Println(sk.Search(1))
	fmt.Println()
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
