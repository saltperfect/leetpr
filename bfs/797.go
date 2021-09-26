package bfs

type data struct {
	val   int
	route []int
}

func allPathsSourceTarget(graph [][]int) [][]int {
	target := len(graph) - 1
	var res [][]int
	var q, nq []data
	nq = []data{{0, []int{0}}}
	for len(nq) > 0 {
		q, nq = nq, make([]data, 0)
		for _, v := range q {
			if v.val == target {
				res = append(res, v.route)
				continue
			}
			for _, child := range graph[v.val] {
				nq = append(nq, data{child, append(v.route, child)})
			}
		}

	}
	return res
}
