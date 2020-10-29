package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

/*树深*/
func GetDepth(node *Node) int {
	if node == null {
		return 0
	}
	ld := GetDepth(node.Left) + 1
	rd := GetDepth(node.Right) + 1
	if ld > rd {
		return ld
	} else {
		return rd
	}
}

/*第K层的节点个数*/
func GetKLevel(node *Node, k int) int {
	if node == nil || k < 1 {
		return 0
	}
	if k == 1 {
		return 1
	}
	return GetKLevel(node.Left, k-1) + GetKLevel(node.Right, k-1)
}

/*两个二叉树结构是否相同*/
func StructurcCmp(node1, node2 *Node) bool {
	if node1 == nil && node2 == nil {
		return true
	} else if node1 == nil || node2 == nil {
		return false
	}
	l := StructurcCmp(node1.Left, node2.Left)
	r := StructurcCmp(node1.Right, node2.Right)
	return l && r
}

func main() {
	fmt.Println("vim-go")
}
