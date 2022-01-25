package solution

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	length := len(candidates)
	sort.Ints(candidates)

	dfs(candidates, length, target, 0, nil, &res)
	return res
}

//
func dfs(candidates []int, length int, residue int, begin int, path []int, res *[][]int) {
	if residue == 0 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return

		//*res = append(*res, path)
		//return
	}

	for i := begin; i < length; i++ {
		if residue < candidates[i] {
			break
		}
		dfs(candidates, length, residue-candidates[i], i, append(path, candidates[i]), res)
	}
}

