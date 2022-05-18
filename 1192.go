package main

func criticalConnections(n int, connections [][]int) [][]int {
	children := map[int][]int{}
	for _, c := range connections {
		children[c[0]] = append(children[c[0]], c[1])
		children[c[1]] = append(children[c[1]], c[0])
	}

	var buf []int
	critic := map[[2]int]bool{}

	var critical func(int, int) bool
	critical = func(a, b int) bool {
		seen := map[int]bool{a: true}
		q := buf[:0]
		q = append(q, a)
		for len(q) != 0 {
			prevlen := len(q)
			for i := 0; i < prevlen; i++ {
				if q[i] == b {
					return false
				}
				for _, child := range children[q[i]] {
					if q[i] == a && child == b {
						continue
					}
					if critic[[2]int{q[i], child}] {
						continue
					}
					if !seen[child] {
						q = append(q, child)
					}
				}
				seen[q[i]] = true
			}
			copy(q, q[prevlen:])
			q = q[:len(q[prevlen:])]
		}
		return true
	}

	out := [][]int{}
	for _, c := range connections {
		if critical(c[0], c[1]) {
			out = append(out, []int{c[0], c[1]})
			critic[[2]int{c[0], c[1]}] = true
			critic[[2]int{c[1], c[0]}] = true
		}
	}

	return out
}
