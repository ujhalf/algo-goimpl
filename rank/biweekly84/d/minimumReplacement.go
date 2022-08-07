package d

func minimumReplacement(nums []int) int64 {
	ret := int64(0)
	min := int(1e9)
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] <= min {
			min = nums[i]
		} else {
			cnt := (nums[i] + min - 1) / min
			ret += int64(cnt - 1)
			min = nums[i] / cnt
		}
	}
	return ret

}
