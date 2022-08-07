package c

func validPartition(nums []int) bool {
	n := len(nums)
	f := make([]bool, n+1)
	f[0] = true

	for i, x := range nums {
		if (i-1 >= 0 && f[i-1] && nums[i] == nums[i-1]) || (i-2 >= 0 && f[i-2] && ((nums[i-2]+1 == nums[i-1] && nums[i-1]+1 == x) || (x == nums[i-2] && x == nums[i-1]))) {
			f[i+1] = true
		}
	}
	return f[n]

}
