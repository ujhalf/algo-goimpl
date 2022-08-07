package d

func longestIdealString(s string, k int) int {
	d := [26]int{}
	ret := 0
	for _, v := range s {
		id := int(v - 'a')
		for j := max(0, id-k); j <= min(25, id+k); j++ {
			d[id] = max(d[id], d[j])
		}
		d[id]++
		ret = max(ret, d[id])
	}
	return ret

}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
