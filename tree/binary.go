package tree

import (
	"algor/queue"
	"algor/stack"
	"algor/vals"
	"fmt"
)

type Binary struct {
	Data  vals.AlgorType
	Left  *Binary
	Right *Binary
}

func InitBinary(dt vals.AlgorType) *Binary {
	return &Binary{
		Data:  dt,
		Left:  nil,
		Right: nil,
	}
}

//先
func FirstErgodic(root *Binary) {
	if root != nil {
		fmt.Printf("%v\t", root.Data)
		FirstErgodic(root.Left)
		FirstErgodic(root.Right)
	}
}

//中
func MiddleErgodic(root *Binary) {
	if root != nil {
		MiddleErgodic(root.Left)
		fmt.Printf("%v\t", root.Data)
		MiddleErgodic(root.Right)
	}
}

//后
func LastErgodic(root *Binary) {
	if root != nil {
		LastErgodic(root.Left)
		LastErgodic(root.Right)
		fmt.Printf("%v\t", root.Data)
	}
}

//先 非递归
func FirstErgodicNoRecursion(root *Binary) {
	stack := stack.Stack{}
	for root != nil || stack.Size != 0 {
		if root != nil {
			fmt.Printf("%v\t", root.Data)
			stack.Push(root)
			root = root.Left
		} else {
			root = stack.Pop().(*Binary)
			root = root.Right
		}
	}
}

//中 非递归
func MiddleErgodicNoRecursion(root *Binary) {
	stack := stack.Stack{}
	for root != nil || stack.Size != 0 {
		if root != nil {
			stack.Push(root)
			root = root.Left
		} else {
			root = stack.Pop().(*Binary)
			fmt.Printf("%v\t", root.Data)
			root = root.Right
		}
	}
}

/*
思路(CSDN博主  千与千寻之前)：
	首先要搞清楚什么时候才能输出根节点的值，必须等到左节点和右节点都访问完的情况才能访问根节点
  * 因此，访问根节点的情况有两种：
  * 一、当前节点的左右子树都为nil时，可以直接访问根节点
  * 二、当前节点的左右子树都访问过时，则可以直接访问根节点
  * 则可以设置一个标志，让他标识上一个访问的结点
  * 遍历栈时，得到栈顶元素，判断他的左右子树为空或者prev标志是否等于当前结点的左右孩子的其中一个时，
  * 当这两种情况满足其中一种，则就可以访问当前根节点。
  *
  * 至于为什么要判断prev标志是否等于当前结点的左右孩子的其中一个，而不是判断它是否等于当前结点的右孩子？
  * 原因是这样的：因为可能存在这样一种情况，当前结点只有左孩子结点，则如果去判断它是否等于当前结点的右孩子结点
  * 得到的一定是false，这样当前结点永远都不会被访问，并且会陷入死循环，左孩子一直循环入栈的情况。当根结点左
  * 子树访问完，如果存在右子树，则会从栈顶得到右子树结点，而此时prev标志等于根节点的左子结点，因此第二个条件不成立（
  * 即prev不会等于根节点右子结点的任何一个子结点）
*/
//后 非递归
func LastErgodicNoRecursion(root *Binary) {
	if root == nil {
		return
	}
	stack := stack.Stack{}
	stack.Push(root)
	var prev *Binary = nil
	var current *Binary = nil
	for stack.Size != 0 {
		current = stack.GetTopButNoPop().(*Binary)
		//fmt.Println("current:", current, " prev:", prev, " size:", stack.Size, "if:", (current.Left == nil && current.Right == nil) || (prev != nil && (current.Left == prev || current.Right == prev)))
		if (current.Left == nil && current.Right == nil) || (prev != nil && (current.Left == prev || current.Right == prev)) {
			fmt.Printf("%v\t", current.Data)
			stack.Pop()
			prev = current
		} else {
			if current.Right != nil {
				stack.Push(current.Right)
			}
			if current.Left != nil {
				stack.Push(current.Left)
			}
		}
	}
}

//深度  栈  = 非递归先序
func DeepErgodic(root *Binary) {
	if root == nil {
		return
	}
	stack := stack.Stack{}
	stack.Push(root)
	for stack.Size > 0 {
		node := stack.Pop().(*Binary)
		fmt.Printf("%v\t", node.Data)
		if node.Right != nil {
			stack.Push(node.Right)
		}
		if node.Left != nil {
			stack.Push(node.Left)
		}
	}
}

//广度 层次 队列
func WideErgodic(root *Binary) {
	if root == nil {
		return
	}
	q := queue.Queue{}
	q.LPush(root)
	for q.Size != 0 {
		node := q.RPop().(*Binary)
		fmt.Printf("%v\t", node.Data)
		if node.Left != nil {
			q.LPush(node.Left)
		}
		if node.Right != nil {
			q.LPush(node.Right)
		}
	}
}

//拷贝
func CopyBinary(root *Binary) *Binary {
	newRoot := new(Binary)
	if root != nil {
		newRoot.Data = root.Data
		newRoot.Left = CopyBinary(root.Left)
		newRoot.Right = CopyBinary(root.Right)
		return newRoot
	}
	return nil
}

/*  测试结构
		  a
	b_____|_____c
d___|___e   f___|___g

*/

func TestBinary() {
	println("递归，先中后")
	root := InitBinary("a")
	LoadTestBinary(root)
	FirstErgodic(root)
	println()
	MiddleErgodic(root)
	println()
	LastErgodic(root)
	println("\n非递归，先中后")
	FirstErgodicNoRecursion(root)
	println()
	MiddleErgodicNoRecursion(root)
	println()
	LastErgodicNoRecursion(root)
	println("\n非递归，深度")
	DeepErgodic(root)
	println("\n非递归，广度")
	WideErgodic(root)
	newRoot := CopyBinary(root)
	println("\nnewRootCopy,广度")
	WideErgodic(newRoot)
}

// 加载treenodes
func LoadTestBinary(root *Binary) {
	root.Left = &Binary{
		Data: "b",
		Left: &Binary{
			Data:  "d",
			Left:  nil,
			Right: nil,
		},
		Right: &Binary{
			Data:  "e",
			Left:  nil,
			Right: nil,
		},
	}

	root.Right = &Binary{
		Data: "c",
		Left: &Binary{
			Data:  "f",
			Left:  nil,
			Right: nil,
		},
		Right: &Binary{
			Data:  "g",
			Left:  nil,
			Right: nil,
		},
	}
}
