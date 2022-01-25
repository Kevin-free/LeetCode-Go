package _621_task_scheduler

import (
	"sort"
)

func leastInterval(tasks []byte, n int) int {
	// 用于记录各个任务的次数，范围为26大写英文字母
	count := make([]int, 26)

	// 统计各个任务的次数
	for _, task := range tasks {
		count[task-'A']++ // 因为是 byte，需要处理
	}

	// 按次数排序，默认升序，count[25]是频率最高的
	sort.Ints(count)

	// 次数最多的任务数
	maxCount := 0

	// 统计有多少个频率最高的任务
	for i := 25; i >= 0; i-- {
		if count[i] != count[25] {
			break
		}
		maxCount++
	}

	// 公式算出的值可能会比数组的长度小，取两者中最大的那个
	return max((count[25]-1)*(n+1) + maxCount, len(tasks))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}