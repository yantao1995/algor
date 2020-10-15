package leetcode

import "strconv"

/*
 * @lc app=leetcode.cn id=788 lang=golang
 *
 * [788] 旋转数字
 */

// @lc code=start
func rotatedDigits(N int) int {
	ht := map[byte]bool{
		'0': true,
		'1': true,
		'2': true,
		'5': true,
		'6': true,
		'8': true,
		'9': true,
	}
	ht2 := map[byte]byte{
		'0': '0',
		'1': '1',
		'2': '5',
		'5': '2',
		'6': '9',
		'8': '8',
		'9': '6',
	}
	count := 0
	for i := 1; i <= N; i++ {
		temp := strconv.Itoa(i)
		temp2 := make([]byte, len(temp))
		for k := range temp {
			if !ht[temp[k]] {
				goto lab
			} else {
				temp2[k] = ht2[temp[k]]
			}
		}
		if string(temp2) != temp {
			count++
		}
	lab:
	}
	return count
}

// @lc code=end

// func rotatedDigits(N int) int {
// 	ht := map[int]bool{
// 		2: true,
// 		5: true,
// 		6: true,
// 		9: true,
// 	}
// 	bit := 1
// 	for bit < N {
// 		bit *= 10
// 	}
// 	count := 0
// 	for i := 1; i <= N; i++ {
// 		temp := i
// 		tempBit := bit
// 		//	fmt.Println("\n1temp:", temp, "tempBit", tempBit)
// 		for temp > 0 {
// 			for temp < tempBit {
// 				tempBit /= 10
// 			}
// 			//	fmt.Print("     temp:  ", temp, "tempBit:  ", tempBit)
// 			if !ht[temp/tempBit] {
// 				//	i += tempBit / 10
// 				goto Label
// 			}
// 			temp -= temp / tempBit * tempBit
// 		}
// 		count++
// 	Label:
// 	}
// 	return count
// }
