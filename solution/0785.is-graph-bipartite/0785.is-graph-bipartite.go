package _785_is_graph_bipartite

// BFS
//func isBipartite(graph [][]int) bool {
//  visited := make([]int, len(graph))   // 0：未染色，1：蓝色，-1：黄色
//
//	for i := 0; i < len(graph); i++ { // 遍历每个顶点
//      if visited[i] != 0 {       // 已经访问过
//		   continue
//      }
//      q := []int{i} // 队列初始推入顶点 i
//      visited[i] = 1 // 染色为 1
//
//      for len(q) != 0 { // 遍历顶点 i 所有相邻的顶点
//          cur := q[0] // 出列的顶点
//          q = q[1:] // 维护队列
//          curColor := visited[cur] // 出列顶点的颜色
//          neighborColor := -curColor // 它的相邻顶点应该有的颜色
//
//          for _, neighbor := range graph[cur] { // 给他们都上色
//              if visited[neighbor] == 0 { // 如果没上色
//                  visited[neighbor] = neighborColor // 相邻顶点上色
//                  q = append(q, neighbor) // 并推入队列
//              }else if visited[neighbor] != neighborColor { // 上了不对的颜色
//                  return false
//              }
//          }
//      }
//  }
//  // 遍历完所有顶点，没有不对
//  return true
//}

// 记录图是否符合二分图性质
var ok bool
// 记录图中节点的颜色，false 和 true 代表两种颜色
// var color []bool
// 记录图中节点是否被访问过
var visited []int

// 输入邻接表，判断是否是二分图
func isBipartite(graph [][]int) bool {
	n := len(graph)
	ok = true
	visited = make([]int, n)
	for v := 0; v < n && ok; v++ {
		if visited[v] == 0 {
			traverse(graph, v)
		}
	}
	return ok
}

// DFS 遍历
func traverse(graph [][]int, cur int) {
	// 如果已经确定不是二分图了，就不用再递归遍历了
    if !ok {
        return
    }
	visited[cur] = 1
	curColor := visited[cur]
	neighborColor := -curColor
	for _, neighbor := range graph[cur] {
		if visited[neighbor] == 0 {
			// 相邻节点 neighbor 没有访问过
			// 那么应该给节点 neighbor 涂上和节点 cur 不同的颜色
			visited[neighbor] = neighborColor
			// color[w] = !color[v]
			// 继续遍历 neighbor
			traverse(graph, neighbor)
		}else {
			// 相邻节点 neighbor 已经被访问过
			// 根据 cur 和 neighbor 的颜色判断是否为二分图
			if visited[cur] == visited[neighbor] {
				// 若相同，则不是二分图
				ok = false
				return
			}
		}
	}
}