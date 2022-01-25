package _538_convert_bst_to_greater_tree

import (
	"fmt"
	"testing"
)

func Test0538(t *testing.T) {
	root := &TreeNode{
		//5, &TreeNode{2, nil, nil}, &TreeNode{13, nil, nil},
		2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil},
	}

	convertBST(root)

	fmt.Println(root)

}
