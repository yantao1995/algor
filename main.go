package main

import (
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
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	fmt.Println(gcd(2, 2))
}
