package a

import "sort"

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	m := map[int]int{}
	for _, v := range items1 {
		m[v[0]] += v[1]
	}
	for _, v := range items2 {
		m[v[0]] += v[1]
	}
	var ret [][]int
	for i, v := range m {
		ret = append(ret, []int{i, v})
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i][0] < ret[j][0]
	})
	return ret
}
