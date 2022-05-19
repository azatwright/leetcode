package main

func criticalConnections(n int, connections [][]int) [][]int {
	edgedic := map[[2]int]bool{}
	for _, c := range connections {
		edgedic[[2]int{c[0], c[1]}] = true
	}

	children := map[int][]int{}
	for _, c := range connections {
		children[c[0]] = append(children[c[0]], c[1])
		children[c[1]] = append(children[c[1]], c[0])
	}

	edgeused := map[[2]int]bool{}
	nodenum := map[int]int{}
	nodecounter := 0

	var dfs func(int) int
	dfs = func(node int) int {
		if n, ok := nodenum[node]; ok {
			return n
		}
		nodenum[node] = nodecounter
		nodecounter++

		rank := nodecounter - 1
		for _, child := range children[node] {
			if !edgeused[[2]int{node, child}] && !edgeused[[2]int{child, node}] {
				edgeused[[2]int{node, child}] = true
				if ret := dfs(child); ret <= nodenum[node] {
					delete(edgedic, [2]int{node, child})
					delete(edgedic, [2]int{child, node})
					rank = min(rank, ret)
				}
			}
		}

		return rank
	}

	dfs(0)

	out := [][]int{}
	for c := range edgedic {
		out = append(out, []int{c[0], c[1]})
	}
	return out
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
