package leetcode

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	fmt.Println()
}

// str 用例  转 一维 string 数组
func str2SliceString(str string) []string {
	result := []string{}
	json.Unmarshal([]byte(str), &result)
	return result
}

// str 用例  转 一维 string 数组
func str2SliceString2(str string) [][]string {
	result := [][]string{}
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
	result := [][]int{}
	json.Unmarshal([]byte(str), &result)
	return result
}

// str 用例  转 单链表
func str2ListNode(str string) *ListNode {
	nodes := str2SliceInt(str)
	var listNode, h *ListNode
	for i, node := range nodes {
		nd := &ListNode{
			Val:  node,
			Next: nil,
		}
		if i == 0 {
			listNode = nd
			h = listNode
		}
		h.Next = nd
		h = h.Next
	}
	return listNode
}

// 打印单链表
func printListNode(list *ListNode) {
	fmt.Printf("%s", "[")
	for h := list; h != nil; h = h.Next {
		fmt.Printf("%d", h.Val)
		if h.Next != nil {
			fmt.Printf(",")
		}
	}
	fmt.Println("]")
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
