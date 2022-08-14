package leetcode

/*
 * @lc app=leetcode.cn id=768 lang=golang
 *
 * [768] 最多能完成排序的块 II
 */

// @lc code=start
func maxChunksToSorted(arr []int) int {
lab:
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j <= i; j++ {
			for k := i + 1; k < len(arr); k++ {
				if arr[k] < arr[j] {
					continue lab
				}
			}
		}
		return maxChunksToSorted(arr[:i+1]) + maxChunksToSorted(arr[i+1:])
	}
	return 1
}

// @lc code=end

/*
	同 769.最多能完成排序的块 ，一样分治法
	既然要分区，那就用分治法
	左区元素全部小于右区元素，那就直接划分成两个区。
	如果不能划区，那就只能为1个区，返回1
*/
