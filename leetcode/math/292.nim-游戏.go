package leetcode

/*
 * @lc app=leetcode.cn id=292 lang=golang
 *
 * [292] Nim 游戏
 */

// @lc code=start
func canWinNim(n int) bool {
	return n%4 != 0
}

// @lc code=end

/*
	举例几个case：
	1 true
	2 true
	3 true
	4 false
	5 true
	6 true
	7 true
	8 false

	因为可以拿1-3块石头，所以，只要最终结果是1-3，那么就可以拿到
	但是如果结果是 4的话，在可以拿 1-3 的情况下，会剩下 3，2，1，那么对手就可以拿到

	如果再往上推。也就是说，最后4块该自己拿，那自己就输
	所以就控制让对手 剩4块的时候拿。如何让对手剩4块。那么自己可以拿 5，6，7。

	又推到了8，可以看到全是4的倍数。那么就是 4的倍数，谁拿谁输。

	因为每一步都是最优解。所以只需要第一步的时候，是4的倍数，谁拿谁输
*/
