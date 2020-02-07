package My_LeetCode_In_Go

import (
	"math"

	. "My_LeetCode_In_Go/util/tree"
)

/*
  给定一个二叉树，判断其是否是一个有效的二叉搜索树。
  假设一个二叉搜索树具有如下特征：
    节点的左子树只包含小于当前节点的数。
    节点的右子树只包含大于当前节点的数。
    所有左子树和右子树自身必须也是二叉搜索树。
*/

func isValidBST(root *TreeNode) bool {
	var dfs func(int, int, *TreeNode) bool
	dfs = func(min, max int, root *TreeNode) bool {
		if root == nil {
			return true
		}
		if root.Val <= min || root.Val >= max {
			return false
		}
		return dfs(min, root.Val, root.Left) && dfs(root.Val, max, root.Right)
	}
	return dfs(math.MinInt64, math.MaxInt64, root)
}