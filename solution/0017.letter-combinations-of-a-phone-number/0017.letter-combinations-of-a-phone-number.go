package _017_letter_combinations_of_a_phone_number

var res []string
var dic = []string{"","","abc","def",
"ghi","jkl","mno",
"pqrs","tuv","wxyz"}

func letterCombinations(digits string) []string {
	res = []string{}
	n := len(digits)
	if n == 0 {
		return res
	}

	findCombination(digits, 0, "")
	return res
}

func findCombination(digits string, index int, s string) {
	// 长度相等，返回
	if index == len(digits) {
		res = append(res, s)
		return
	}
	b := int(digits[index]-'0') // 对应的数字（digits[index]拿出的是数字对应的ASCII码，所以减去0）
	letters := dic[b] // 数字对应的字符串
	for i := 0; i < len(letters); i++ {
		findCombination(digits, index+1, s+string(letters[i]))
	}
	return
}
