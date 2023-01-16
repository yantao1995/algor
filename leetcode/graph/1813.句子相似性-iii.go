package leetcode

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=1813 lang=golang
 *
 * [1813] 句子相似性 III
 */

// @lc code=start
func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	s1, s2 := strings.Split(sentence1, " "), strings.Split(sentence2, " ")
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}
	if len(s1) == len(s2) {
		return sentence1 == sentence2
	}
	flag := false
	var dfs func(i, j int, bits []int)
	dfs = func(i, j int, bits []int) {
		if flag { //剪枝
			return
		}
		if i >= len(s1) {
			count := 0
			if bits[0] > 0 {
				count++
			}
			ti := 1
			for ; ti < len(bits); ti++ {
				if bits[ti]-bits[ti-1] > 1 {
					count++
				}
			}
			if bits[ti-1] != len(s2)-1 {
				count++
			}
			if count < 2 {
				flag = true
			}
			return
		}
		for ; j < len(s2); j++ {
			if s1[i] == s2[j] {
				dfs(i+1, j, append(bits, j))
			}
		}
	}
	dfs(0, 0, []int{})
	return flag
}

// @lc code=end

/*
	将 sentence1 和 sentence2 按空格分隔成单词数组
	保证s1比s2小，如果相等，那么直接返回两个句子是否相等即可
	然后将 s1 按顺序放入 s2 内，如果有多种选择，那么进入dfs选择
	最后得到放入顺序序列 bits
	然后开始处理 bits 序列
	可以分为3中情况(导致计数器count加1)：
		第0个不为0： 需要在开头插入句子
		末尾一个不为s2末尾一个： 需要在中间插入句子
		中间不连续：需要在末尾插入句子
	所以 count 只能小于2，即最多插入一个句子，通过flag带出结果
*/
