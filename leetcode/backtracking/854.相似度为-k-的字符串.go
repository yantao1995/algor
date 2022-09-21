package leetcode

/*
 * @lc app=leetcode.cn id=854 lang=golang
 *
 * [854] 相似度为 K 的字符串
 */

// @lc code=start
func kSimilarity(s1 string, s2 string) int {
	bs1, bs2 := []byte(s1), []byte(s2)
	minK := len(s1)
	var backtrace func(i, k int)
	backtrace = func(i, k int) {
		if k >= minK {
			return
		}
		for ; i < len(bs1); i++ {
			if bs1[i] != bs2[i] {
				k++
				if k >= minK {
					return
				}
				for j := i + 1; j < len(bs1); j++ {
					if bs1[j] == bs2[i] {
						bs1[i], bs1[j] = bs1[j], bs1[i]
						backtrace(i, k)
						bs1[i], bs1[j] = bs1[j], bs1[i]
					}
				}
			}
		}
		if k < minK {
			minK = k
		}
	}
	backtrace(0, 0)
	return minK
}

// @lc code=end

/*
	回溯
	为什么需要回溯，如下例子：
		 s1       s2
		ababa   bbaaa
	如果和第1个 b 交换 得到  baaba ,那么就还需要再次交换a和b得到bbaaa,总共就需要2次交换
	很显然，s1中第1个 a ,和 第2个 b 交换就只需要1次交换，即可达成 bbaaa

	所以程序并不知道在哪儿交换是最优的选择，所以需要依次尝试不同的位置。

	剪枝:
		1.当前 k>=minK 时，就应该抛弃本条路径，因为继续往下走，
		也只会大于等于mink,增加了无意义的计算
		2.外层循环内也判断当前 k>=minK ,可以减少内层重复元素的回溯调用
*/
