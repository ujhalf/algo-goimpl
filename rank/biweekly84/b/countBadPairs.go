package b

func countBadPairs(nums []int) int64 {
	m := map[int]int{}
	var ret int64
	for i, v := range nums {
		id := v - i
		ret += int64(i)
		ret -= int64(m[id])
		m[id]++
	}
	return ret
}
