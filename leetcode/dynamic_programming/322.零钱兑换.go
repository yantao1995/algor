package leetcode

/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	repateMap := map[int]int{}
	return recursion322(coins, amount, repateMap)
}

func recursion322(coins []int, amount int, repateMap map[int]int) int {
	if v, ok := repateMap[amount]; ok {
		return v
	}
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	min := 2<<31 - 1
	for k := range coins {
		count := recursion322(coins, amount-coins[k], repateMap)
		if count == -1 {
			continue
		}
		if min > count+1 {
			min = count + 1
		}
	}
	repateMap[amount] = min
	if min != 2<<31-1 {
		return min
	}
	return -1
}

// @lc code=end
