package _076_minimum_window_substring

import "math"

// 76.最小覆盖子串：困难
// tag: 哈希表、字符串、滑动窗口

// 如果一个字符进入窗口，应该增加window计数器；
// 如果一个字符将移出窗口的时候，应该减少window计数器；
// 当valid满足need时应该收缩窗口；应该在收缩窗口的时候更新最终结果。
func minWindow(s string, t string) string {
	need, window := map[byte]int{}, map[byte]int{}
	for i, _ := range t {
		need[t[i]]++
	}

	left, right := 0, 0
	// 记录覆盖子串的个数
	valid := 0
	// 记录最小覆盖子串的起始索引及长度
	start, length := 0, math.MaxInt32
	//length := math.MaxInt8
	for right < len(s) {
		// c 是将移入窗口的字符
		c := s[right]
		// 右移窗口
		right++
		// 进行窗口内数据的一系列更新
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		// 判断左侧窗口是否要收缩
		for valid == len(need) {
			// 在这里更新最小覆盖子串
			if right-left < length {
				start = left
				length = right - left
			}
			// d 是将移除窗口的字符
			d := s[left]
			// 左移窗口
			left++
			// 进行窗口内数据的一系列更新
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	// 返回最小覆盖子串
	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length]
}
