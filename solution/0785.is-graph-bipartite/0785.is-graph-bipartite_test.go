package _785_is_graph_bipartite

import (
	"fmt"
	"testing"
)

func Test0785(t *testing.T) {
	graph := [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}
	ans := isBipartite(graph)
	fmt.Print(ans)
}
