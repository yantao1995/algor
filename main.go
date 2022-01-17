package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	test()
}

func test() {
	fmt.Println(getRandQuestion())
}

func getRandQuestion() int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(2088)
}
