package _394_decode_string

func decodeString(s string) string {
	res := "" // 结果字符串
	multi := 0 // 倍数
	smulti, sstring := make([]int, 0), make([]string, 0) // 数字栈 和 字符串栈
	for _, c := range s { // 逐个字符扫描
		if c == '[' { // 遇到 【
			smulti = append(smulti, multi) // 数字入栈
			sstring = append(sstring, res) // 字符串入栈
			multi = 0 //清零
			res = "" //清空
		}else if c == ']' { // 遇到 】
			tmp := "" // 临时变量
			curMulti := smulti[len(smulti)-1] //拿出数字
			smulti = smulti[:len(smulti)-1] //⚠️go中别忘了加这句=移除最后一个元素
			for i := 0; i < curMulti; i++ { //重复数字次
				tmp += res
			}
			res = sstring[len(sstring)-1] + tmp // 构建子串
			sstring = sstring[:len(sstring)-1] //
		}else if c >= '0' && c <= '9' {
			multi = multi * 10 + int(c) - '0' //⚠️考虑数字是两位数以上的情况
		}else {
			res += string(c) // 拼接字符串
		}
	}
	return res //返回结果
}

//func decodeString(s string) string {
//	numStack := []int{}
//	strStack := []string{}
//	num := 0
//	result := ""
//	for _, char := range s {
//		if char >= '0' && char <= '9' {
//			n, _ := strconv.Atoi(string(char))
//			num = num*10 + n
//		} else if char == '[' {
//			strStack = append(strStack, result)
//			result = ""
//			numStack = append(numStack, num)
//			num = 0
//		} else if char == ']' {
//			count := numStack[len(numStack)-1]
//			numStack = numStack[:len(numStack)-1]
//			str := strStack[len(strStack)-1]
//			strStack = strStack[:len(strStack)-1]
//			result = string(str) + strings.Repeat(result, count)
//		} else {
//			result += string(char)
//		}
//	}
//	return result
//}