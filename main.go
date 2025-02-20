package main

import (
	"algor/utils"
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
	ts := int64(1740758400)
	fmt.Println(utils.GetMonthStart(time.Local, ts))
	fmt.Println(utils.GetMonthEnd(time.Local, ts))
}
