package leetcode

/*
 * @lc app=leetcode.cn id=658 lang=golang
 *
 * [658] 找到 K 个最接近的元素
 */

// @lc code=start
func findClosestElements(arr []int, k int, x int) []int {
	if x <= arr[0] {
		return arr[:k]
	}
	if x >= arr[len(arr)-1] {
		return arr[len(arr)-k:]
	}
	current := 0
	left, right := 0, len(arr)-1
	for mid := (left + right) / 2; left+1 < right; mid = (left + right) / 2 {
		if x > arr[mid] {
			left = mid
		} else if x < arr[mid] {
			right = mid
		} else {
			left = mid - 1
			right = mid + 1
			current++
			break
		}
	}
	getAbs := func(a, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}
	lr, rr := 0, 0
	for {
		if right < len(arr) {
			rr = getAbs(x, arr[right])
		}
		if left >= 0 {
			lr = getAbs(x, arr[left])
		}
		if (left >= 0 && lr <= rr) || right >= len(arr) {
			if current+1 == k {
				return arr[left : left+k]
			}
			left--
		} else {
			right++
			if current+1 == k {
				return arr[right-k : right]
			}
		}
		current++
	}
}

// @lc code=end

/*
	二分法，找到该x所在位置的左右区间
	left,right 分别为左右区间，current记录当前找到的数的个数
	getAbs 获取当前两个数的绝对值
	只需要找到满足k个数的左右区间即可
*/
