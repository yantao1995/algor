package main

import (
	"algor/utils/chain"
	"fmt"
	"math/rand"
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
	chain.MiningCrash()
}
