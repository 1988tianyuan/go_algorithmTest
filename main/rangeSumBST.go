package main

import "fmt"

func main()  {
	root := new(TreeNode)
	root.Val = 10
	root.Left = &TreeNode{Val:5, Left:&TreeNode{Val:3, Left:nil, Right:nil}, Right:&TreeNode{Val:7, Left:nil, Right:nil}}
	root.Right = &TreeNode{Val:15, Left:nil, Right:&TreeNode{Val:18, Left:nil, Right:nil}}
	fmt.Println(rangeSumBST(root, 7, 15))
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	var sum int
	rootVal := root.Val
	if rootVal < L {
		sum = rangeSumBST(root.Right, L, R)
	} else if rootVal > R {
		sum = rangeSumBST(root.Left, L, R)
	} else {
		sum = rangeSumBST(root.Left, L, rootVal) + rangeSumBST(root.Right, rootVal, R) + rootVal
	}
	return sum
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}