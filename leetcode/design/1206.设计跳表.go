package leetcode

import (
	"fmt"
	"math/rand"
)

/*
 * @lc app=leetcode.cn id=1206 lang=golang
 *
 * [1206] 设计跳表
 */

// @lc code=start
type Skiplist struct {
	rd            *rand.Rand
	probable      float64         //层数生成概率
	constMaxLevel *int            //最大层数限制
	headLevel     []*skipListNode //头结点
	count         int             //结点数量
	getLevel      func() int      //获取生成结点的层高
}

type skipListNode struct {
	level []*skipListNode //前进指针
	score float64         //分值权重 【当前分值 = data值】
	data  []int           //存放相同数据
}

// func Constructor() Skiplist {
// 	maxLevel := 8
// 	sk := Skiplist{
// 		rd:            rand.New(rand.NewSource(time.Now().Unix())),
// 		probable:      0.5,
// 		constMaxLevel: &maxLevel,
// 		headLevel:     make([]*skipListNode, maxLevel),
// 		count:         0,
// 	}
// 	sk.getLevel = func() int {
// 		level := 1
// 		for level < *sk.constMaxLevel &&
// 			sk.probable <= sk.rd.Float64() {
// 			level++
// 		}
// 		return level
// 	}
// 	return sk
// }
func (this *Skiplist) search(target int) *skipListNode {
	if this.count > 0 {
		fTarget := float64(target)
		var node *skipListNode
		for level := len(this.headLevel) - 1; level >= 0; level-- {
			node = this.headLevel[level]
			if node != nil && node.score <= fTarget {
				for ; level >= 0; level-- {
					for {
						if node.score == fTarget {
							return node
						} else if node.level[level] != nil && node.level[level].score <= fTarget {
							node = node.level[level]
						} else {
							break
						}
					}
				}
			}
		}
	}
	return nil
}
func (this *Skiplist) Search(target int) bool {
	return this.search(target) != nil
}

func (this *Skiplist) Add(num int) {
	this.count++
	searchNode := this.search(num)
	if searchNode != nil {
		searchNode.data = append(searchNode.data, num)
		return
	}
	addNode := &skipListNode{
		level: make([]*skipListNode, this.getLevel()),
		score: float64(num),
		data:  []int{num},
	}
	for level := len(addNode.level) - 1; level >= 0; level-- {
		this.headLevel[level], addNode.level[level] = addNode, this.headLevel[level]
	}
	node := this.headLevel[0]
	for ; node.level[0] != nil && node.score > node.level[0].score; node = node.level[0] {
		node.score, node.data, node.level[0].score, node.level[0].data =
			node.level[0].score, node.level[0].data, node.score, node.data
	}
}

func (this *Skiplist) Erase(num int) bool {
	searchNode := this.search(num)
	if searchNode == nil {
		return false
	}
	this.count--
	if len(searchNode.data) > 1 {
		searchNode.data = searchNode.data[:len(searchNode.data)-1]
	} else {
		var node *skipListNode
		for level := len(searchNode.level) - 1; level >= 0; level-- {
			if this.headLevel[level] == searchNode {
				this.headLevel[level] = searchNode.level[level]
				continue
			}
			if node == nil {
				node = this.headLevel[level]
			}
			for ; node.level[level] != nil; node = node.level[level] {
				if node.level[level] == searchNode {
					node.level[level] = searchNode.level[level]
					break
				}
			}
		}
	}
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

func (this *Skiplist) Print() {
	for level := len(this.headLevel) - 1; level >= 0; level-- {
		fmt.Printf("headLevel[%d]:", level)
		for node := this.headLevel[0]; node != nil; node = node.level[0] {
			if len(node.level) > level {
				fmt.Printf("\t-->\t%d%d", node.data[0], len(node.data))
			} else {
				if node.level[0] != nil {
					fmt.Printf("\t---\t---")
				} else {
					fmt.Printf("\t-->\tnil")
				}
			}
		}
		fmt.Println()
	}
	fmt.Println("--------------------------")
}

/*
	常规实现

type Skiplist struct {
	rd            *rand.Rand		//用于生成层高
	probable      float64         //层数生成概率
	constMaxLevel *int            //最大层数限制
	headLevel     []*skipListNode //头结点的层数映射
	count         int             //结点数量
	getLevel      func() int      //获取生成结点的层高  【概率指数，每加一层，都要算一次概率】
}

type skipListNode struct {
	level []*skipListNode //前进指针
	score float64         //分值权重 【当前分值 = data值】
	data  []int           //存放相同数据
}

	添加结点:
		直接添加到第一个位置，然后交换score和data到指定位置
	删除结点：
		data存储多个，则不删除结点，只删除数据
		data只有1个，则需要删除结点,层序遍历，找到指定结点之后，前一个结点的指针直接指向后一个结点
	查找：
		从顶层依次向下查找
*/
