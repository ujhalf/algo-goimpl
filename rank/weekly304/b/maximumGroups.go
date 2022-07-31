package b

import "math"

func maximumGroups(grades []int) int {

	return (int(math.Sqrt(float64(len(grades)*8+1))) - 1) / 2
}
