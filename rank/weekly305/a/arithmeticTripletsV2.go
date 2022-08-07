package a

func arithmeticTripletsV2(nums []int, diff int) int {
	i, j := 0, 1
	ret := 0
	for _, x := range nums {
		for nums[j]+diff < x {
			j++
		}
		if nums[j]+diff != x {
			continue
		}
		for nums[i]+diff < nums[j] {
			i++
		}
		if nums[i]+diff == nums[j] {
			ret++
		}
	}
	return ret
}
