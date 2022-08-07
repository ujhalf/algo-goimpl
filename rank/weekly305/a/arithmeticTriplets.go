package a

func arithmeticTriplets(nums []int, diff int) int {
	set := map[int]bool{}
	for _, v := range nums {
		set[v] = true
	}
	ret := 0
	for _, v := range nums {
		if set[v+diff] && set[v+2*diff] {
			ret++
		}
	}
	return ret
}
