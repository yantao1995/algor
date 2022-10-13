package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	test2()
}

// func test() {
// 	fmt.Println(utils.GetRandQuestion(2088))
// }

func test2() {
	rd := rand.NewSource(math.MaxInt)
	for i := 0; i < 10; i++ {
		fmt.Print(" ", rd.Int63())
	}
}
