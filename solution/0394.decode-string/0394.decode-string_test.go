package _394_decode_string

import (
	"fmt"
	"testing"
)

func Test0394(t *testing.T) {
	in := "100[leetcode]"
	//in := "3[a]2[bc]"
	out := decodeString(in)
	fmt.Println(out)
}
