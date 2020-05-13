package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//简单暴力的分层宽度遍历方法
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		tmp := []int{}
		tmpNodes := []*TreeNode{}
		for i := 0; i < len(nodes); i++ {
			tmp = append(tmp, nodes[i].Val)
			if nodes[i].Left != nil {
				tmpNodes = append(tmpNodes, nodes[i].Left)
			}
			if nodes[i].Right != nil {
				tmpNodes = append(tmpNodes, nodes[i].Right)
			}
		}
		res = append(res, tmp)
		nodes = tmpNodes
	}
	return res
}
