package leetcode

/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	ll := (len(nums1) + len(nums2)) - 2
// 	mid := ll / 2
// 	if ll%2 == 1 {
// 		mid = ll/2 + 1
// 	}
// 	if ll%2 == -1 {
// 		mid = 0
// 	}
// 	mid1 := mid + 1
// 	midnum, mid1num := 0, 0
// 	temp, count := 0, 0
// 	fmt.Println("len(num1):", len(nums1), "len(nums2):", len(nums2), "mid:", mid, "mid1:", mid1)
// 	for i, j := 0, 0; ; {
// 		devided1, devided2 := 0, 0
// 		if i < len(nums1) {
// 			devided1 = temp - nums1[i]
// 		}
// 		if j < len(nums2) {
// 			devided2 = temp - nums2[j]
// 		}
// 		if (devided1 >= devided2 && i < len(nums1)) || j >= len(nums2) { //差值小
// 			temp = nums1[i]
// 			i++
// 		} else {
// 			temp = nums2[j]
// 			j++
// 		}
// 		if count == mid {
// 			midnum = temp
// 			if ll%2 == 1 || (mid == 0 && ll%2 == -1) {
// 				return float64(temp)
// 			}
// 		}
// 		if count == mid1 {
// 			mid1num = temp
// 			break
// 		}
// 		fmt.Println(i, "i _ j", j, "temp:", temp, "count:", count, "devide1:", devided1, "devide2:", devided2)
// 		count++
// 	}
// 	fmt.Println(midnum, "-", mid1num)
// 	return (float64(midnum) + float64(mid1num)) / 2
// }

// @lc code=end
