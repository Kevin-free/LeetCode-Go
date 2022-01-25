package _543_diameter_of_binary_tree

import (
	"fmt"
	"testing"
)

func Test0543(t *testing.T) {
	root := &TreeNode{
		1,
		&TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}},
		&TreeNode{3, nil, nil},
	}

	ans := diameterOfBinaryTree(root)
	fmt.Println(ans)
}
