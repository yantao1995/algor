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
	str := "asdf"
	fmt.Printf("%p\n", &str)
	str = "asdf" + "aasdasdasdasdasdasdasd"
	fmt.Printf("%p", &str)
}
