package p05

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		node := TreeNode{
			Val: val,
		}
		node.Left = root
		return &node
	} else {
		dfs(root, val, depth-1, 1)
		return root
	}

}

func dfs(root *TreeNode, val, target, cur int) {
	if root != nil {
		if cur == target {
			left := TreeNode{
				Val: val,
			}
			right := TreeNode{
				Val: val,
			}
			left.Left = root.Left
			right.Right = root.Right
			root.Left = &left
			root.Right = &right
		} else {
			dfs(root.Left, val, target, cur+1)
			dfs(root.Right, val, target, cur+1)
		}
	}
}
