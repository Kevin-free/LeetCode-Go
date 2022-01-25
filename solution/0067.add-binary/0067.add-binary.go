package _067_add_binary

import (
	"myutil"
	"strconv"
)

func addBinary(a string, b string) string {
	ans := "" // 记录结果
	c := 0    // 是否进一位
	lena, lenb := len(a), len(b)
	n := myutil.MaxInt(lena, lenb) // a,b 中长度最大值

	for i := 0; i < n; i++ {
		if i < lena {
			c += int(a[lena-1-i] - '0')
		}
		if i < lenb {
			c += int(b[lenb-1-i] - '0')
		}
		// 注意：因为这里是 ans = sum + ans（而不是append尾加）所以最后不需要反转
		ans = strconv.Itoa(c%2) + ans // 如果二者都为1  那么c%2应该刚好为0 否则为1
		c /= 2 // 如果二者都为1  那么ca 应该刚好为1 否则为0
	}
	if c == 1 {
		ans = "1" + ans // 有进位，在前面加1
	}
	return ans
}
