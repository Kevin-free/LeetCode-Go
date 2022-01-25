package _049_group_anagrams

func groupAnagrams(strs []string) [][]string {
	dic := map[byte]int{'a': 2, 'b': 3, 'c': 5, 'd': 7, 'e': 11, 'f': 13, 'g': 17, 'h': 19, 'i': 23, 'j': 29, 'k': 31, 'l': 37, 'm': 41,
		'n': 43, 'o': 47, 'p': 53, 'q': 59, 'r': 61, 's': 67, 't': 71, 'u': 73, 'v': 79, 'w': 83, 'x': 89, 'y': 97, 'z': 101,
	}
	resMap := make(map[int][]string)
	resList := make([][]string, 0)
	for _, str := range strs {
		m := 1
		for i := 0; i < len(str); i++ {
			m *= dic[str[i]]
		}
		if resMap[m] == nil {
			resMap[m] = []string{}
		}
		resMap[m] = append(resMap[m], str)
	}

	for _, v := range resMap {
		resList = append(resList, v)
	}

	return resList
}
