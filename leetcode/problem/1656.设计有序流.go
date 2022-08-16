package leetcode

/*
 * @lc app=leetcode.cn id=1656 lang=golang
 *
 * [1656] 设计有序流
 */

// @lc code=start
type OrderedStream struct {
	stream []string
	ptr    int
}

// func Constructor(n int) OrderedStream {
// 	return OrderedStream{
// 		stream: make([]string, n),
// 		ptr:    0,
// 	}
// }

func (this *OrderedStream) Insert(idKey int, value string) []string {
	result := []string{}
	idKey--
	this.stream[idKey] = value
	for this.ptr < len(this.stream) && this.stream[this.ptr] != "" {
		result = append(result, this.stream[this.ptr])
		this.ptr++
	}
	return result
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
// @lc code=end

/*
	stream 存储value流，根据 idKey-1 的值来存储，因为输入的数据是 [1 ~ n],而索引是 [0  ~ n-1]
	ptr 存储当前指针
	每当有 insert 时，都先向 stream[idKey-1] 的位置存储 value
	然后ptr判断当前位置是否占用，如果占用就加入到结果集，ptr向后移动，然后再循环判断是否占用........
*/
