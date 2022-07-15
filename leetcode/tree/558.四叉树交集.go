package leetcode

/*
 * @lc app=leetcode.cn id=558 lang=golang
 *
 * [558] 四叉树交集
 */

// @lc code=start
/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

// func intersect(quadTree1 *Node, quadTree2 *Node) *Node {
// 	if quadTree1.IsLeaf && quadTree2.IsLeaf {
// 		if quadTree1.Val {
// 			return quadTree1
// 		}
// 		return quadTree2
// 	} else if quadTree1.IsLeaf {
// 		if quadTree1.Val {
// 			return quadTree1
// 		}
// 		return quadTree2
// 	} else if quadTree2.IsLeaf {
// 		if quadTree2.Val {
// 			return quadTree2
// 		}
// 		return quadTree1
// 	} else {
// 		var node *Node = &Node{false, false, nil, nil, nil, nil}
// 		node.TopLeft = intersect(quadTree1.TopLeft, quadTree2.TopLeft)
// 		node.TopRight = intersect(quadTree1.TopRight, quadTree2.TopRight)
// 		node.BottomLeft = intersect(quadTree1.BottomLeft, quadTree2.BottomLeft)
// 		node.BottomRight = intersect(quadTree1.BottomRight, quadTree2.BottomRight)
// 		if node.TopLeft != nil && node.TopRight != nil &&
// 			node.BottomLeft != nil && node.BottomRight != nil &&

// 			node.TopLeft.IsLeaf && node.TopRight.IsLeaf &&
// 			node.BottomLeft.IsLeaf && node.BottomRight.IsLeaf &&

// 			node.TopLeft.Val == node.TopRight.Val &&
// 			node.TopRight.Val == node.BottomLeft.Val &&
// 			node.BottomLeft.Val == node.BottomRight.Val {
// 			node.IsLeaf = true
// 			node.Val = node.TopLeft.Val
// 			node.TopLeft, node.TopRight = nil, nil
// 			node.BottomLeft, node.BottomRight = nil, nil
// 		}
// 		return node
// 	}
// }

// @lc code=end

/*
	题意，4叉树映射成二维矩阵，将两个矩阵内的每个对应值进行按位 或运算
	然后将网格4等分，如果每份区域内的值数字全为1或者全为0，那么就是叶子结点
		即： IsLeaf = true   val = 1或者1
	所以非叶子结点的val是无效的，可以为任意值

	思路：
		不必进行nil判断，因为在判断某一个为叶子结点的时候，就已经返回了
		1.树1或者树2 有一个为叶子结点，而另一棵树还没有到叶子结点，
			那么可以判断当前叶子结点的值是否为1：
			如果为1，那么按位或就会将所有值置为1，那么当前也会变成叶子结点，
			那么直接返回当前叶子结点即可。
			如果为0，那么不会影响另一个子树的结果，那么直接返回另一个子树的结果即可。
		2.两个都为叶子结点
			那么可以直接对两个值按位或，然后返回结果值。
			对应到代码关系就是只有有一个val 是true，就直接返回任意一个为true的
			否则返回任意一个为false的
		3.两个都不为叶子结点。那么应该对4个位置都进行递归处理，然后再判断返回的
			4个结点的值是否相同，如果值都一样，那么此时的结点退化成叶子结点。
*/
