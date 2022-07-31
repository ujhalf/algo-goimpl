package a

func minimumOperations(nums []int) int {
	rec := make(map[int]bool)
	for _, v := range nums {
		rec[v] = true
	}
	delete(rec, 0)
	return len(rec)

}
