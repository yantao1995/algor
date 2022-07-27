package leetcode

import (
	"fmt"
	"math/rand"
	"time"
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
	getLevel      func() int      //获取生成结点的数量
}

type skipListNode struct {
	level []*skipListNode //前进指针
	score float64         //分值权重 【当前分值 = data值】
	data  []int           //存放相同数据
}

func Constructor() Skiplist {
	maxLevel := 8
	sk := Skiplist{
		rd:            rand.New(rand.NewSource(time.Now().Unix())),
		probable:      0.5,
		constMaxLevel: &maxLevel,
		headLevel:     make([]*skipListNode, maxLevel),
		count:         0,
	}
	sk.getLevel = func() int {
		level := 1
		for level <= *sk.constMaxLevel &&
			sk.probable <= sk.rd.Float64() {
			level++
		}
		return level
	}
	return sk
}
func (this *Skiplist) search(target int) *skipListNode {
	if this.count == 0 {
		return nil
	}
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
	if len(searchNode.data) > 1 {
		searchNode.data = searchNode.data[:len(searchNode.data)-1]
	} else {
		score, data := 0, []int{}
		for node := this.headLevel[0]; node != nil; node = node.level[0] {

		}
	}
	this.count--
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
			if len(node.level) >= level {
				fmt.Printf("\t-->\t%d[%d", node.data[0], len(node.data))
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
备份，常规添加
func (this *Skiplist) Add(num int) {
	this.count++
	addNode := &skipListNode{
		level: make([]*skipListNode, this.getLevel()),
		score: float64(num),
		data:  []int{num},
	}
	var node *skipListNode
	for level := len(addNode.level) - 1; level >= 0; level-- {
		node = this.headLevel[level]
		if node == nil {
			this.headLevel[level] = addNode
		} else {
			if node.score == addNode.score {
				node.data = append(node.data, num)
				node.level = append(node.level, addNode.level[level+1:]...)
				return
			} else if node.score < addNode.score {
				for ; level >= 0; level-- {
					for {
						if node.score == addNode.score {
							node.data = append(node.data, num)
							node.level = append(node.level, addNode.level[level+1:]...)
							return
						} else if node.level[level] != nil && node.level[level].score <= addNode.score {
							node = node.level[level]
						} else {
							this.headLevel[level], addNode.level[level] = addNode, this.headLevel[level]
							break
						}
					}
				}
			} else if node.score > addNode.score {
				this.headLevel[level], addNode.level[level] = addNode, this.headLevel[level]
			}
		}
	}
}
*/
