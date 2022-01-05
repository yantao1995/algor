package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func main() {
	test()
}

func test() {
	a := []int(nil)
	b := []int{}
	fmt.Println(reflect.DeepEqual(a, b))
	fmt.Println(reflect.TypeOf(b))

	//fmt.Println(getRandQuestion())
}

func getRandQuestion() int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(2088)
}
