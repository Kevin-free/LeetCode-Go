package _049_group_anagrams

import (
	"fmt"
	"testing"
)

func Test0049(t *testing.T) {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	output := groupAnagrams(input)
	fmt.Println(output)
}
