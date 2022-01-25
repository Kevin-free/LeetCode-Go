package _094_binary_tree_inorder_traversal

import (
	"fmt"
	"testing"
)

func Test0094(t *testing.T) {
	root := &TreeNode{
		1,
		nil, &TreeNode{2,
			&TreeNode{3, nil, nil}, nil},
	}
	res := inorderTraversal(root)
	fmt.Println(res)
}
