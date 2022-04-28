package leetcode

/*
 * @lc app=leetcode.cn id=190 lang=golang
 *
 * [190] 颠倒二进制位
 */

// @lc code=start
func reverseBits(num uint32) uint32 {
	result := uint32(0)
	for i := 0; i < 32; i++ {
		result <<= 1
		result += num & 1
		num >>= 1
	}
	return result
}

// @lc code=end

/*
	开始想用 num >0 作为条件，但是 uint32 的首位还有很多0，无法判断
	所以直接i<32来做判断
	然后result按位累加num的末尾
*/
