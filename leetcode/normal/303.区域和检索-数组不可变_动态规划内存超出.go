package leetcode

/*
 * @lc app=leetcode.cn id=303 lang=golang
 *
 * [303] 区域和检索 - 数组不可变
 */

// @lc code=start
// type NumArray struct {
// 	Array []int
// 	Table [][]int
// }

// func Constructor(nums []int) NumArray {
// 	table := [][]int{}
// 	for k := range nums {
// 		temp := make([]int, len(nums))
// 		for i := 0; i < len(temp); i++ {
// 			if i == k {
// 				temp[k] = nums[k]
// 			} else {
// 				temp[i] = -9999
// 			}
// 		}
// 		table = append(table, temp)
// 	}
// 	// 初始化备忘录，sumRange时也可以初始化，避免没必要的操作
// 	// i, j := 0, len(nums)-1
// 	// for m := 1; m < j-i+1; m++ {
// 	// 	for k := i; k+m <= j; k++ {
// 	// 		if table[k][k+m] == -9999 {
// 	// 			table[k][k+m] = table[k][k+m-1] + table[k+m][k+m]
// 	// 		}
// 	// 	}
// 	// }
// 	//
// 	return NumArray{
// 		Array: nums,
// 		Table: table,
// 	}
// }

// func (this *NumArray) SumRange(i int, j int) int {

// 	// 备忘录
// 	// if this.Table[i][j] == -9999 {
// 	// 	for m := 1; m < j-i+1; m++ {
// 	// 		for k := i; k+m <= j; k++ {
// 	// 			if this.Table[k][k+m] == -9999 {
// 	// 				this.Table[k][k+m] = this.Table[k][k+m-1] + this.Table[k+m][k+m]
// 	// 			}
// 	// 		}
// 	// 	}
// 	// }
// 	return this.Table[i][j]
// }

// func (this *NumArray) Ptfn() {
// 	for k := range this.Table {
// 		fmt.Println(this.Table[k])
// 	}
// 	fmt.Println("---------------------------")
// }

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
// @lc code=end
