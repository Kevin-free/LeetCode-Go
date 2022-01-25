package _076_minimum_window_substring

import (
	"fmt"
	"testing"
)

func Test0076(t *testing.T) {
	str := "ADOBECODEBANC"
	tar := "ABC"
	//str := "ab"
	//tar := "a"
	res := minWindow(str, tar)
	fmt.Print(res)
}