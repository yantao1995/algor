package leetcode

import "sort"

/*
给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。

示例 1：
输入: s1 = "abc", s2 = "bca"
输出: true

示例 2：
输入: s1 = "abc", s2 = "bad"
输出: false

说明：
0 <= len(s1) <= 100
0 <= len(s2) <= 100

*/
func CheckPermutation(s1 string, s2 string) bool {
	bts1 := []byte(s1)
	bts2 := []byte(s2)
	sort.Slice(bts1, func(i, j int) bool {
		return bts1[i] < bts1[j]
	})
	sort.Slice(bts2, func(i, j int) bool {
		return bts2[i] < bts2[j]
	})
	return string(bts1) == string(bts2)
}

/*
	排序相等即可，或者统计每个字符的个数是否对应
*/
