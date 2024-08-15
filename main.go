package main

import (
	"fmt"
	"math"
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
	a, b := 5, 2
	fmt.Println(math.Ceil(float64(a) / float64(b)))
}
