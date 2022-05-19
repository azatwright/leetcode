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
	nodeindex := map[int]int{}
	nodecounter := 0

	var dfs func(int) int
	dfs = func(node int) int {
		if n, ok := nodeindex[node]; ok {
			return n
		}
		nodeindex[node] = nodecounter
		nodecounter++

		// dfs may reach any visited node, thus detecting a loop.
		// Of all visited nodes that can be reached from the
		// current node, minindex is the index of the node that
		// was visited earlier than any other, thus detecting the
		// longest loop.
		//
		// Because dfs returns minindex, the parent can know if
		// its child makes a loop with it or its ancestors.
		//
		minindex := nodeindex[node]

		for _, child := range children[node] {
			if !edgeused[[2]int{node, child}] && !edgeused[[2]int{child, node}] {
				edgeused[[2]int{node, child}] = true
				if ret := dfs(child); ret <= nodeindex[node] {
					delete(edgedic, [2]int{node, child})
					delete(edgedic, [2]int{child, node})
					minindex = min(minindex, ret)
				}
			}
		}

		return minindex
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
