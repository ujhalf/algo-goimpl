package c

const INF = int(1e7)

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	val := int(1e6)
	ret := -1
	a := dfs(edges, node1)
	b := dfs(edges, node2)
	for i := 0; i < len(edges); i++ {
		curVal := max(a[i], b[i])
		if curVal < val {
			val = curVal
			ret = i
		}
	}
	return ret
}

func dfs(e []int, source int) []int {
	ret := make([]int, len(e))
	for i := 0; i < len(e); i++ {
		ret[i] = INF
	}
	ret[source] = 0
	dis := 0
	for nxt := e[source]; nxt != -1 && ret[nxt] == INF; nxt = e[nxt] {
		dis++
		ret[nxt] = dis
	}
	return ret
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
