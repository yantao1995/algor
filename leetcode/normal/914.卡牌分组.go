package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=914 lang=golang
 *
 * [914] 卡牌分组
 */

// @lc code=start
func hasGroupsSizeX(deck []int) bool {
	sort.Ints(deck)
	minLength := len(deck)
	lengthSlice := []int{}
	lastIndex := 0
	for i := 0; i < len(deck); i++ {
		if i+1 < len(deck) && deck[i] != deck[i+1] {
			if i-lastIndex+1 < minLength {
				minLength = i - lastIndex + 1
			}
			lengthSlice = append(lengthSlice, i-lastIndex+1)
			lastIndex = i + 1
		}
	}
	if len(deck)-lastIndex < minLength {
		minLength = len(deck) - lastIndex
	}
	lengthSlice = append(lengthSlice, len(deck)-lastIndex)
	if minLength < 2 {
		return false
	}
	ht := getPrimeFactor914(minLength) //所有质因数
	htKeys := []int{}
	for k := range ht {
		htKeys = append(htKeys, k)
	}
	htKeys = append(htKeys, minLength) // 加上自己
	sort.Ints(htKeys)                  //小质数机会更大
	test := false
	for m := range htKeys {
		test = false
		for k := range lengthSlice {
			if lengthSlice[k]%htKeys[m] != 0 {
				test = true
				break
			}
		}
		if !test {
			break
		}
	}
	return !test
}

// 获取某个数的所有质因数
func getPrimeFactor914(n int) map[int]bool {
	//厄拉多塞筛法
	ht := map[int]bool{}
	for i := 2; i < n; i++ {
		ht[i] = true
	}
	temp := 0
	for i := 2; i < n; i++ {
		temp = i * 2
		if ht[i] {
			for temp < n {
				if ht[temp] {
					delete(ht, temp)
				}
				temp += i
			}
		}
	}
	for k := range ht { //获得质因数
		if n%k != 0 {
			delete(ht, k)
		}
	}
	return ht
}

// @lc code=end
