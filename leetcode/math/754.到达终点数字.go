package leetcode

/*
 * @lc app=leetcode.cn id=754 lang=golang
 *
 * [754] 到达终点数字
 */

// @lc code=start
func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}
	numMoves := 0
	current := 0
	for current < target {
		numMoves++
		current += numMoves
	}
	if (current-target)%2 == 0 {
		return numMoves
	}
	return numMoves + 1 + numMoves%2
}

// @lc code=end

/*
	target 由 1-n个数排列相加或相减:
	1  2  3  4 5  .....  n
	相当于在其中填充正负号,如果当前找到了一个填充可以满足的，
		那么 target的正负 无非就是填充的正变负，负变正
	所以全部将 target 看成正的即可，如果为负，修改为正
	最短的策略为，当前全部为 + 号，只要相加的和current >= target,那么这个序列就有可能能完成
	如果 = ，那么直接完成
	如果 > , 那么分情况讨论:
		1. 如果 sub = current - target 为偶数，那么 sub / 2 的数只需要取一个负号，即可完成
			所以这个情况可以和 =  一样完成。即  （current-target）%2 == 0 ,也就是当前的数就是最小移动次数
		2. 如果 sub 为奇数，那么 当前的最后一个数为 偶数， 则 只需要 多走一步 奇数，来与 （sub）/2 +1 的奇数
			抵消即可。如果 当前的最后一个数为奇数，则需要再走出一个奇数来做抵消，也就是走一个偶数，再走一个奇数，
			那么需要多走两步。
*/
