package leetcode

/*
给定一个长度为 n 的整数数组 nums ，其中 nums 是范围为 [1，n] 的整数的排列。还提供了一个 2D 整数数组 sequences ，其中 sequences[i] 是 nums 的子序列。
检查 nums 是否是唯一的最短 超序列 。最短 超序列 是 长度最短 的序列，并且所有序列 sequences[i] 都是它的子序列。对于给定的数组 sequences ，可能存在多个有效的 超序列 。
例如，对于 sequences = [[1,2],[1,3]] ，有两个最短的 超序列 ，[1,2,3] 和 [1,3,2] 。
而对于 sequences = [[1,2],[1,3],[1,2,3]] ，唯一可能的最短 超序列 是 [1,2,3] 。[1,2,3,4] 是可能的超序列，但不是最短的。
如果 nums 是序列的唯一最短 超序列 ，则返回 true ，否则返回 false 。
子序列 是一个可以通过从另一个序列中删除一些元素或不删除任何元素，而不改变其余元素的顺序的序列。
*/
func sequenceReconstruction(nums []int, sequences [][]int) bool {
	type kv struct{ a, b int }
	m := map[kv]bool{}
	for k := range sequences {
		for i := 1; i < len(sequences[k]); i++ {
			m[kv{sequences[k][i-1], sequences[k][i]}] = true
		}
	}

	for i := 1; i < len(nums); i++ {
		if !m[kv{nums[i-1], nums[i]}] {
			return false
		}
	}
	return true
}

/*
	map 存储每一个拓扑前后映射关系
	需要nums唯一且最短，那么每个拓扑关系都应该在map中存在
	如果不存在，那么说明有特殊位置交换元素也可以组成新的序列
*/
