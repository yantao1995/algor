package other

import (
	"fmt"
)

/*
	01 背包爆搜 递归版本
*/

func bagRecursion() {
	var (
		goods       = []string{"a", "b", "c", "d", "e"}
		values      = []int{2, 3, 4, 1, 4}
		weights     = []int{6, 3, 7, 4, 5}
		totalWeight = 10
		maxValue    = 0
	)
	var recursion func(chooseList []int, currentIndex, currentValue, currentWeight int)
	recursion = func(chooseList []int, currentIndex, currentValue, currentWeight int) {
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
	for i := 0; i < len(goods); i++ {
		recursion([]int{i}, i+1, values[i], weights[i])
	}
	fmt.Println("最大价值为：", maxValue)
}

func bagDP() {
	var (
		goods       = []string{"a", "b", "c", "d", "e"}
		values      = []int{2, 3, 4, 1, 4}
		weights     = []int{6, 3, 7, 4, 5}
		totalWeight = 10
	)
	dp := make([][]int, len(goods)) // 定义 二维数组dp
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, totalWeight+1) // 定义容量的大小为 0 到 背包的最大容量 totalWeight，所以要+1
	}
	//初始化dp 数组 的第一行 ， 只要容量大于等于 第0个物品， 那么价值就是第一件物品
	for i := 1; i <= totalWeight; i++ {
		if i >= weights[0] {
			dp[0][i] = values[0]
		}
	}
	//定义max 函数， C++内可以使用  math.max()
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// 开始进行 状态转移
	for i := 1; i < len(goods); i++ { //从第一个物品开始，因为第0个已经初始化了
		for j := 1; j <= totalWeight; j++ { //容量从1开始
			dp[i][j] = dp[i-1][j] //这个是不加当前物品，那么最大价值应该和 总共有 i-1 件物品时的最大价值一样
			if j >= weights[i] {  // 如果当前容量 j 的大小比 当前物品大，那么 j-当前物品 时的最大价值的基础上就可以加上当前物品
				dp[i][j] = max(dp[i][j], dp[i-1][j-weights[i]]+values[i]) // 状态转移   j-当前物品 时的最大价值的基础上就可以加上当前物品
			}
		}
	}
	// 打印一下 dp 数组 给你看看
	fmt.Println("打印dp数组:")
	n := len(dp)
	m := len(dp[0])
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			fmt.Print(dp[i][j], " ")
		}
		fmt.Println()
	}
	// 最大价值就是 dp[所有物品][最大容量] 的价值
	fmt.Println("最大价值为：", dp[n-1][m-1])
}
