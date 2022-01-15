package main

type Tree struct {
	root   TreeNode
	height int
}

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	data  int
}

func new(data int) Tree {
	return Tree{root: TreeNode{data: data}}
}

func main()  {
	
}


