package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}
