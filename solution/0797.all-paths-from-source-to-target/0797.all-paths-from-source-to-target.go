package _797_all_paths_from_source_to_target

// 797. 所有可能的路径：中等
// tags：图、深度优先搜索、回溯

var target int //搜索目标
var dag [][]int //图的领接矩阵
var paths [][]int //存储所有路径
var path []int //存储当前搜索到的路径

// DFS 深度优先搜索
func DFS(cur int) {
	path = append(path, cur)
	defer func(){
		path = path[:len(path)-1]
	}()
	if cur == target {
		ans := make([]int, len(path))
		copy(ans, path)
		paths = append(paths, ans)
		return
	}
	for _, next := range dag[cur] {
		DFS(next)
	}
}

func allPathsSourceTarget(graph [][]int) [][]int {
	target = len(graph)-1
	dag = graph
	paths = make([][]int, 0)
	path = make([]int, 0)
	DFS(0)
	return paths
}