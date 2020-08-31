package main

import (
	"algor/leetcode"
	"fmt"
)

func main() {
	debug()
}

func debug() {
	//linklist.TestLink()
	//linklist.TestDoubleLink()
	//stack.TestStack()
	//queue.TestQueue()
	//tree.TestBinary()
	//graph.TestGraph()
	//search.TestSearch()
	//search.TestBlock()
	//seqsort.TestSorts()
	leetcode.TestAlgor()
	test()
}

func test() {
	items := []int{1, 2, 3, 4, 5}
	doubleSlice1 := make([][]int, 1)
	doubleSlice2 := make([][]int, 1)
	singleSlice := items[1:3]
	doubleSlice1[0] = items[2:3]
	doubleSlice2[0] = items[1:3]
	items[2] = 999
	fmt.Println(items)
	fmt.Println(singleSlice)
	fmt.Println(doubleSlice1)
	fmt.Printf("singleSlice:%p,%p\n", &singleSlice, singleSlice)
	fmt.Printf("doubleSlice1:%p,%p\n", &doubleSlice1[0], doubleSlice1[0])
	fmt.Printf("doubleSlice2:%p,%p\n", &doubleSlice2[0], doubleSlice2[0])
}
