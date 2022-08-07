package b

func reachableNodes(n int, edges [][]int, restricted []int) int {
	set := map[int]bool{}

	for _, v := range restricted {
		set[v] = true
	}

	rec := map[int]map[int]bool{}

	for _, v := range edges {
		a := v[0]
		b := v[1]
		if !set[a] && !set[b] {
			aNb := rec[a]
			if aNb == nil {
				aNb = map[int]bool{}
			}
			aNb[b] = true
			rec[a] = aNb

			bNa := rec[b]
			if bNa == nil {
				bNa = map[int]bool{}
			}
			bNa[a] = true
			rec[b] = bNa
		}
	}
	seen := map[int]bool{}
	return dfs(rec, seen, 0)

}

func dfs(rec map[int]map[int]bool, seen map[int]bool, s int) int {
	if seen[s] {
		return 0
	}
	ret := 1
	aNb := rec[s]
	seen[s] = true
	if aNb != nil {
		for v := range aNb {
			ret += dfs(rec, seen, v)
		}
	}
	return ret
}
