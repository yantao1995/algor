package easy

/*
 * @lc app=leetcode.cn id=605 lang=golang
 *
 * [605] 种花问题
 */

// @lc code=start
func canPlaceFlowers(flowerbed []int, n int) bool {
	result := 0
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			if i+2 < len(flowerbed) {
				if i == 0 && flowerbed[i+1] == 0 {
					flowerbed[i] = 1
					result++
				} else if flowerbed[i+1] == 0 && flowerbed[i+2] == 0 {
					flowerbed[i+1] = 1
					result++
					i = i + 1
				}
			} else {
				if i == len(flowerbed)-2 && flowerbed[len(flowerbed)-1] == 0 {
					flowerbed[len(flowerbed)-1] = 1
					result++
					break
				}
			}
		}
	}
	if len(flowerbed) == 1 && flowerbed[0] == 0 {
		result = 1
	}
	if result >= n {
		return true
	}
	return false
}

// @lc code=end
