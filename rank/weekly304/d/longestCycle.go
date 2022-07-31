package main

import "fmt"

func longestCycle(edges []int) int {
	n := len(edges)
	ans := -1
	time := make([]int, n)
	clock := 0
	for i := 0; i < n; i++ {
		if time[i] == 0 {
			for j, start_time := i, clock; j != -1; j = edges[j] {
				if time[j] > 0 {
					if time[j] >= start_time {
						ans = max(ans, clock-time[j])
					}
					break
				}
				time[j] = clock
				clock++
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	edges := []int{3, 3, 4, 2, 3}
	cycle := longestCycle(edges)
	fmt.Println(cycle)
}
