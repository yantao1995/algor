package main

import (
	"fmt"
	"math/rand"
	"slices"
	"time"
)

func main() {
	test2()
}

func test() {
	fmt.Println(getRandQuestion())
}

func getRandQuestion() int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(2088)
}

func test2() {
	stack := []int{1}
	stack = slices.Insert(stack, 1, 111)
	fmt.Println(stack)
}
