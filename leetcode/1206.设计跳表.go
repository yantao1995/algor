package leetcode

/*
 * @lc app=leetcode.cn id=1206 lang=golang
 *
 * [1206] 设计跳表
 */

// @lc code=start
type Skiplist struct {
	maxNodeLevel int           //最大层数结点的层数
	head, tail   *skipListNode //头尾结点
	count        int           //结点数量
}

type skipListNode struct {
	level []*nodeLevel
	prev  *skipListNode //前置结点
	score float64       //分值权重 【当前分值 = data值】
	data  []int         //存放相同数据
}

type nodeLevel struct {
	skipList *skipListNode //前进指针
	span     int           //到下一个结点的跨度
}

func Constructor() Skiplist {
	return Skiplist{}
}

func (this *Skiplist) Search(target int) bool {
	return true
}

func (this *Skiplist) Add(num int) {

}

func (this *Skiplist) Erase(num int) bool {
	return true
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
// @lc code=end
