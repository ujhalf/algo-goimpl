package c

func taskSchedulerII(tasks []int, space int) int64 {
	m := map[int]int64{}
	for _, v := range tasks {
		m[v] = int64(-space - 1)
	}
	today := int64(1)
	for _, v := range tasks {
		if today-m[v] <= int64(space) {
			today = int64(space) + m[v] + 1
		}
		m[v] = today
		today++
	}
	return today - 1

}
