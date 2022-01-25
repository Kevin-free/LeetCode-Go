package _206_reverse_linked_list

import (
	"fmt"
	"testing"
)

func Test0206(t *testing.T) {
	head := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	ans := reverseList(head)
	fmt.Print(ans)
}
