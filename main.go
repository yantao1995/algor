package main

import "fmt"

func main() {
	test2()
}

// func test() {
// 	fmt.Println(utils.GetRandQuestion(2088))
// }

var (
	goods       = []string{"a", "b", "c", "d"}
	values      = []int{2, 3, 4, 1, 4}
	weights     = []int{6, 3, 7, 4, 5}
	totalWeight = 10
	maxValue    = 0
)

func test2() {
	for i := 0; i < len(goods); i++ {
		recursion([]int{i}, i+1, values[i], weights[i])
	}
	fmt.Println("最大价值为：", maxValue, " 最大重量为:", totalWeight)
}

func recursion(chooseList []int, currentIndex, currentValue, currentWeight int) {
	if currentWeight > totalWeight {
		return
	}
	fmt.Println("物品列表:", chooseList, " 当前的总价值:", currentValue, " 当前的总重量", currentWeight)
	if currentValue > maxValue {
		maxValue = currentValue
	}
	for i := currentIndex; i < len(goods); i++ {
		recursion(append(chooseList, i), i+1, currentValue+values[i], currentWeight+weights[i])
	}
}
