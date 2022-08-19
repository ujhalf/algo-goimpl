package p20

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return construct(nums, 0, len(nums)-1)
}

func construct(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	id := l
	for i := l + 1; i <= r; i++ {
		if nums[id] < nums[i] {
			id = i
		}
	}
	node := &TreeNode{
		Val: nums[id],
	}
	node.Left = construct(nums, l, id-1)
	node.Right = construct(nums, id+1, r)
	return node
}
