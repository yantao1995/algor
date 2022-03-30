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
	i := 0
	pi := &i
	*pi++
	fmt.Println(*pi)
	fmt.Println(*pi - 1)
	fmt.Println(*pi)
}
