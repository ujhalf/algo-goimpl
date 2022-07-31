package p31

func maxLevelSum(root *TreeNode) int {
	val := -(int)(2e9)
	ans := -1
	q := []*TreeNode{root}
	for level := 1; len(q) > 0; level++ {
		tmp := q
		sum := 0
		q = nil
		for _, node := range tmp {
			sum += node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		if sum > val {
			val = sum
			ans = level
		}
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
