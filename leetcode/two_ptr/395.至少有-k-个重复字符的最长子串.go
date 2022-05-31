package leetcode

/*
 * @lc app=leetcode.cn id=395 lang=golang
 *
 * [395] 至少有 K 个重复字符的最长子串
 */

// @lc code=start
func longestSubstring(s string, k int) int {
	maxLength := 0
	cm := map[byte]int{}
	count := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		cm[s[i]]++
		count[i] = cm[s[i]]
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	temp := map[byte]int{}
	need := map[byte]bool{}
	got := map[byte]bool{}
	for i := len(count) - 1; i >= 0; i-- {
		temp = map[byte]int{}
		need = map[byte]bool{}
		got = map[byte]bool{}
		for j := i; j >= 0 && (count[i] >= k || need[s[j]] || got[s[j]]); j-- {
			if count[j] < k && !need[s[j]] && !got[s[j]] {
				if len(got) == 0 {
					i = j
				}
				break
			}
			if temp[s[j]] == 0 {
				need[s[j]] = true
				temp[s[j]] = count[j]
			}
			if temp[s[j]]-count[j]+1 >= k {
				got[s[j]] = true
				delete(need, s[j])
			}
			if len(need) == 0 {
				maxLength = max(maxLength, i-j+1)
			}
		}
	}
	return maxLength
}

// @lc code=end

/*
	第一时间想到了暴力破解,但是暴力肯定就超时了，暴力代码在最下面
		暴力思路：
			从当前i开始，每次检查从i到j长度的字符串是否满足条件，
			使用need标记当前j走过的长度内，哪些字符还不满足个数，
			使用count记录当前j走过的长度内，s[j]的个数是多少

	因为有些被重复计算了，所以就需要跳过某些元素
	比如 "aab" k=3 , 那么第一轮从0开始，a其实是不满足的，但是a从1开始又会走一遍

	所以：思路2，目地是跳过一些不需要重复计算的
	例如字符串 ："aaabbbcdefcdefgggggggggggggggcde", k=3
	可以定义 count 数组来表示每个字符到当前索引位置时，当前元素的累积个数，例如上面的字符串可以表示为：
	[1 2 3 1 2 3 1 1 1 1 2 2 2 2 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 3 3 3]
	外层循环从后向前遍历，只要当前元素的个数小于k，那么就跳过当前元素
	只要进入了当前循环，那么在 for j 时，当前元素个数就算小于k，只要之前遇到过，即 need或got中存在，
	那么就应该继续在for j 中向前遍历。
	need 记录当前元素还未满足k个，got记录当前元素已经满足k个，如果当前元素未出现过，
	而且count小于k,就进入退出判断条件，如果当前没有一个元素满足k个，那么 for i 就直接跳到当前 j 的位置继续，
	这样就跳过了一部分不满足k的元素
*/

/* 暴力超时
func longestSubstring(s string, k int) int {
	maxLength := 0
	need := map[byte]bool{}
	count := map[byte]int{}
	for i := 0; i < len(s)-maxLength; i++ {
		need = map[byte]bool{}
		count = map[byte]int{}
		for j := i; j < len(s); j++ {
			if count[s[j]] < k {
				need[s[j]] = true
			}
			count[s[j]]++
			if count[s[j]] >= k {
				delete(need, s[j])
			}
			if len(need) == 0 && j-i+1 > maxLength {
				maxLength = j - i + 1
			}
		}
	}
	return maxLength
}
*/
