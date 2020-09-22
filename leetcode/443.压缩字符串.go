package leetcode

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=443 lang=golang
 *
 * [443] 压缩字符串
 */

// @lc code=start
func compress(chars []byte) int {
	length := 0
	last := byte(' ')
	count := 0
	lastIndex := 0
	for i := 0; i < len(chars); i++ {
		if chars[i] != last {
			if count == 1 {
				length++
				chars[lastIndex] = last
				lastIndex++
			}
			if count > 1 {
				length++
				length += getBit(count)
				chars[lastIndex] = last
				lastIndex++
				if last != ' ' {
					lengthBytes := []byte(strconv.Itoa(count))
					for k := range lengthBytes {
						chars[lastIndex] = lengthBytes[k]
						lastIndex++
					}
				}
			}
			last = chars[i]
			count = 1
			continue
		} else {
			count++
		}
	}
	chars[lastIndex] = last
	lastIndex++
	if count == 1 {
		length++
	}
	if count > 1 {
		length++
		length += getBit(count)
		lengthBytes := []byte(strconv.Itoa(count))
		if lengthBytes[0] == '1' && len(lengthBytes) == 1 {
			return length
		}
		for k := range lengthBytes {
			chars[lastIndex] = lengthBytes[k]
			lastIndex++
		}
	}
	return length
}

//获取位数
func getBit(num int) int {
	bt := 1
	weight := 1
	for weight*10 <= num {
		weight *= 10
		bt++
	}
	return bt
}

// @lc code=end
